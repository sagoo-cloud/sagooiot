package system

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/mssola/user_agent"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/utils"
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
func (s *sLogin) Login(ctx context.Context, verifyKey string, captcha string, userName string, password string) (loginUserOut *model.LoginUserOut, token string, err error) {
	//判断验证码是否正确
	if !service.Captcha().VerifyString(verifyKey, captcha) {
		err = gerror.New("验证码输入错误")
		return
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
			Module:    "系统后台",
		})
		return
	}
	//生成token
	key := gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + gmd5.MustEncryptString(userInfo.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(userInfo.Id) + "-" + gmd5.MustEncryptString(userInfo.UserName) + "-" + gmd5.MustEncryptString(userInfo.UserPassword+ip+userAgent)
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
		Username:  userName,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "系统后台",
	})

	//查看是否在线用户信息是否都存在不存在则删除
	userOnlines, _ := service.SysUserOnline().GetAll(ctx)
	if len(userOnlines) > 0 {
		var onlineIds []uint
		for _, online := range userOnlines {
			if !common.Cache().Get(ctx, online.Key).Bool() {
				onlineIds = append(onlineIds, online.Id)
			}
		}
		if len(onlineIds) > 0 {
			err = service.SysUserOnline().DelByIds(ctx, onlineIds)
		}
	}

	//保存在线用户信息
	ua := user_agent.New(userAgent)
	os := ua.OS()
	explorer, _ := ua.Browser()
	service.SysUserOnline().Invoke(ctx, &entity.SysUserOnline{
		Uuid:     guid.S(),
		UserName: userName,
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
	common.Cache().Remove(ctx, userOnline.Key)

	err = service.SysUserOnline().DelByToken(ctx, token[1])
	return
}
