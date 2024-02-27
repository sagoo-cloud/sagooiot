package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"

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
func (s *sSysMenuColumn) GetList(ctx context.Context, input *model.MenuColumnDoInput) (data []*model.UserMenuColumnOut, err error) {
	menuColumnOut, err := s.GetData(ctx, input)
	if err != nil {
		return
	}
	var parentNodeOut []*model.UserMenuColumnOut
	if menuColumnOut != nil {
		//获取所有的根节点
		for _, v := range menuColumnOut {
			var parentNode *model.UserMenuColumnOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeOut = append(parentNodeOut, parentNode)
			}
		}
		data = ColumnTree(parentNodeOut, menuColumnOut)
	}
	return
}

// ColumnTree MenuColumnTree 生成菜单列表树结构
func ColumnTree(parentNodeOut []*model.UserMenuColumnOut, data []model.UserMenuColumnOut) (dataTree []*model.UserMenuColumnOut) {
	//循环所有一级菜单
	for k, v := range parentNodeOut {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.UserMenuColumnOut
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeOut[k].Children = append(parentNodeOut[k].Children, node)
			}
		}
		ColumnTree(v.Children, data)
	}
	return parentNodeOut
}

// GetData 执行获取数据操作
func (s *sSysMenuColumn) GetData(ctx context.Context, input *model.MenuColumnDoInput) (data []model.UserMenuColumnOut, err error) {
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
	}).Scan(&data)
	return
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
	_, err = dao.SysMenuColumn.Ctx(ctx).Data(do.SysMenuColumn{
		ParentId:    menuColumn.ParentId,
		MenuId:      menuColumn.MenuId,
		Name:        menuColumn.Name,
		Code:        menuColumn.Code,
		Description: menuColumn.Description,
		Status:      menuColumn.Status,
		IsDeleted:   menuColumn.IsDeleted,
		CreatedBy:   menuColumn.CreatedBy,
		CreatedAt:   gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuColumn.MenuId)
	if err != nil {
		return
	}
	//所有的菜单列表
	_, err = s.GetAll(ctx)
	if err != nil {
		return
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
	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuColumn.MenuId)
	if err != nil {
		return
	}
	//所有的菜单列表
	_, err = s.GetAll(ctx)
	if err != nil {
		return
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
			dao.SysMenuColumn.Columns().DeletedAt: gtime.Now(),
			dao.SysMenuColumn.Columns().IsDeleted: 1,
		}).Where(dao.SysMenuColumn.Columns().Id, Id).
		Update()

	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuColumn.MenuId)
	if err != nil {
		return
	}
	//所有的菜单列表
	_, err = s.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

// EditStatus 修改状态
func (s *sSysMenuColumn) EditStatus(ctx context.Context, id int, menuId int, status int) (err error) {
	var menuColumn *entity.SysMenuColumn
	_ = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().Id: id,
	}).Scan(&menuColumn)
	if menuColumn == nil {
		return gerror.New("ID错误")
	}
	if menuColumn.MenuId != menuId {
		return gerror.New("列表字段不属于当前菜单,无法修改")
	}
	if menuColumn != nil && menuColumn.IsDeleted == 1 {
		return gerror.New("列表字段已删除,无法修改")
	}
	if menuColumn != nil && menuColumn.Status == status {
		return gerror.New("列表已禁用或启用,无须重复修改")
	}
	loginUserId := service.Context().GetUserId(ctx)
	menuColumn.Status = status
	menuColumn.UpdatedBy = loginUserId

	_, err = dao.SysMenuColumn.Ctx(ctx).Data(menuColumn).Where(g.Map{
		dao.SysMenuColumn.Columns().Id: id,
	}).Update()
	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuColumn.MenuId)
	if err != nil {
		return
	}
	//所有的菜单列表
	_, err = s.GetAll(ctx)
	if err != nil {
		return
	}

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
	var tmpData *gvar.Var
	tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenuColumn)
	if err != nil {
		return
	}

	var tmpSysMenuColumn []*entity.SysMenuColumn

	var menuColumnInfo []*entity.SysMenuColumn
	//根据菜单ID数组获取菜单列表信息
	if tmpData.Val() != nil {
		if err = json.Unmarshal([]byte(tmpData.Val().(string)), &tmpSysMenuColumn); err != nil {
			return
		}
		for _, id := range ids {
			for _, tmp := range tmpSysMenuColumn {
				if id == int(tmp.Id) {
					menuColumnInfo = append(menuColumnInfo, tmp)
					continue
				}
			}
		}
	}
	if menuColumnInfo != nil && len(menuColumnInfo) >= 0 {
		data = menuColumnInfo
		return
	}

	err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().IsDeleted: 0,
		dao.SysMenuColumn.Columns().Status:    1,
	}).WhereIn(dao.SysMenuColumn.Columns().Id, ids).Scan(&data)
	return
}

// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
func (s *sSysMenuColumn) GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuColumn, err error) {
	//获取缓存菜单按钮信息
	for _, v := range menuIds {
		var tmpData *gvar.Var
		tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenuColumn+"_"+gconv.String(v))
		if err != nil {
			return
		}
		if tmpData.Val() != nil {
			var sysMenuColumn []*entity.SysMenuColumn
			err = json.Unmarshal([]byte(tmpData.Val().(string)), &sysMenuColumn)
			data = append(data, sysMenuColumn...)
		}
	}
	if data == nil || len(data) == 0 {
		err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
			dao.SysMenuColumn.Columns().IsDeleted: 0,
			dao.SysMenuColumn.Columns().Status:    1,
		}).WhereIn(dao.SysMenuColumn.Columns().MenuId, menuIds).Scan(&data)
	}
	return
}

// GetInfoByMenuId 根据菜单ID获取菜单信息
func (s *sSysMenuColumn) GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuColumn, err error) {
	err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().IsDeleted: 0,
		dao.SysMenuColumn.Columns().Status:    1,
		dao.SysMenuColumn.Columns().MenuId:    menuId,
	}).Scan(&data)
	if err != nil {
		return
	}
	if data != nil && len(data) > 0 {
		err = cache.Instance().Set(ctx, consts.CacheSysMenuColumn+"_"+gconv.String(menuId), data, 0)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = cache.Instance().Remove(ctx, consts.CacheSysMenuColumn+"_"+gconv.String(menuId))
	}
	return
}

// GetAll 获取所有的列表信息
func (s *sSysMenuColumn) GetAll(ctx context.Context) (data []*entity.SysMenuColumn, err error) {
	err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().IsDeleted: 0,
		dao.SysMenuColumn.Columns().Status:    1,
	}).Scan(&data)
	if err != nil {
		return
	}
	if data != nil && len(data) > 0 {
		err = cache.Instance().Set(ctx, consts.CacheSysMenuColumn, data, 0)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = cache.Instance().Remove(ctx, consts.CacheSysMenuColumn)
	}
	return
}
