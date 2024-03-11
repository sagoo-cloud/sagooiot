package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sort"
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
	if err != nil {
		return
	}
	if data != nil && len(data) > 0 {
		err = cache.Instance().Set(ctx, consts.CacheSysMenu, data, 0)
		if err != nil {
			return
		}
	} else {
		_, err = cache.Instance().Remove(ctx, consts.CacheSysMenu)
	}
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

				var isExist = false
				for _, menuOut := range parentNodeOut {
					if menuOut.Id == parentNode.Id {
						isExist = true
						break
					}
				}
				if !isExist {
					parentNodeOut = append(parentNodeOut, parentNode)
				}
			} else {
				//查找根节点
				var parentMenu *entity.SysMenu
				parentMenu, err = FindMenuParentByChildrenId(ctx, int(v.ParentId))
				if err != nil {

				}
				if err = gconv.Scan(parentMenu, &parentNode); err != nil {
					return
				}
				var isExist = false
				for _, menuOut := range parentNodeOut {
					if menuOut.Id == int64(parentMenu.Id) {
						isExist = true
						break
					}
				}
				if !isExist {
					parentNodeOut = append(parentNodeOut, parentNode)
				}
			}
		}
		//对父节点进行排序
		sort.SliceStable(parentNodeOut, func(i, j int) bool {
			return parentNodeOut[i].Weigh > parentNodeOut[j].Weigh
		})
		data = MenuTree(parentNodeOut, menuInfo)
	}

	return
}

// FindMenuParentByChildrenId 根据子节点获取根节点
func FindMenuParentByChildrenId(ctx context.Context, parentId int) (out *entity.SysMenu, err error) {
	err = dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().Id:        parentId,
		dao.SysMenu.Columns().IsDeleted: 0,
	}).Scan(&out)
	if err != nil {
		return
	}
	if out.ParentId != -1 {
		return FindMenuParentByChildrenId(ctx, out.ParentId)
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
		//对子节点进行排序
		sort.SliceStable(v.Children, func(i, j int) bool {
			return v.Children[i].Weigh > v.Children[j].Weigh
		})
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
	_, err = dao.SysMenu.Ctx(ctx).Data(do.SysMenu{
		ParentId:   menu.ParentId,
		Name:       menu.Name,
		Title:      menu.Title,
		Icon:       menu.Icon,
		Condition:  menu.Condition,
		Remark:     menu.Remark,
		MenuType:   menu.MenuType,
		Weigh:      menu.Weigh,
		IsHide:     menu.IsHide,
		Path:       menu.Path,
		Component:  menu.Component,
		IsLink:     menu.IsLink,
		ModuleType: menu.ModuleType,
		ModelId:    menu.ModelId,
		IsIframe:   menu.IsIframe,
		IsCached:   menu.IsCached,
		Redirect:   menu.Redirect,
		IsAffix:    menu.IsAffix,
		LinkUrl:    menu.LinkUrl,
		Status:     menu.Status,
		IsDeleted:  menu.IsDeleted,
		CreatedBy:  menu.CreatedBy,
		CreatedAt:  menu.CreatedAt,
	}).Insert()
	if err != nil {
		return err
	}
	//获取所有的菜单
	_, err = s.GetAll(ctx)
	if err != nil {
		return
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
	//获取所有的菜单
	_, err = s.GetAll(ctx)
	if err != nil {
		return
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
	//查询是否有子节点
	num, _ := dao.SysMenu.Ctx(ctx).Where(g.Map{
		dao.SysMenu.Columns().ParentId:  menuId,
		dao.SysMenu.Columns().IsDeleted: 0,
	}).Count()
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.SysMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		//删除菜单
		_, err = dao.SysMenu.Ctx(ctx).Data(g.Map{
			dao.SysMenu.Columns().DeletedBy: uint(loginUserId),
			dao.SysMenu.Columns().DeletedAt: gtime.Now(),
			dao.SysMenu.Columns().IsDeleted: 1,
		}).Where(dao.SysMenu.Columns().Id, menuId).Update()

		//删除关联Api
		_, err = dao.SysMenuApi.Ctx(ctx).Data(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 1,
			dao.SysMenuApi.Columns().DeletedBy: uint(loginUserId),
			dao.SysMenuApi.Columns().DeletedAt: gtime.Now()}).Where(dao.SysMenuApi.Columns().MenuId, menuId).Update()
		if err != nil {
			return
		}

		//删除关联按钮
		_, err = dao.SysMenuButton.Ctx(ctx).Data(g.Map{
			dao.SysMenuButton.Columns().IsDeleted: 1,
			dao.SysMenuButton.Columns().DeletedBy: uint(loginUserId),
			dao.SysMenuButton.Columns().DeletedAt: gtime.Now()}).Where(dao.SysMenuButton.Columns().MenuId, menuId).Update()
		if err != nil {
			return
		}

		//删除关联列表
		_, err = dao.SysMenuColumn.Ctx(ctx).Data(g.Map{
			dao.SysMenuColumn.Columns().IsDeleted: 1,
			dao.SysMenuColumn.Columns().DeletedBy: uint(loginUserId),
			dao.SysMenuColumn.Columns().DeletedAt: gtime.Now()}).Where(dao.SysMenuColumn.Columns().MenuId, menuId).Update()
		if err != nil {
			return
		}

		return
	})
	//获取所有的菜单
	_, err = s.GetAll(ctx)
	if err != nil {
		return
	}
	return
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
	var tmpData *gvar.Var
	tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenu)
	if err != nil {
		return
	}

	var tmpMenuInfo []*entity.SysMenu

	var menuInfo []*entity.SysMenu
	//根据菜单ID数组获取菜单列表信息
	if tmpData.Val() != nil {
		if err = json.Unmarshal([]byte(tmpData.Val().(string)), &tmpMenuInfo); err != nil {
			return
		}
		for _, menuId := range menuIds {
			for _, menuTmp := range tmpMenuInfo {
				if menuId == int(menuTmp.Id) {
					menuInfo = append(menuInfo, menuTmp)
					continue
				}
			}
		}
	}
	if menuInfo != nil && len(menuInfo) >= 0 {
		data = menuInfo
		return
	}
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
