package common

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"math/rand"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/utility/utils"
	"sagooiot/pkg/utility/version"
	"strings"
	"time"
)

type sSysInfo struct {
}

func sysInfo() *sSysInfo {
	return &sSysInfo{}
}

func init() {
	service.RegisterSysInfo(sysInfo())
}

func (s *sSysInfo) GetSysInfo(ctx context.Context) (out g.Map, err error) {
	cfgSystemName, err := service.ConfigData().GetConfigByKey(ctx, consts.SysSystemName)
	systemName := "沙果IOT"
	if cfgSystemName != nil {
		systemName = cfgSystemName.ConfigValue
	}

	cfgSystemCopyright, err := service.ConfigData().GetConfigByKey(ctx, consts.SysSystemCopyright)
	systemCopyright := "Sagoo inc."
	if cfgSystemName != nil {
		systemCopyright = cfgSystemCopyright.ConfigValue
	}

	cfgSystemLogo, err := service.ConfigData().GetConfigByKey(ctx, consts.SysSystemLogo)
	systemLogo := ""
	if cfgSystemLogo != nil {
		systemLogo = cfgSystemLogo.ConfigValue
	}
	cfgSystemLogoMini, err := service.ConfigData().GetConfigByKey(ctx, consts.SysSystemLogoMini)
	systemLogoMini := ""
	if cfgSystemLogoMini != nil {
		systemLogoMini = cfgSystemLogoMini.ConfigValue
	}
	cfgSystemLoginPic, err := service.ConfigData().GetConfigByKey(ctx, consts.SysSystemLoginPic)
	systemLoginPic := ""
	if cfgSystemLoginPic != nil {
		systemLoginPic = cfgSystemLoginPic.ConfigValue
	}
	cfgHomePageRoute, err := service.ConfigData().GetConfigByKey(ctx, consts.HomePageRoute)
	systemHomePageRoute := ""
	if cfgHomePageRoute != nil {
		systemHomePageRoute = cfgHomePageRoute.ConfigValue
	}

	cfgSysPasswordChangePeriod, err := service.ConfigData().GetConfigByKey(ctx, consts.SysPasswordChangePeriod)
	sysPasswordChangePeriod := "90"
	if cfgSysPasswordChangePeriod != nil {
		sysPasswordChangePeriod = cfgSysPasswordChangePeriod.ConfigValue
	}

	cfgSysIsSecurityControlEnabled, err := service.ConfigData().GetConfigByKey(ctx, consts.SysIsSecurityControlEnabled)
	isSecurityControlEnabled := "0"
	if cfgSysIsSecurityControlEnabled != nil {
		isSecurityControlEnabled = cfgSysIsSecurityControlEnabled.ConfigValue
	}

	cfgSysIsRsaEnabled, err := service.ConfigData().GetConfigByKey(ctx, consts.SysIsRsaEnabled)
	isRsaEnabled := "0"
	if cfgSysIsRsaEnabled != nil {
		isRsaEnabled = cfgSysIsRsaEnabled.ConfigValue
	}

	out = g.Map{
		"systemName":          systemName,
		"systemCopyright":     systemCopyright,
		"systemLogo":          systemLogo,
		"systemLogoMini":      systemLogoMini,
		"systemLoginPIC":      systemLoginPic,
		"systemHomePageRoute": systemHomePageRoute,

		"buildVersion": version.BuildVersion,
		"buildTime":    version.BuildTime,
		"commitID":     version.CommitID,
		"target":       gbase64.EncodeToString([]byte(sysPasswordChangePeriod + "|" + isSecurityControlEnabled + "|" + isRsaEnabled + "|SAGOOIOT")),
	}
	return
}

// ServerInfoEscalation 客户端服务信息上报
func (s *sSysInfo) ServerInfoEscalation(ctx context.Context) (err error) {
	num := rand.Intn(10)
	time.Sleep(time.Duration(num) * time.Second)

	ip, err := utils.GetLocalIP()
	if err != nil {
		err = gerror.New("获取客户端信息失败")
		return
	}
	var tmpData *gvar.Var
	tmpData, err = cache.Instance().Get(ctx, consts.CacheServerInfo)
	var serverInfos []map[string]interface{}
	if tmpData.Val() != nil {
		err = json.Unmarshal([]byte(tmpData.Val().(string)), &serverInfos)
		if err != nil {
			err = gerror.New("解析已上报客户端信息失败")
			return
		}
	}
	//判断IP是否存在
	isExist := false
	if serverInfos != nil && len(serverInfos) > 0 {
		for _, serverInfo := range serverInfos {
			if strings.EqualFold(serverInfo["ip"].(string), ip) {
				isExist = true
				//重新启动时间
				serverInfo["date"] = gtime.Now()
				break
			}

		}
	}
	if !isExist {
		serverInfo := make(map[string]interface{})
		serverInfo["ip"] = ip
		serverInfo["date"] = gtime.Now()
		serverInfos = append(serverInfos, serverInfo)
	}
	err = cache.Instance().Set(ctx, consts.CacheServerInfo, serverInfos, 0)

	return
}
