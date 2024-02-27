package product

import (
	"context"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"

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
	/*param.CreateBy = uint(loginUserId)*/

	_, err = dao.DevDeviceTag.Ctx(ctx).Data(do.DevDeviceTag{
		DeptId:    service.Context().GetUserDeptId(ctx),
		DeviceId:  param.DeviceId,
		DeviceKey: param.DeviceKey,
		Key:       param.Key,
		Name:      param.Name,
		Value:     param.Value,
		CreatedBy: uint(loginUserId),
		CreatedAt: gtime.Now(),
	}).Insert()
	return
}

func (s *sDevDeviceTag) Edit(ctx context.Context, in *model.EditTagDeviceInput) (err error) {
	var deviceTag *entity.DevDeviceTag
	err = dao.DevDeviceTag.Ctx(ctx).Where(dao.DevDeviceTag.Columns().Id, in.Id).Scan(&deviceTag)
	if deviceTag == nil {
		return gerror.New("标签不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DevDeviceTag
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdatedBy = uint(loginUserId)
	param.Id = nil

	_, err = dao.DevDeviceTag.Ctx(ctx).Data(param).Where(dao.DevDeviceTag.Columns().Id, in.Id).Update()
	return
}

func (s *sDevDeviceTag) Del(ctx context.Context, id uint) (err error) {
	var deviceTag *entity.DevDeviceTag
	err = dao.DevDeviceTag.Ctx(ctx).Where(dao.DevDeviceTag.Columns().Id, id).Scan(&deviceTag)
	if deviceTag == nil {
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

func (s *sDevDeviceTag) Update(ctx context.Context, deviceId uint, list []model.AddTagDeviceInput) (err error) {
	var tagIds []int
	var add []model.AddTagDeviceInput
	for _, v := range list {
		rs, err := dao.DevDeviceTag.Ctx(ctx).
			Fields(dao.DevDeviceTag.Columns().Id).
			Where(dao.DevDeviceTag.Columns().DeviceId, deviceId).
			Where(dao.DevDeviceTag.Columns().Key, v.Key).
			Where(dao.DevDeviceTag.Columns().Name, v.Name).
			Where(dao.DevDeviceTag.Columns().Value, v.Value).
			Value()
		if err != nil {
			return err
		}
		if rs.Int() > 0 {
			tagIds = append(tagIds, rs.Int())
		} else {
			add = append(add, v)
		}
	}
	if len(tagIds) > 0 {
		_, err = dao.DevDeviceTag.Ctx(ctx).
			Where(dao.DevDeviceTag.Columns().DeviceId, deviceId).
			WhereNotIn(dao.DevDeviceTag.Columns().Id, tagIds).
			Unscoped().Delete()
		if err != nil {
			return
		}
	}
	for _, v := range add {
		newV := v
		if err = service.DevDeviceTag().Add(ctx, &newV); err != nil {
			return
		}
	}
	return
}
