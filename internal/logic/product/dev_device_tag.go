package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDevDeviceTag struct{}

func init() {
	service.RegisterDevDeviceTag(deviceTagNew())
}

func deviceTagNew() *sDevDeviceTag {
	return &sDevDeviceTag{}
}

func (s *sDevDeviceTag) Add(ctx context.Context, in *model.AddTagDeviceInput) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDeviceTag
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)

	_, err = dao.DevDeviceTag.Ctx(ctx).Data(param).Insert()
	return
}

func (s *sDevDeviceTag) Edit(ctx context.Context, in *model.EditTagDeviceInput) (err error) {
	total, _ := dao.DevDeviceTag.Ctx(ctx).Where(dao.DevDeviceTag.Columns().Id, in.Id).Count()
	if total == 0 {
		return gerror.New("标签不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDeviceTag
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevDeviceTag.Ctx(ctx).Data(param).Where(dao.DevDeviceTag.Columns().Id, in.Id).Update()
	return
}

func (s *sDevDeviceTag) Del(ctx context.Context, id uint) (err error) {
	total, _ := dao.DevDeviceTag.Ctx(ctx).Where(dao.DevDeviceTag.Columns().Id, id).Count()
	if total == 0 {
		return gerror.New("标签不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	_, err = dao.DevDeviceTag.Ctx(ctx).
		Data(do.DevDeviceTag{
			DeletedBy: uint(loginUserId),
			DeletedAt: gtime.Now(),
		}).
		Where(dao.DevDeviceTag.Columns().Id, id).
		Unscoped().
		Update()
	return
}
