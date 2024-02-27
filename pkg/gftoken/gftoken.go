package gftoken

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type GfToken struct {
	//  server name
	ServerName string
	// 缓存key (每创建一个实例CacheKey必须不相同)
	CacheKey string
	// 超时时间 默认10天（秒）
	Timeout int64
	// 缓存刷新时间 默认5天（秒）
	// 处理携带token的请求时当前时间大于超时时间并小于缓存刷新时间时token将自动刷新即重置token存活时间
	// MaxRefresh值为0时,token将不会自动刷新
	MaxRefresh int64
	// 是否允许多点登录
	MultiLogin bool
	// Token加密key 32位
	EncryptKey []byte
	// 缓存 (缓存模式:gcache 或 gredis)
	cache *gcache.Cache
	// 拦截排除地址
	ExcludePaths g.SliceStr
	// jwt
	userJwt *JwtSign
}

// TokenData Token 数据
type TokenData struct {
	JwtToken string `json:"jwtToken"`
	UuId     string `json:"uuId"`
}

// 存活时间 (存活时间 = 超时时间 + 缓存刷新时间)
func (m *GfToken) diedLine() time.Time {
	return time.Now().Add(time.Second * time.Duration(m.Timeout+m.MaxRefresh))
}

// 生成token
func (m *GfToken) GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error) {
	var (
		uuid   string
		tData  *TokenData
		tokens string
	)
	// 支持多端重复登录，返回新token
	if m.MultiLogin {
		tData, err = m.getCache(ctx, m.CacheKey+key)
		if err != nil {
			return
		}
		if tData != nil {
			key = gstr.SubStr(key, 0, len(key)-16) + grand.Letters(16)
			keys, uuid, err = m.EncryptToken(ctx, key, tData.UuId)
			m.doRefresh(ctx, key, tData) //刷新token
			return
		}
	}
	tokens, err = m.userJwt.CreateToken(CustomClaims{
		data,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-10, 0)), // 生效开始时间
			ExpiresAt: jwt.NewNumericDate(m.diedLine()),                       // 失效截止时间
		},
	})
	if err != nil {
		return
	}
	keys, uuid, err = m.EncryptToken(ctx, key)
	if err != nil {
		return
	}
	err = m.setCache(ctx, m.CacheKey+key, TokenData{
		JwtToken: tokens,
		UuId:     uuid,
	})
	if err != nil {
		return
	}
	return
}

// 解析token (只验证格式并不验证过期)
func (m *GfToken) ParseToken(r *ghttp.Request) (*CustomClaims, error) {
	token, err := m.GetToken(r)
	if err != nil {
		return nil, err
	}
	if customClaims, err := m.userJwt.ParseToken(token.JwtToken); err == nil {
		return customClaims, nil
	} else {
		return &CustomClaims{}, errors.New(ErrorsParseTokenFail)
	}
}

// 检查缓存的token是否有效且自动刷新缓存token
func (m *GfToken) IsEffective(ctx context.Context, token string) bool {
	cacheToken, key, err := m.GetTokenData(ctx, token)
	if err != nil {
		g.Log().Info(ctx, err)
		return false
	}
	_, code := m.IsNotExpired(cacheToken.JwtToken)
	if JwtTokenOK == code {
		// 刷新缓存
		if m.IsRefresh(cacheToken.JwtToken) {
			return m.doRefresh(ctx, key, cacheToken)
		}
		return true
	}
	return false
}

func (m *GfToken) doRefresh(ctx context.Context, key string, cacheToken *TokenData) bool {
	if newToken, err := m.RefreshToken(cacheToken.JwtToken); err == nil {
		cacheToken.JwtToken = newToken
		err = m.setCache(ctx, m.CacheKey+key, cacheToken)
		if err != nil {
			g.Log().Error(ctx, err)
			return false
		}
	}
	return true
}

func (m *GfToken) GetTokenData(ctx context.Context, token string) (tData *TokenData, key string, err error) {
	var uuid string
	key, uuid, err = m.DecryptToken(ctx, token)
	if err != nil {
		return
	}
	tData, err = m.getCache(ctx, m.CacheKey+key)
	if tData == nil || tData.UuId != uuid {
		err = gerror.New("token is invalid")
	}
	return
}

// 检查token是否过期 (过期时间 = 超时时间 + 缓存刷新时间)
func (m *GfToken) IsNotExpired(token string) (*CustomClaims, int) {
	if customClaims, err := m.userJwt.ParseToken(token); err == nil {
		if time.Now().Unix()-customClaims.ExpiresAt.Unix() < 0 {
			// token有效
			return customClaims, JwtTokenOK
		} else {
			// 过期的token
			return customClaims, JwtTokenExpired
		}
	} else {
		// 无效的token
		return customClaims, JwtTokenInvalid
	}
}

// 刷新token的缓存有效期
func (m *GfToken) RefreshToken(oldToken string) (newToken string, err error) {
	if newToken, err = m.userJwt.RefreshToken(oldToken, m.diedLine().Unix()); err != nil {
		return
	}
	return
}

// token是否处于刷新期
func (m *GfToken) IsRefresh(token string) bool {
	if m.MaxRefresh == 0 {
		return false
	}
	if customClaims, err := m.userJwt.ParseToken(token); err == nil {
		now := time.Now().Unix()
		if now < customClaims.ExpiresAt.Unix() && now > (customClaims.ExpiresAt.Unix()-m.MaxRefresh) {
			return true
		}
	}
	return false
}

// EncryptToken token加密方法
func (m *GfToken) EncryptToken(ctx context.Context, key string, randStr ...string) (encryptStr, uuid string, err error) {
	if key == "" {
		err = gerror.New("encrypt key empty")
		return
	}
	// 生成随机串
	if len(randStr) > 0 {
		uuid = randStr[0]
	} else {
		uuid = gmd5.MustEncrypt(grand.Letters(10))
	}
	token, err := gaes.Encrypt([]byte(key+uuid), m.EncryptKey)
	if err != nil {
		g.Log().Error(ctx, "[GFToken]encrypt error Token:", key, err)
		err = gerror.New("encrypt error")
		return
	}
	encryptStr = gbase64.EncodeToString(token)
	return
}

// DecryptToken token解密方法
func (m *GfToken) DecryptToken(ctx context.Context, token string) (DecryptStr, uuid string, err error) {
	if token == "" {
		err = gerror.New("decrypt Token empty")
		return
	}
	token64, err := gbase64.Decode([]byte(token))
	if err != nil {
		g.Log().Info(ctx, "[GFToken]decode error Token:", token, err)
		err = gerror.New("decode error")
		return
	}
	decryptToken, err := gaes.Decrypt(token64, m.EncryptKey)
	if err != nil {
		g.Log().Info(ctx, "[GFToken]decrypt error Token:", token, err)
		err = gerror.New("decrypt error")
		return
	}
	length := len(decryptToken)
	uuid = string(decryptToken[length-32:])
	DecryptStr = string(decryptToken[:length-32])
	return
}

// RemoveToken 删除token
func (m *GfToken) RemoveToken(ctx context.Context, token string) (err error) {
	var key string
	_, key, err = m.GetTokenData(ctx, token)
	if err != nil {
		return
	}
	err = m.removeCache(ctx, m.CacheKey+key)
	return
}
