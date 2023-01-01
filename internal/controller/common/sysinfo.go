package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/sagoo-cloud/sagooiot/api/v1/common"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/version"
)

type cSysInfo struct{}

var SysInfo = cSysInfo{}

func (s *cSysInfo) GetSysInfo(ctx context.Context, req *common.GetSysInfoReq) (res *common.GetSysInfoRes, err error) {

	cfgSystemName, err := service.ConfigData().GetConfigByKey(ctx, "sys.system.name")
	systemName := "沙果IOT"
	if cfgSystemName != nil {
		systemName = cfgSystemName.ConfigValue
	}

	cfgSystemCopyright, err := service.ConfigData().GetConfigByKey(ctx, "sys.system.copyright")
	systemCopyright := "Sagoo inc."
	if cfgSystemName != nil {
		systemCopyright = cfgSystemCopyright.ConfigValue
	}

	res = &common.GetSysInfoRes{
		"systemName":      systemName,
		"systemCopyright": systemCopyright,
		"buildVersion":    version.BuildVersion,
		"buildTime":       version.BuildTime,
	}

	return
}

// IsToken 验证token是否正确
func (s *cSysInfo) IsToken(ctx context.Context, req *common.IsTokenReq) (res *common.IsTokenRes, err error) {
	authorization := ghttp.RequestFromCtx(ctx).Header.Get("Authorization")
	if authorization == "" {
		err = gerror.New("请先登录!")
		return
	}
	var isToken = false
	var expiresAt int64
	//验证TOKEN是否正确
	data, _ := service.SysToken().ParseToken(ghttp.RequestFromCtx(ctx))
	if data != nil {
		isToken = true
		expiresAt = data.ExpiresAt
	}
	res = &common.IsTokenRes{
		IsToken:   isToken,
		ExpiresAt: expiresAt,
	}
	return
}
