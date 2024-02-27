package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mssola/useragent"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/utility/utils"
	"strconv"
	"strings"
)

type sLogin struct {
}

func loginNew() *sLogin {
	return &sLogin{}
}

func init() {
	service.RegisterLogin(loginNew())
}

// Login 登录
func (s *sLogin) Login(ctx context.Context, verifyKey string, captcha string, userName string, password string) (loginUserOut *model.LoginUserOut, token string, isChangePassword int, err error) {

	//判断验证码是否正确
	if !service.Captcha().VerifyString(verifyKey, captcha) {
		err = gerror.New("验证码输入错误")
		return
	}

	//判断是否启用了安全控制和启用了RSA
	configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysIsRsaEnabled}
	configDatas, err := service.ConfigData().GetByKeys(ctx, configKeys)
	if err != nil {
		return
	}
	isSecurityControlEnabled := "0" //是否启动安全控制
	isRsaEnbled := "0"              //是否启用RSA
	for _, configData := range configDatas {
		if strings.EqualFold(configData.ConfigKey, consts.SysIsSecurityControlEnabled) {
			isSecurityControlEnabled = configData.ConfigValue
		}
		if strings.EqualFold(configData.ConfigKey, consts.SysIsRsaEnabled) {
			isRsaEnbled = configData.ConfigValue
		}
	}

	if strings.EqualFold(isSecurityControlEnabled, "1") {
		//验证密码错误次数
		err = s.CheckPwdErrorNum(ctx, userName)
		if err != nil {
			return
		}

		//判断用户是否需要更改密码
		isChangePassword = s.IsChangePwd(ctx, userName)
		if isChangePassword == 1 {
			return
		}

		if strings.EqualFold(isRsaEnbled, "1") {
			//对账号进行解密
			password, err = utils.Decrypt(consts.RsaPrivateKeyFile, password, consts.RsaOAEP)
			if err != nil {
				return
			}
		}
	}

	//获取IP地址
	ip := utils.GetClientIp(ctx)
	//获取user-agent
	userAgent := utils.GetUserAgent(ctx)
	//根据账号密码获取用户信息
	userInfo, err := service.SysUser().GetAdminUserByUsernamePassword(ctx, userName, password)
	if err != nil {
		// 保存登录失败的日志信息
		service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
			Status:    0,
			Username:  userName,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "账号密码登录",
		})
		return
	}

	loginUserOut, token, err = s.GenUserToken(ctx, isSecurityControlEnabled, ip, userAgent, userInfo, "账号密码登录")
	if err != nil {
		return
	}

	return
}

// CheckPwdErrorNum 验证密码错误次数
func (s *sLogin) CheckPwdErrorNum(ctx context.Context, userName string) (err error) {
	tmpData, err := cache.Instance().Get(ctx, consts.CacheSysErrorPrefix+"_"+gconv.String(userName))
	if tmpData.Val() != nil {
		//获取密码错误次数和限制登录时间
		configKeys := []string{consts.SysPasswordErrorNum, consts.SysAgainLoginDate}
		var configDatas []*entity.SysConfig
		configDatas, err = service.ConfigData().GetByKeys(ctx, configKeys)
		if err != nil {
			return
		}
		errorNum := 3       //密码错误次数
		againLoginDate := 1 //限制登录时间
		for _, configData := range configDatas {
			if strings.EqualFold(configData.ConfigKey, consts.SysPasswordErrorNum) {
				errorNum, err = strconv.Atoi(configData.ConfigValue)
				if err != nil {
					err = gerror.New("密码错误次数配置错误")
					return
				}
			}
			if strings.EqualFold(configData.ConfigKey, consts.SysAgainLoginDate) {
				againLoginDate, err = strconv.Atoi(configData.ConfigValue)
				if err != nil {
					err = gerror.New("允许再次登录时间配置错误")
					return
				}
			}
		}

		tempValue, _ := strconv.Atoi(tmpData.Val().(string))
		if tempValue >= errorNum {
			err = gerror.Newf("密码错误次数过多, 请%d分钟后再重试", againLoginDate)
			return
		}
	}
	return
}
func (s *sLogin) IsChangePwd(ctx context.Context, userName string) (isChangePwd int) {

	changePasswordForFirstLogin := "0" //是否开启首次登录更改密码
	passwordChangePeriod := "90"       //是否开启是否密码定期更换

	//判断是否启用了安全控制和启用了RSA
	configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysChangePasswordForFirstLogin, consts.SysPasswordChangePeriodSwitch}
	configDatas, err := service.ConfigData().GetByKeys(ctx, configKeys)
	if err != nil {
		return
	}
	for _, configData := range configDatas {
		if strings.EqualFold(configData.ConfigKey, consts.SysChangePasswordForFirstLogin) {
			changePasswordForFirstLogin = configData.ConfigValue
		}
		if strings.EqualFold(configData.ConfigKey, consts.SysPasswordChangePeriodSwitch) {
			passwordChangePeriod = configData.ConfigValue
		}
	}

	var userInfo *entity.SysUser
	userInfo, err = service.SysUser().GetUserByUsername(ctx, userName)
	if userInfo == nil {
		isChangePwd = 0
		return
	}

	//获取用户修改密码的历史记录
	var history []*entity.SysUserPasswordHistory
	err = dao.SysUserPasswordHistory.Ctx(ctx).Where(dao.SysUserPasswordHistory.Columns().UserId, userInfo.Id).OrderDesc(dao.SysUserPasswordHistory.Columns().CreatedAt).Scan(&history)

	if strings.EqualFold(changePasswordForFirstLogin, "1") {
		if history == nil {
			isChangePwd = 1
			return
		}
	}

	if strings.EqualFold(passwordChangePeriod, "1") {
		//获取密码更换周期录系统参数
		var configDataByPwdChangePeriod *entity.SysConfig
		configDataByPwdChangePeriod, err = service.ConfigData().GetConfigByKey(ctx, consts.SysPasswordChangePeriod)
		if err != nil {
			return
		}
		if configDataByPwdChangePeriod != nil {
			//获取用户修改密码的历史记录
			// 将字符串转换为整数
			var days int
			days, err = strconv.Atoi(configDataByPwdChangePeriod.ConfigValue)
			if err != nil {
				fmt.Println("无法将字符串转换为整数:", err)
				return
			}
			changeTime := gtime.Now()
			if history != nil || len(history) > 0 {
				changeTime = history[0].ChangeTime
			}
			if changeTime.AddDate(0, 0, days).Before(gtime.Now()) {
				isChangePwd = 1
				return
			}
		}
	}

	return
}

// GenUserToken 生成用户TOKEN
func (s *sLogin) GenUserToken(ctx context.Context, isSecurityControlEnabled string, ip string, userAgent string, userInfo *entity.SysUser, logMoudel string) (loginUserOut *model.LoginUserOut, token string, err error) {
	var configData *entity.SysConfig
	if strings.EqualFold(isSecurityControlEnabled, "1") {
		//获取是否单一登录系统参数
		configData, err = service.ConfigData().GetConfigByKey(ctx, consts.SysIsSingleLogin)
		if err != nil {
			return
		}
	}

	//生成token
	key := "Login:" + gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName)
	if configData != nil && strings.EqualFold(configData.ConfigValue, "1") {
		key = "Login:" + gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + "-" + gmd5.MustEncryptString(ip+userAgent)
	}
	userInfo.UserPassword = ""
	token, err = service.SysToken().GenerateToken(ctx, key, userInfo)
	if err != nil {
		return
	}
	//转换用户信息
	if err = gconv.Scan(userInfo, &loginUserOut); err != nil {
		return
	}
	//修改用户信息
	err = service.SysUser().UpdateLoginInfo(ctx, userInfo.Id, ip)
	if err != nil {
		return
	}
	// 保存登录成功的日志信息
	service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
		Status:    1,
		Username:  userInfo.UserName,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    logMoudel,
	})

	//查看是否在线用户信息是否都存在不存在则删除
	userOnlines, _ := service.SysUserOnline().GetAll(ctx)
	if len(userOnlines) > 0 {
		var onlineIds []int
		for _, online := range userOnlines {
			onliceCache, _ := cache.Instance().Get(ctx, online.Key)
			if !onliceCache.Bool() {
				onlineIds = append(onlineIds, online.Id)
			}
		}
		if len(onlineIds) > 0 {
			err = service.SysUserOnline().DelByIds(ctx, onlineIds)
		}
	}

	//保存在线用户信息
	ua := useragent.New(userAgent)
	os := ua.OS()
	explorer, _ := ua.Browser()
	service.SysUserOnline().Invoke(ctx, &entity.SysUserOnline{
		Uuid:     guid.S(),
		UserName: userInfo.UserName,
		Key:      key,
		Token:    token,
		Ip:       ip,
		Os:       os,
		Explorer: explorer,
	})
	return
}

func (s *sLogin) LoginOut(ctx context.Context) (err error) {
	authorization := ghttp.RequestFromCtx(ctx).Header.Get("Authorization")
	if authorization == "" {
		err = gerror.New("请先登录!")
		return
	}
	//查看是否已经登录
	token := strings.Split(authorization, " ")
	if len(token) < 2 || len(token) >= 2 && token[1] == "" {
		err = gerror.New("TOKEN错误!")
		return
	}
	userOnline, _ := service.SysUserOnline().GetInfoByToken(ctx, token[1])
	if userOnline == nil {
		err = gerror.New("未登录,无法退出!")
		return
	}
	//增加删除缓存信息

	loginUserId := service.Context().GetUserId(ctx)
	_, err = cache.Instance().Remove(ctx, userOnline.Key)
	_, err = cache.Instance().Remove(ctx, consts.CacheUserAuthorize+"_"+gconv.String(loginUserId))
	_, err = cache.Instance().Remove(ctx, consts.CacheUserInfo+"_"+gconv.String(loginUserId))

	err = service.SysUserOnline().DelByToken(ctx, token[1])
	return
}
