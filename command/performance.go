package command

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/internal/sse/sysenv"
	"sagooiot/pkg/plugins"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/utility/utils"
)

// InitSystemStatistics 初始化系统统计
func initSystemStatistics(ctx context.Context) error {
	sysenv.LocalIP, _ = utils.GetLocalIP()       //获取本机IP
	sysenv.PublicIP, _ = utils.GetPublicIP()     //获取公网IP
	SysRunDir, _ := os.Getwd()                   //获取当前运行目录
	sysenv.GoDiskSize = utils.DirSize(SysRunDir) //获取当前运行目录磁盘大小

	return nil
}

// initPlugins 初始化插件
func initPlugins(ctx context.Context) (err error) {
	var inputData = new(model.GetSysPluginsListInput)
	inputData.PageSize = 1000
	inputData.Status = 1
	_, _, pluginDataList, err := service.SysPlugins().GetSysPluginsList(context.Background(), inputData)
	for _, p := range pluginDataList {
		//加载插件
		switch p.Types {
		case PluginType.Protocol:
			_, err = plugins.GetProtocolPlugin().GetProtocolByName(p.Name)
		case PluginType.Notice:
			_, err = plugins.GetNoticePlugin().GetNoticeByName(p.Name)
		}
		if err != nil {
			g.Log().Debug(ctx, p.Name, "插件加载出错", err.Error())
		}
	}

	//更新插件配置缓存
	err = service.SystemPluginsConfig().UpdateAllPluginsConfigCache(context.Background())
	if err != nil {
		return err
	}
	return
}
