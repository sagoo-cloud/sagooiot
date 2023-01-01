package system

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysMenuColumn struct {
}

func sysMenuColumnNew() *sSysMenuColumn {
	return &sSysMenuColumn{}
}

func init() {
	service.RegisterSysMenuColumn(sysMenuColumnNew())
}

// GetList 获取全部菜单列表数据
func (s *sSysMenuColumn) GetList(ctx context.Context, input *model.MenuColumnDoInput) (data []model.UserMenuColumnRes, err error) {
	var menuColumn []model.UserMenuColumnRes
	menuColumn, err = s.GetData(ctx, input, menuColumn)
	return menuColumn, err
}

// GetData 执行获取数据操作
func (s *sSysMenuColumn) GetData(ctx context.Context, input *model.MenuColumnDoInput, menuColumn []model.UserMenuColumnRes) (data []model.UserMenuColumnRes, err error) {
	m := dao.SysMenuColumn.Ctx(ctx)
	//模糊查询菜单列表名称
	if input.Name != "" {
		m = m.WhereLike(dao.SysMenuColumn.Columns().Name, "%"+input.Name+"%")
	}
	//查询菜单列表菜单ID关联数据
	if input.MenuId != "" {
		m = m.Where(dao.SysMenuColumn.Columns().MenuId, input.MenuId)
	}
	//查询菜单列表上级
	if input.ParentId != "" {
		m = m.Where(dao.SysMenuColumn.Columns().ParentId, input.ParentId)
	}
	//模糊查询菜单列表状态
	if input.Status != -1 {
		m = m.Where(dao.SysMenuColumn.Columns().Status, input.Status)
	}
	err = m.Where(g.Map{
		dao.SysMenuColumn.Columns().MenuId:    input.MenuId,
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).
		Scan(&menuColumn)
	return menuColumn, err
}

// Add 添加菜单列表
func (s *sSysMenuColumn) Add(ctx context.Context, input *model.AddMenuColumnInput) (err error) {
	var menuColumn *entity.SysMenuColumn
	//根据名称查看是否存在
	menuColumn = checkMenuColumnName(ctx, input.MenuId, input.Name, menuColumn)
	if menuColumn != nil {
		return gerror.New("菜单列表已存在,无法添加")
	}

	menuColumn = checkMenuColumnCode(ctx, input.MenuId, input.Code, menuColumn)
	if menuColumn != nil {
		return gerror.New("菜单列表CODE已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	menuColumn = new(entity.SysMenuColumn)
	if err := gconv.Scan(input, &menuColumn); err != nil {
		return err
	}
	menuColumn.IsDeleted = 0
	menuColumn.CreatedBy = uint(loginUserId)
	_, err = dao.SysMenuColumn.Ctx(ctx).Data(menuColumn).Insert()
	if err != nil {
		return err
	}
	return
}

// Detail 菜单列表详情
func (s *sSysMenuColumn) Detail(ctx context.Context, Id int64) (entity *entity.SysMenuColumn, err error) {
	_ = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().Id: Id,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("ID错误")
	}
	return
}

// Edit 修改菜单列表
func (s *sSysMenuColumn) Edit(ctx context.Context, input *model.EditMenuColumnInput) (err error) {
	var menuColumn, menuColumn2 *entity.SysMenuColumn
	//根据ID查看菜单列表是否存在
	menuColumn = checkMenuColumnId(ctx, input.Id, menuColumn)
	if menuColumn == nil {
		return gerror.New("菜单列表不存在")
	}
	//根据名称查看是否存在
	menuColumn2 = checkMenuColumnName(ctx, input.MenuId, input.Name, menuColumn)
	if menuColumn2 != nil && int(menuColumn2.Id) != input.Id {
		return gerror.New("菜单列表已存在,无法添加")
	}

	menuColumn2 = checkMenuColumnCode(ctx, input.MenuId, input.Code, menuColumn)
	if menuColumn2 != nil && int(menuColumn2.Id) != input.Id {
		return gerror.New("菜单列表CODE已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &menuColumn); err != nil {
		return err
	}
	menuColumn.UpdatedBy = int(uint(loginUserId))
	_, err = dao.SysMenuColumn.Ctx(ctx).Data(menuColumn).
		Where(dao.SysMenuColumn.Columns().Id, input.Id).Update()
	if err != nil {
		return gerror.New("修改失败")
	}
	return
}

// Del 根据ID删除菜单列表信息
func (s *sSysMenuColumn) Del(ctx context.Context, Id int64) (err error) {
	var menuColumn *entity.SysMenuColumn
	_ = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().Id: Id,
	}).Scan(&menuColumn)
	if menuColumn == nil {
		return gerror.New("ID错误")
	}
	//查询是否存在下级
	num, err := dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().ParentId:  Id,
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}
	loginUserId := service.Context().GetUserId(ctx)
	//更新菜单列表信息
	_, err = dao.SysMenuColumn.Ctx(ctx).
		Data(g.Map{
			dao.SysMenuColumn.Columns().DeletedBy: uint(loginUserId),
			dao.SysMenuColumn.Columns().IsDeleted: 1,
		}).Where(dao.SysMenuColumn.Columns().Id, Id).
		Update()
	//删除菜单列表信息
	_, err = dao.SysMenuColumn.Ctx(ctx).Where(dao.SysMenuColumn.Columns().Id, Id).Delete()
	return
}

// EditStatus 修改状态
func (s *sSysMenuColumn) EditStatus(ctx context.Context, id int, menuId int, status int) (err error) {
	var menuColum *entity.SysMenuColumn
	_ = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().Id: id,
	}).Scan(&menuColum)
	if menuColum == nil {
		return gerror.New("ID错误")
	}
	if menuColum.MenuId != menuId {
		return gerror.New("列表字段不属于当前菜单,无法修改")
	}
	if menuColum != nil && menuColum.IsDeleted == 1 {
		return gerror.New("列表字段已删除,无法修改")
	}
	if menuColum != nil && menuColum.Status == status {
		return gerror.New("列表已禁用或启用,无须重复修改")
	}
	loginUserId := service.Context().GetUserId(ctx)
	menuColum.Status = status
	menuColum.UpdatedBy = loginUserId

	_, err = dao.SysMenuColumn.Ctx(ctx).Data(menuColum).Where(g.Map{
		dao.SysMenuColumn.Columns().Id: id,
	}).Update()
	return
}

// 检查指定ID的数据是否存在
func checkMenuColumnId(ctx context.Context, Id int, menuColumn *entity.SysMenuColumn) *entity.SysMenuColumn {
	_ = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().Id:        Id,
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).Scan(&menuColumn)
	return menuColumn
}

// 检查相同菜单按钮名称的数据是否存在
func checkMenuColumnName(ctx context.Context, menu int, name string, menuColumn *entity.SysMenuColumn) *entity.SysMenuColumn {
	m := dao.SysMenuColumn.Ctx(ctx)
	_ = m.Where(g.Map{
		dao.SysMenuColumn.Columns().Name:      name,
		dao.SysMenuColumn.Columns().MenuId:    menu,
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).Scan(&menuColumn)
	return menuColumn
}

func checkMenuColumnCode(ctx context.Context, menu int, code string, menuColumn *entity.SysMenuColumn) *entity.SysMenuColumn {
	m := dao.SysMenuColumn.Ctx(ctx)
	_ = m.Where(g.Map{
		dao.SysMenuColumn.Columns().Code:      code,
		dao.SysMenuColumn.Columns().MenuId:    menu,
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).Scan(&menuColumn)
	return menuColumn
}

// GetInfoByColumnIds 根据列表ID数组获取菜单信息
func (s *sSysMenuColumn) GetInfoByColumnIds(ctx context.Context, ids []int) (data []*entity.SysMenuColumn, err error) {
	err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysMenuColumn.Columns().Id, ids).Scan(&data)
	return
}

// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
func (s *sSysMenuColumn) GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuColumn, err error) {
	err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).WhereIn(dao.SysMenuColumn.Columns().MenuId, menuIds).Scan(&data)
	return
}
