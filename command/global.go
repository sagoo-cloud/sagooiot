package command

import (
	"context"
	"sagooiot/internal/consts"
	"sagooiot/internal/logic/analysis"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/statistics"
	"sagooiot/pkg/utility/helper"
	"sagooiot/utility/validate"

	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmode"
)

func AllSystemInit(ctx context.Context) {
	// 设置gf运行模式
	//SetGFMode(ctx)
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}
	// 初始化链路追踪
	InitTrace(ctx)
	// 设置缓存适配器
	cache.SetAdapter(ctx)
	// 初始化统计设备数据
	statistics.InitCountDeviceData()
	// 初始化系统配置参数
	err := dcache.InitSystemConfig(ctx)
	if err != nil {
		g.Log().Debug(ctx, "初始化系统配置参数失败")
	}

	//清除设备统计缓存
	analysis.RemoveDeviceStatusCountCache(ctx)

}

// SetGFMode 设置gf运行模式
func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(helper.AppName(ctx), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	helper.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}
