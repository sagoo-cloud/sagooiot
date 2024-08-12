package oauth

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/utility/utils"
	"strings"
)

var ErrRecordNotFound = errors.New("Record not found")

type sOauthUser struct{}

func init() {
	service.RegisterOauthUser(oauthUserNew())
}

func oauthUserNew() *sOauthUser {
	return &sOauthUser{}
}

func (s *sOauthUser) GetByUserId(ctx context.Context, userId uint64) (out []*entity.OauthUser, err error) {
	err = dao.OauthUser.Ctx(ctx).Where(dao.OauthUser.Columns().UserId, userId).Scan(&out)
	return
}

func (s *sOauthUser) SaveUser(ctx context.Context, in *entity.OauthUser) (out *entity.OauthUser, err error) {
	out = &entity.OauthUser{
		Openid:    in.Openid,
		AvatarUrl: in.AvatarUrl,
		Provider:  in.Provider,
		UpdatedAt: gtime.Now(),
		CreatedAt: gtime.Now(),
		Nickname:  in.Nickname,
	}
	user, err := s.GetUser(ctx, out.Provider, out.Openid)
	if errors.Is(err, ErrRecordNotFound) {
		if _, err = dao.OauthUser.Ctx(ctx).Save(out); err != nil {
			return nil, err
		}
		return out, nil
	} else if err != nil {
		return nil, err
	}
	_, err = dao.OauthUser.Ctx(ctx).
		Where(dao.OauthUser.Columns().Provider, user.Provider).
		Where(dao.OauthUser.Columns().Openid, user.Openid).Update(user)
	return user, err
}

func (s *sOauthUser) GetUser(ctx context.Context, provider string, openid string) (out *entity.OauthUser, err error) {
	err = dao.OauthUser.Ctx(ctx).
		Where(dao.OauthUser.Columns().Provider, provider).
		Where(dao.OauthUser.Columns().Openid, openid).
		Scan(&out)
	if out == nil && err == nil {
		return nil, ErrRecordNotFound
	}
	return
}
func (s *sOauthUser) UserBinding(ctx context.Context, provider string, openid string) (out *entity.OauthUser, err error) {
	loginUser := service.Context().GetLoginUser(ctx)
	if loginUser == nil {
		return nil, errors.New("需要先登录系统进行绑定！")
	}
	out, err = s.GetUser(ctx, provider, openid)
	if err != nil {
		return nil, err
	}
	out.UserId = loginUser.Id
	_, err = dao.OauthUser.Ctx(ctx).
		Where(dao.OauthUser.Columns().Provider, provider).
		Where(dao.OauthUser.Columns().Openid, openid).Update(out)
	return
}

func (s *sOauthUser) List(ctx context.Context, in *entity.OauthUser) (out []*entity.OauthUser, err error) {
	find := dao.OauthUser.Ctx(ctx)
	if in.Openid != "" {
		find = find.Where(dao.OauthUser.Columns().Openid, in.Openid)
	}
	if in.UserId > 0 {
		find = find.Where(dao.OauthUser.Columns().UserId, in.UserId)
	}
	if in.Provider != "" {
		find = find.Where(dao.OauthUser.Columns().Provider, in.Provider)
	}
	if in.Nickname != "" {
		find = find.Where(dao.OauthUser.Columns().Nickname, in.Nickname)
	}
	err = find.Scan(&out)
	return
}

func (s *sOauthUser) AuthLogin(ctx context.Context, in *entity.OauthUser) (*model.LoginUserOut, string, error) {
	//判断是否启用了安全控制和启用了RSA
	configKeys := []string{consts.SysIsSecurityControlEnabled, consts.SysIsRsaEnabled}
	configDatas, err := service.ConfigData().GetByKeys(ctx, configKeys)
	if err != nil {
		return nil, "", err
	}

	isSecurityControlEnabled := "0" //是否启动安全控制
	for _, configData := range configDatas {
		if strings.EqualFold(configData.ConfigKey, consts.SysIsSecurityControlEnabled) {
			isSecurityControlEnabled = configData.ConfigValue
		}
	}

	//获取IP地址
	ip := utils.GetClientIp(ctx)
	//获取user-agent
	userAgent := utils.GetUserAgent(ctx)
	u, err := service.SysUser().GetUserById(ctx, uint(in.UserId))
	if err != nil {
		return nil, "", err
	}
	//根据账号用户信息
	userInfo, err := service.SysUser().GetUserByUsername(ctx, u.UserName)
	if err != nil {
		// 保存登录失败的日志信息
		service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
			Status:    0,
			Username:  u.UserName,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    in.Provider + "授权登录",
		})
		return nil, "", err
	}
	return service.Login().GenUserToken(ctx, isSecurityControlEnabled, ip, userAgent, userInfo, in.Provider+"授权登录")
}
