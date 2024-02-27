package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysNotifications = cSysNotifications{}

type cSysNotifications struct{}

// 获取列表
func (u *cSysNotifications) GetNotificationsList(ctx context.Context, req *system.GetNotificationsListReq) (res *system.GetNotificationsListRes, err error) {
	var input *model.GetNotificationsListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, currentPage, out, err := service.SysNotifications().GetSysNotificationsList(ctx, input)
	if err != nil {
		return
	}
	res = new(system.GetNotificationsListRes)
	res.Total = total
	res.CurrentPage = currentPage
	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}
	return
}

// 获取指定ID数据
func (u *cSysNotifications) GetNotificationsById(ctx context.Context, req *system.GetNotificationsByIdReq) (res *system.GetNotificationsByIdRes, err error) {
	data, err := service.SysNotifications().GetSysNotificationsById(ctx, req.Id)
	res = new(system.GetNotificationsByIdRes)
	err = gconv.Scan(data, &res.Data)
	return
}

// 添加数据
func (u *cSysNotifications) AddSysNotifications(ctx context.Context, req *system.AddNotificationsReq) (res *system.AddNotificationsRes, err error) {
	var data = model.NotificationsAddInput{}
	err = gconv.Scan(req, &data)
	err = service.SysNotifications().AddSysNotifications(ctx, data)
	return
}

// 修改数据
func (u *cSysNotifications) EditSysNotifications(ctx context.Context, req *system.EditNotificationsReq) (res *system.EditNotificationsRes, err error) {
	var data = model.NotificationsEditInput{}
	err = gconv.Scan(req, &data)
	err = service.SysNotifications().EditSysNotifications(ctx, data)
	return
}

// 删除数据
func (u *cSysNotifications) DeleteSysNotifications(ctx context.Context, req *system.DeleteNotificationsReq) (res *system.DeleteNotificationsRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.SysNotifications().DeleteSysNotifications(ctx, req)
	return
}
