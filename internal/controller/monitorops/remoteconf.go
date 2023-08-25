package monitorops

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/monitorops"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var Remoteconf = cMonitoropsRemoteconf{}

type cMonitoropsRemoteconf struct{}

// GetRemoteconfList 获取列表
func (u *cMonitoropsRemoteconf) GetRemoteconfList(ctx context.Context, req *monitorops.GetRemoteconfListReq) (res *monitorops.GetRemoteconfListRes, err error) {
	dataList, err := service.MonitoropsRemoteconf().GetRemoteconfList(ctx, &req.GetRemoteconfListInput)
	res = new(monitorops.GetRemoteconfListRes)
	res.Data = dataList
	return
}

// GetRemoteconfById 获取指定ID数据
func (u *cMonitoropsRemoteconf) GetRemoteconfById(ctx context.Context, req *monitorops.GetRemoteconfByIdReq) (res *monitorops.GetRemoteconfByIdRes, err error) {
	data, err := service.MonitoropsRemoteconf().GetRemoteconfById(ctx, req.Id)
	res = new(monitorops.GetRemoteconfByIdRes)
	gconv.Scan(data, &res)
	return
}

// AddRemoteconf 添加数据
func (u *cMonitoropsRemoteconf) AddRemoteconf(ctx context.Context, req *monitorops.AddRemoteconfReq) (res *monitorops.AddRemoteconfRes, err error) {
	userInfo := service.Context().GetLoginUser(ctx)
	if userInfo == nil {
		err = gerror.New("未登录或TOKEN失效,请重新登录")
		return
	}
	glog.Infof(ctx,fmt.Sprintf("配置文件字节数：%d", len([]byte(req.ConfigContent))))
	if len(req.ConfigSize) > consts.RemoteconfLimitKB {
		err = gerror.New("配置文件上限64KB,请重新修改配置")
		return
	}
	if len([]byte(req.ConfigContent)) > consts.RemoteconfLimitB {
		err = gerror.New("配置文件上限64KB,请重新修改配置")
		return
	}
	// data.CreatedBy = userInfo.Id
	err = service.MonitoropsRemoteconf().AddRemoteconf(ctx, *req.RemoteconfAddInput)
	return
}

// EditRemoteconf 修改数据
func (u *cMonitoropsRemoteconf) EditRemoteconf(ctx context.Context, req *monitorops.EditRemoteconfReq) (res *monitorops.EditRemoteconfRes, err error) {
	userInfo := service.Context().GetLoginUser(ctx)
	if userInfo == nil {
		err = gerror.New("未登录或TOKEN失效,请重新登录")
		return
	}
	// data.UpdateBy = userInfo.Id //如果需要保存信息，把这个打开
	err = service.MonitoropsRemoteconf().EditRemoteconf(ctx, *req.RemoteconfEditInput)
	return
}

// DeleteRemoteconf 删除数据
func (u *cMonitoropsRemoteconf) DeleteRemoteconf(ctx context.Context, req *monitorops.DeleteRemoteconfReq) (res *monitorops.DeleteRemoteconfRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.MonitoropsRemoteconf().DeleteRemoteconf(ctx, req.Ids)
	return
}
