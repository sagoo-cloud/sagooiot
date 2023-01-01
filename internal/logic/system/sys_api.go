package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysApi struct {
}

func sysApiNew() *sSysApi {
	return &sSysApi{}
}

func init() {
	service.RegisterSysApi(sysApiNew())
}

// GetInfoByIds 根据接口APIID数组获取接口信息
func (s *sSysApi) GetInfoByIds(ctx context.Context, ids []int) (data []*entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().IsDeleted: 0,
		dao.SysApi.Columns().Status:    1,
	}).WhereIn(dao.SysApi.Columns().Id, ids).Scan(&data)
	return
}

// GetApiByMenuId 根据ApiID获取接口信息
func (s *sSysApi) GetApiByMenuId(ctx context.Context, apiId int) (data []*entity.SysApi, err error) {
	var apiApi []*entity.SysMenuApi
	err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().MenuId:    apiId,
		dao.SysMenuApi.Columns().IsDeleted: 0,
	}).Scan(&apiApi)
	//获取接口ID数组
	if apiApi != nil {
		var ids []int
		for _, api := range apiApi {
			ids = append(ids, api.ApiId)
		}
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().IsDeleted: 0,
			dao.SysApi.Columns().Status:    1,
		}).WhereIn(dao.SysApi.Columns().Id, ids).Scan(&data)
	}
	return
}

// GetInfoById 根据ID获取API
func (s *sSysApi) GetInfoById(ctx context.Context, id int) (entity *entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Scan(&entity)
	return
}

// GetApiAll 获取所有接口
func (s *sSysApi) GetApiAll(ctx context.Context) (data []*entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().IsDeleted: 0,
		dao.SysApi.Columns().Status:    1,
		dao.SysApi.Columns().Types:     2,
	}).Scan(&data)
	return
}

// GetApiTree 获取Api数结构数据
func (s *sSysApi) GetApiTree(ctx context.Context, name string, address string, status int, types int) (out []*model.SysApiTreeOut, err error) {
	var e []*entity.SysApi
	m := dao.SysApi.Ctx(ctx)
	if name != "" {
		m = m.WhereLike(dao.SysApi.Columns().Name, "%"+name+"%")
	}
	if address != "" {
		m = m.WhereLike(dao.SysApi.Columns().Address, "%"+address+"%")
	}
	if status != -1 {
		m = m.Where(dao.SysApi.Columns().Status, status)
	}
	if types != -1 {
		m = m.Where(dao.SysApi.Columns().Types, types)
	}
	m = m.Where(dao.SysApi.Columns().IsDeleted, 0)

	err = m.OrderAsc(dao.SysApi.Columns().Sort).Scan(&e)

	if len(e) > 0 {
		out, err = GetApiTree(e)
		if err != nil {
			return
		}
	}
	return
}

// Add 添加Api列表
func (s *sSysApi) Add(ctx context.Context, input *model.AddApiInput) (err error) {
	if input.Types == 2 {
		if input.Address == "" || len(input.MenuIds) == 0 {
			err = gerror.New("参数错误")
			return
		}
		if input.ParentId == -1 {
			err = gerror.New("接口不能为根节点")
			return
		}
	} else {
		if input.Address != "" || len(input.MenuIds) > 0 {
			err = gerror.New("参数错误")
			return
		}
	}
	if input.ParentId != -1 {
		var parentApiInfo *entity.SysApi
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().Id: input.ParentId,
		}).Scan(&parentApiInfo)
		if parentApiInfo.IsDeleted != 0 {
			return gerror.New("上级节点已删除，无法新增")
		}
		if parentApiInfo.Status != 1 {
			return gerror.New("上级节点未启用，无法新增")
		}
		if parentApiInfo.Types != 1 {
			return gerror.New("上级节点不是分类，无法新增")
		}
	}
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		var apiInfo *entity.SysApi
		//根据名称查看是否存在
		apiInfo = checkApiName(ctx, input.Name, 0)
		if apiInfo != nil {
			return gerror.New("Api名字已存在,无法添加")
		}
		if input.Types == 2 {
			//根据名称查看是否存在
			apiInfo = checkApiAddress(ctx, input.Address, 0)
			if apiInfo != nil {
				return gerror.New("Api地址,无法添加")
			}
		}
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		apiInfo = new(entity.SysApi)
		if apiInfoErr := gconv.Scan(input, &apiInfo); apiInfoErr != nil {
			return
		}
		apiInfo.IsDeleted = 0
		apiInfo.CreateBy = uint(loginUserId)
		apiInfoId, err := dao.SysApi.Ctx(ctx).Data(apiInfo).InsertAndGetId()
		if err != nil {
			return err
		}
		//绑定菜单
		err = AddMenuApi(ctx, int(apiInfoId), input.MenuIds, loginUserId)
		return
	})
	return
}

// Detail Api列表详情
func (s *sSysApi) Detail(ctx context.Context, id int) (out *model.SysApiOut, err error) {
	var e *entity.SysApi
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Scan(&e)
	if e != nil {
		if err = gconv.Scan(e, &out); err != nil {
			return nil, err
		}
		menuApiInfo, _ := service.SysMenuApi().GetInfoByApiId(ctx, out.Id)
		var menuIds []int
		for _, menuApi := range menuApiInfo {
			menuIds = append(menuIds, menuApi.MenuId)
		}
		out.MenuIds = append(out.MenuIds, menuIds...)
	}
	return
}

func AddMenuApi(ctx context.Context, id int, menuIds []int, loginUserId int) (err error) {
	//添加菜单
	var sysMenuApis []*entity.SysMenuApi
	for _, menuId := range menuIds {
		var menuInfo *entity.SysMenu
		err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, menuId).Scan(&menuInfo)
		if menuInfo == nil {
			err = gerror.New("菜单ID错误")
			return
		}
		if menuInfo != nil && menuInfo.IsDeleted == 1 {
			err = gerror.New(menuInfo.Name + "已删除,无法绑定")
			return
		}
		if menuInfo != nil && menuInfo.Status == 0 {
			err = gerror.New(menuInfo.Name + "已禁用,无法绑定")
			return
		}

		//解除旧绑定关系
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 1,
			dao.SysMenuApi.Columns().DeletedBy: loginUserId,
		}).Where(dao.SysMenuApi.Columns().ApiId, id).Update()
		_, err = dao.SysMenuApi.Ctx(ctx).Where(dao.SysMenuApi.Columns().ApiId, id).Delete()

		var sysMenuApi = new(entity.SysMenuApi)
		sysMenuApi.MenuId = menuId
		sysMenuApi.ApiId = id
		sysMenuApi.IsDeleted = 0
		sysMenuApi.CreatedBy = uint(loginUserId)
		sysMenuApis = append(sysMenuApis, sysMenuApi)
	}
	if sysMenuApis != nil {
		//添加
		_, addErr := dao.SysMenuApi.Ctx(ctx).Data(sysMenuApis).Insert()
		if addErr != nil {
			err = gerror.New("添加失败")
			return
		}
	}
	return
}

// Edit 修改Api列表
func (s *sSysApi) Edit(ctx context.Context, input *model.EditApiInput) (err error) {
	if input.Types == 2 {
		if input.Address == "" || len(input.MenuIds) == 0 {
			err = gerror.New("参数错误")
			return
		}
		if input.ParentId == -1 {
			err = gerror.New("接口不能为根节点")
			return
		}
	} else {
		if input.Address != "" || len(input.MenuIds) > 0 {
			err = gerror.New("参数错误")
			return
		}
	}
	if input.ParentId != -1 {
		var parentApiInfo *entity.SysApi
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().Id: input.ParentId,
		}).Scan(&parentApiInfo)
		if parentApiInfo.IsDeleted != 0 {
			return gerror.New("上级节点已删除，无法新增")
		}
		if parentApiInfo.Status != 1 {
			return gerror.New("上级节点已启用，无法新增")
		}
		if parentApiInfo.Types != 1 {
			return gerror.New("上级节点不是分类，无法新增")
		}
	}
	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		var apiInfo, apiInfo2 *entity.SysApi
		//根据ID查看Api列表是否存在
		apiInfo = checkApiId(ctx, input.Id, apiInfo)
		if apiInfo == nil {
			return gerror.New("Api列表不存在")
		}
		apiInfo2 = checkApiName(ctx, input.Name, input.Id)
		if apiInfo2 != nil {
			return gerror.New("相同Api名称已存在,无法修改")
		}
		if input.Types == 2 {
			apiInfo2 = checkApiAddress(ctx, input.Address, input.Id)
			if apiInfo2 != nil {
				return gerror.New("Api地址已存在,无法修改")
			}
		}
		//获取当前登录用户ID
		loginUserId := service.Context().GetUserId(ctx)
		if apiInfoErr := gconv.Scan(input, &apiInfo); apiInfoErr != nil {
			return
		}
		apiInfo.UpdatedBy = uint(loginUserId)
		_, err = dao.SysApi.Ctx(ctx).Data(apiInfo).
			Where(dao.SysApi.Columns().Id, input.Id).Update()
		if err != nil {
			return gerror.New("修改失败")
		}
		//绑定菜单
		err = AddMenuApi(ctx, input.Id, input.MenuIds, loginUserId)
		return
	})

	return
}

// Del 根据ID删除Api列表信息
func (s *sSysApi) Del(ctx context.Context, Id int) (err error) {
	var apiColumn *entity.SysApi
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: Id,
	}).Scan(&apiColumn)
	if apiColumn == nil {
		return gerror.New("ID错误")
	}
	//判断是否为分类
	if apiColumn.Types == 1 {
		//判断是否存在子节点
		num, _ := dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().ParentId:  Id,
			dao.SysApi.Columns().IsDeleted: 0,
		}).Count()
		if num > 0 {
			return gerror.New("存在子节点，无法删除")
		}
	}
	loginUserId := service.Context().GetUserId(ctx)
	//获取当前时间
	time, err := gtime.StrToTimeFormat(gtime.Datetime(), "2006-01-02 15:04:05")
	if err != nil {
		return
	}
	//开启事务管理
	err = dao.SysApi.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		//更新Api列表信息
		_, err = dao.SysApi.Ctx(ctx).
			Data(g.Map{
				dao.SysApi.Columns().DeletedBy: uint(loginUserId),
				dao.SysApi.Columns().IsDeleted: 1,
				dao.SysApi.Columns().DeletedAt: time,
			}).Where(dao.SysApi.Columns().Id, Id).
			Update()
		//删除于菜单关系绑定
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 0,
			dao.SysMenuApi.Columns().DeletedBy: loginUserId,
			dao.SysMenuApi.Columns().DeletedAt: time,
		}).Where(dao.SysMenuApi.Columns().ApiId, Id).Update()
		return
	})
	return
}

// EditStatus 修改状态
func (s *sSysApi) EditStatus(ctx context.Context, id int, status int) (err error) {
	var apiInfo *entity.SysApi
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Scan(&apiInfo)
	if apiInfo == nil {
		return gerror.New("ID错误")
	}
	if apiInfo != nil && apiInfo.IsDeleted == 1 {
		return gerror.New("列表字段已删除,无法修改")
	}
	if apiInfo != nil && apiInfo.Status == status {
		return gerror.New("API已禁用或启用,无须重复修改")
	}
	loginUserId := service.Context().GetUserId(ctx)
	apiInfo.Status = status
	apiInfo.UpdatedBy = uint(loginUserId)

	_, err = dao.SysApi.Ctx(ctx).Data(apiInfo).Where(g.Map{
		dao.SysApi.Columns().Id: id,
	}).Update()
	return
}

// GetInfoByAddress 根据Address获取API
func (s *sSysApi) GetInfoByAddress(ctx context.Context, address string) (entity *entity.SysApi, err error) {
	err = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Address:   address,
		dao.SysApi.Columns().IsDeleted: 0,
		dao.SysApi.Columns().Status:    1,
	}).Scan(&entity)
	return
}

// 检查指定ID的数据是否存在
func checkApiId(ctx context.Context, Id int, apiColumn *entity.SysApi) *entity.SysApi {
	_ = dao.SysApi.Ctx(ctx).Where(g.Map{
		dao.SysApi.Columns().Id:        Id,
		dao.SysApi.Columns().IsDeleted: 0,
	}).Scan(&apiColumn)
	return apiColumn
}

// 检查相同Api名称的数据是否存在
func checkApiName(ctx context.Context, name string, tag int) *entity.SysApi {
	var apiInfo *entity.SysApi
	m := dao.SysApi.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysApi.Columns().Id, tag)
	}
	_ = m.Where(g.Map{
		dao.SysApi.Columns().Name:      name,
		dao.SysApi.Columns().IsDeleted: 0,
	}).Scan(&apiInfo)
	return apiInfo
}

// 检查相同Api地址的数据是否存在
func checkApiAddress(ctx context.Context, address string, tag int) *entity.SysApi {
	var apiInfo *entity.SysApi
	m := dao.SysApi.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysApi.Columns().Id, tag)
	}
	_ = m.Where(g.Map{
		dao.SysApi.Columns().Address:   address,
		dao.SysApi.Columns().IsDeleted: 0,
	}).Scan(&apiInfo)
	return apiInfo
}
