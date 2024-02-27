package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/api/v1/system"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"
)

type sSysNotifications struct{}

func sSysNotificationsNew() *sSysNotifications {
	return &sSysNotifications{}
}
func init() {
	service.RegisterSysNotifications(sSysNotificationsNew())
}

// GetSysNotificationsList 获取列表数据
func (s *sSysNotifications) GetSysNotificationsList(ctx context.Context, input *model.GetNotificationsListInput) (total, page int, list []*model.NotificationsOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysNotifications.Ctx(ctx)
		total, err = m.Count()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		page = input.PageNum
		if input.PageSize == 0 {
			input.PageSize = consts.PageSize
		}
		err = m.Page(page, input.PageSize).Order("created_at desc").Scan(&list)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// GetSysNotificationsById 获取指定ID数据
func (s *sSysNotifications) GetSysNotificationsById(ctx context.Context, id int) (out *model.NotificationsRes, err error) {
	err = dao.SysNotifications.Ctx(ctx).Where("id", id).Scan(&out)
	return
}

// AddSysNotifications 添加数据
func (s *sSysNotifications) AddSysNotifications(ctx context.Context, in model.NotificationsAddInput) (err error) {
	_, err = dao.SysNotifications.Ctx(ctx).Data(do.SysNotifications{
		Title:     in.Title,
		Doc:       in.Doc,
		Source:    in.Source,
		Types:     in.Types,
		CreatedAt: gtime.Now(),
		Status:    in.Status,
	}).Insert()
	return
}

// EditSysNotifications 修改数据
func (s *sSysNotifications) EditSysNotifications(ctx context.Context, in model.NotificationsEditInput) (err error) {
	_, err = dao.SysNotifications.Ctx(ctx).FieldsEx(dao.SysNotifications.Columns().Id).Where(dao.SysNotifications.Columns().Id, in.Id).Update(in)
	return
}

// DeleteSysNotifications 删除数据
func (s *sSysNotifications) DeleteSysNotifications(ctx context.Context, in *system.DeleteNotificationsReq) (err error) {
	_, err = dao.SysNotifications.Ctx(ctx).Delete(dao.SysNotifications.Columns().Id+" in (?)", in.Ids)
	return
}
