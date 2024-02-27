package product

import (
	"context"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

type sDevDeviceTree struct{}

func init() {
	service.RegisterDevDeviceTree(devDeviceTreeNew())
}

func devDeviceTreeNew() *sDevDeviceTree {
	return &sDevDeviceTree{}
}

// List 设备树列表
func (s *sDevDeviceTree) List(ctx context.Context) (out []*model.DeviceTreeListOutput, err error) {

	var list []*model.DeviceTree
	if err = dao.DevDeviceTree.Ctx(ctx).OrderAsc(dao.DevDeviceTree.Columns().Id).Scan(&list); err != nil || len(list) == 0 {
		return nil, err
	}

	m := dao.DevDeviceTreeInfo.Ctx(ctx)

	var infoList []*entity.DevDeviceTreeInfo
	if err = m.Scan(&infoList); err != nil || len(infoList) == 0 {
		return nil, err
	}

	infoMap := make(map[int]string, len(infoList))

	var deviceTreeList []*model.DeviceTree

	for _, v := range infoList {
		infoMap[v.Id] = v.Name
		for _, l := range list {
			if v.Id == l.InfoId {
				deviceTreeList = append(deviceTreeList, l)
				break
			}
		}
	}
	for i, v := range deviceTreeList {
		deviceTreeList[i].Name = infoMap[v.InfoId]
	}

	return tree(deviceTreeList, 0), nil
}

func tree(all []*model.DeviceTree, pid int) (rs []*model.DeviceTreeListOutput) {
	for _, v := range all {
		if v.ParentInfoId == pid {
			var out *model.DeviceTreeListOutput
			if err := gconv.Scan(v, &out); err != nil {
				return
			}
			out.Children = tree(all, v.InfoId)
			rs = append(rs, out)
		}
	}
	return
}

// Change 更换上下级
func (s *sDevDeviceTree) Change(ctx context.Context, infoId, parentInfoId int) error {
	_, err := dao.DevDeviceTree.Ctx(ctx).Where(dao.DevDeviceTree.Columns().InfoId, infoId).Update(g.Map{
		dao.DevDeviceTree.Columns().ParentInfoId: parentInfoId,
	})
	return err
}

// Detail 信息详情
func (s *sDevDeviceTree) Detail(ctx context.Context, infoId int) (out *model.DetailDeviceTreeInfoOutput, err error) {
	if err = dao.DevDeviceTreeInfo.Ctx(ctx).Where(dao.DevDeviceTreeInfo.Columns().Id, infoId).Scan(&out); err != nil || out == nil {
		return
	}
	rs, err := dao.DevDeviceTree.Ctx(ctx).Fields(dao.DevDeviceTree.Columns().ParentInfoId).Where(dao.DevDeviceTree.Columns().InfoId, infoId).Value()
	if err != nil {
		return
	}
	out.ParentId = rs.Int()
	return
}

// check 检查设备是否被绑定：true=可绑定
func (s *sDevDeviceTree) check(ctx context.Context, deviceKey string, infoId int) (b bool, err error) {
	if deviceKey == "" {
		b = true
		return
	}
	m := dao.DevDeviceTreeInfo.Ctx(ctx).Where(dao.DevDeviceTreeInfo.Columns().DeviceKey, deviceKey)
	if infoId > 0 {
		m = m.WhereNot(dao.DevDeviceTreeInfo.Columns().Id, infoId)
	}
	rs, err := m.Count()
	if err != nil {
		return
	}
	return rs == 0, nil
}

// Add 添加设备树基本信息
func (s *sDevDeviceTree) Add(ctx context.Context, in *model.AddDeviceTreeInfoInput) error {
	// 获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	if in.DeviceKey != "" {
		_, err := service.DevDevice().Get(ctx, in.DeviceKey)
		if err != nil {
			return err
		}
	}

	b, err := s.check(ctx, in.DeviceKey, 0)
	if err != nil {
		return err
	}
	if !b {
		return gerror.New("该设备已被绑定")
	}

	err = dao.DevDeviceTreeInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := dao.DevDeviceTreeInfo.Ctx(ctx).Data(do.DevDeviceTreeInfo{
			Code:      guid.S(),
			DeptId:    service.Context().GetUserDeptId(ctx),
			Name:      in.Name,
			Address:   in.Address,
			Lng:       in.Lng,
			Lat:       in.Lat,
			Contact:   in.Contact,
			Phone:     in.Phone,
			StartDate: in.StartDate,
			EndDate:   in.EndDate,
			Image:     in.Image,
			DeviceKey: in.DeviceKey,
			Duration:  in.Duration,
			TimeUnit:  in.TimeUnit,
			Template:  in.Template,
			Category:  in.Category,
			CreatedBy: uint(loginUserId),
			CreatedAt: gtime.Now(),
		}).Insert()
		if err != nil {
			return err
		}

		infoId, err := service.Sequences().GetSequences(ctx, result, dao.DevDeviceTreeInfo.Table(), dao.DevDeviceTreeInfo.Columns().Id)
		if err != nil {
			return err
		}

		_, err = dao.DevDeviceTree.Ctx(ctx).Data(do.DevDeviceTree{
			InfoId:       infoId,
			ParentInfoId: in.ParentId,
		}).Insert()

		return err
	})
	return err
}

// Edit 修改设备树基本信息
func (s *sDevDeviceTree) Edit(ctx context.Context, in *model.EditDeviceTreeInfoInput) error {
	// 获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	if in.DeviceKey != "" {
		_, err := service.DevDevice().Get(ctx, in.DeviceKey)
		if err != nil {
			return err
		}
	}

	b, err := s.check(ctx, in.DeviceKey, in.Id)
	if err != nil {
		return err
	}
	if !b {
		return gerror.New("该设备已被绑定")
	}

	var param *do.DevDeviceTreeInfo
	if err := gconv.Scan(in, &param); err != nil {
		return err
	}
	param.UpdatedBy = uint(loginUserId)

	err = dao.DevDeviceTreeInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.DevDeviceTreeInfo.Ctx(ctx).Where(dao.DevDeviceTreeInfo.Columns().Id, in.Id).Update(param)
		if err != nil {
			return err
		}
		relation := g.Map{
			dao.DevDeviceTree.Columns().ParentInfoId: in.ParentId,
		}
		_, err = dao.DevDeviceTree.Ctx(ctx).Where(dao.DevDeviceTree.Columns().InfoId, in.Id).Update(relation)
		return err
	})
	return err
}

// Del 删除设备树基本信息
func (s *sDevDeviceTree) Del(ctx context.Context, infoId int) error {
	// 获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	n, err := dao.DevDeviceTree.Ctx(ctx).Where(dao.DevDeviceTree.Columns().ParentInfoId, infoId).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.New("请先处理该信息的子集关系")
	}

	err = dao.DevDeviceTreeInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.DevDeviceTreeInfo.Ctx(ctx).
			Data(do.DevDeviceTreeInfo{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DevDeviceTreeInfo.Columns().Id, infoId).
			Unscoped().
			Update()
		if err != nil {
			return err
		}

		_, err = dao.DevDeviceTree.Ctx(ctx).Where(dao.DevDeviceTree.Columns().InfoId, infoId).Delete()
		return err
	})
	return err
}
