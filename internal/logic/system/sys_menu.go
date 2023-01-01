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

type sSysMenu struct {
}

func sysMenuNew() *sSysMenu {
	return &sSysMenu{}
}

func init() {
	service.RegisterSysMenu(sysMenuNew())
}

// GetAll 获取全部菜单数据
func (s *sSysMenu) GetAll(ctx context.Context) (data []*entity.SysMenu, err error) {
	err = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().Status:    1,
		dao.SysMenu.Columns().IsDeleted: 0,
	}).OrderDesc(dao.SysMenu.Columns().Weigh).Scan(&data)
	return
}

// GetTree 获取菜单数据
func (s *sSysMenu) GetTree(ctx context.Context, title string, status int) (data []*model.SysMenuOut, err error) {
	menuInfo, err := s.GetData(ctx, title, status)
	var parentNodeOut []*model.SysMenuOut
	if menuInfo != nil {
		//获取所有的根节点
		for _, v := range menuInfo {
			var parentNode *model.SysMenuOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeOut = append(parentNodeOut, parentNode)
			}
		}
		data = MenuTree(parentNodeOut, menuInfo)
	}

	return
}

// MenuTree 生成树结构
func MenuTree(parentNodeOut []*model.SysMenuOut, data []*model.SysMenuOut) (dataTree []*model.SysMenuOut) {
	//循环所有一级菜单
	for k, v := range parentNodeOut {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.SysMenuOut
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeOut[k].Children = append(parentNodeOut[k].Children, node)
			}
		}
		MenuTree(v.Children, data)
	}
	return parentNodeOut
}

// Add 添加菜单
func (s *sSysMenu) Add(ctx context.Context, input *model.AddMenuInput) (err error) {
	var menu *entity.SysMenu
	//根据名称查看角色是否存在
	menu = checkMenuName(ctx, input.Name, menu, 0)
	if menu != nil {
		return gerror.New("菜单已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	menu = new(entity.SysMenu)
	if err := gconv.Scan(input, &menu); err != nil {
		return err
	}
	menu.IsDeleted = 0
	menu.CreatedBy = uint(loginUserId)
	_, err = dao.SysMenu.Ctx(ctx).Data(menu).Insert()
	if err != nil {
		return err
	}
	return
}

// Detail 菜单详情
func (s *sSysMenu) Detail(ctx context.Context, menuId int64) (entity *entity.SysMenu, err error) {
	_ = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().Id: menuId,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("ID错误")
	}
	return
}

// Edit 修改菜单
func (s *sSysMenu) Edit(ctx context.Context, input *model.EditMenuInput) (err error) {
	var menu, menu2 *entity.SysMenu
	//根据ID查看菜单是否存在
	menu = checkMenuId(ctx, input.Id, menu)
	if menu == nil {
		return gerror.New("菜单不存在")
	}
	menu2 = checkMenuName(ctx, input.Name, menu2, input.Id)
	if menu2 != nil {
		return gerror.New("相同菜单已存在,无法修改")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &menu); err != nil {
		return err
	}
	menu.UpdatedBy = int(uint(loginUserId))
	_, err = dao.SysMenu.Ctx(ctx).Data(menu).
		Where(dao.SysMenu.Columns().Id, input.Id).Update()
	if err != nil {
		return gerror.New("修改失败")
	}
	return
}

// 检查相同菜单名称的数据是否存在
func checkMenuName(ctx context.Context, menuName string, menu *entity.SysMenu, tag int64) *entity.SysMenu {
	m := dao.SysMenu.Ctx(ctx)
	if tag > 0 {
		m = m.WhereNot(dao.SysMenu.Columns().Id, tag)
	}
	_ = m.Where(g.Map{
		dao.SysMenu.Columns().Name:      menuName,
		dao.SysMenu.Columns().IsDeleted: 0,
	}).Scan(&menu)
	return menu
}

// Del 根据ID删除菜单信息
func (s *sSysMenu) Del(ctx context.Context, menuId int64) (err error) {
	var menu *entity.SysMenu
	_ = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().Id: menuId,
	}).Scan(&menu)
	if menu == nil {
		return gerror.New("ID错误")
	}
	errorString := checkMenuJoin(ctx, menuId)
	if errorString != "" {
		return gerror.New(errorString)
	}
	loginUserId := service.Context().GetUserId(ctx)
	_, err = dao.SysMenu.Ctx(ctx).Data(g.Map{
		dao.SysMenu.Columns().DeletedBy: uint(loginUserId),
		dao.SysMenu.Columns().IsDeleted: 1,
	}).Where(dao.SysMenu.Columns().Id, menuId).Update()
	//删除菜单信息
	_, err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, menuId).Delete()
	return
}

// 查询关联数据是否存在
func checkMenuJoin(ctx context.Context, menuId int64) string {
	num := 0
	//查询是否有子节点
	num, _ = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().ParentId:  menuId,
		dao.SysMenu.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return "请先删除子节点!"
	}
	//查询关联Api
	num, _ = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
		dao.SysMenuApi.Columns().MenuId:    menuId,
		dao.SysMenuApi.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return "存在菜单API关联!"
	}
	//查询关联列表
	num, _ = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().MenuId:    menuId,
		dao.SysMenuButton.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return "存在菜单列表关联!"
	}
	//查询关联按钮
	num, _ = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
		dao.SysMenuColumn.Columns().MenuId:    menuId,
		dao.SysMenuColumn.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return "存在菜单按钮关联!"
	}
	return ""
}

// GetData 执行获取数据操作
func (s *sSysMenu) GetData(ctx context.Context, title string, status int) (data []*model.SysMenuOut, err error) {
	m := dao.SysMenu.Ctx(ctx)
	//模糊查询菜单标题名称
	if title != "" {
		m = m.WhereLike(dao.SysMenu.Columns().Title, "%"+title+"%")
	}
	if status != -1 {
		m = m.Where(dao.SysMenu.Columns().Status, status)
	}
	err = m.Where(dao.SysMenu.Columns().IsDeleted, 0).
		OrderDesc(dao.SysMenu.Columns().Weigh).
		Scan(&data)
	if err != nil {
		return
	}
	return
}

// 检查指定ID的数据是否存在
func checkMenuId(ctx context.Context, MenuId int64, menu *entity.SysMenu) *entity.SysMenu {
	_ = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().Id:        MenuId,
		dao.SysMenu.Columns().IsDeleted: 0,
	}).Scan(&menu)
	return menu
}

// GetInfoByMenuIds 根据菜单ID数组获取菜单信息
func (s *sSysMenu) GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenu, err error) {
	err = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().IsDeleted: 0,
		dao.SysMenu.Columns().Status:    1,
	}).WhereIn(dao.SysMenu.Columns().Id, menuIds).OrderDesc(dao.SysMenu.Columns().Weigh).Scan(&data)
	return
}

func (s *sSysMenu) GetInfoById(ctx context.Context, id int) (data *entity.SysMenu, err error) {
	err = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().IsDeleted: 0,
		dao.SysMenu.Columns().Id:        id,
	}).Scan(&data)
	return
}
