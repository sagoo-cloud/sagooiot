package command

import (
	"context"
	"fmt"
	_ "sagooiot/internal/logic"
	"sagooiot/internal/service"
	"sagooiot/network"
	"sagooiot/network/core/logic/model"

	"github.com/gogf/gf/v2/frame/g"
)

type NoDeferFunc struct {
	F    func(ctx context.Context) error
	Desc string
}

var InitFuncNoDeferListForIotCore = []NoDeferFunc{
	{model.InitCoreLogic, "核心处理逻辑"},
	{service.TSLTable().CreateDatabase, "时序数据库创建"},
	{service.TdLogTable().CreateStable, "时序库日志表创建"},
	{service.DevInit().InitProductForTd, "时序库产品表初始化"},
	{service.DevInit().InitDeviceForTd, "时序库设备表初始化"},
	{service.DevDevice().CacheDeviceDetailList, "缓存设备信息"},
	{service.AlarmRule().CacheAllAlarmRule, "缓存告警规则"},
	{network.ReloadNetwork, "网络服务"},
}

var InitFuncNoDeferListWebAdmin = []NoDeferFunc{
	{service.SysAuthorize().InitAuthorize, "系统权限"},
	{initSystemStatistics, "系统统计"},
	{service.SysInfo().ServerInfoEscalation, "集群数据"},
	{initPlugins, "插件"},
}

func InitSystem(ctx context.Context, noDeferFuncList []NoDeferFunc) error {
	for _, funcNode := range noDeferFuncList {
		g.Log().Infof(ctx, "开始初始化%s", funcNode.Desc)
		if err := funcNode.F(ctx); err != nil {
			return fmt.Errorf("初始化%s失败，错误原因是:%w", funcNode.Desc, err)
		} else {
			g.Log().Infof(ctx, "初始化%s成功", funcNode.Desc)
		}
	}
	return nil
}

var initFuncWithDeferList = []DeferFunc{
	{RunQueue, "消息队列"},
	{wrapperMqtt, "mqtt连接"},
}

func InitSystemDeferFunc(ctx context.Context) ([]func(context.Context) error, error) {
	deferFuncList := make([]func(context.Context) error, len(initFuncWithDeferList))
	for index, deferFuncNode := range initFuncWithDeferList {
		g.Log().Infof(ctx, "开始初始化%s", deferFuncNode.Desc)
		if err, deferFunc := deferFuncNode.F(ctx); err != nil {
			return nil, fmt.Errorf("初始化%s失败，错误原因是:%w", deferFuncNode.Desc, err)
		} else {
			deferFuncList[index] = deferFunc
			g.Log().Infof(ctx, "初始化%s成功", deferFuncNode.Desc)
		}
	}
	return deferFuncList, nil

}
