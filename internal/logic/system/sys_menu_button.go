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

type sSysMenuButton struct {
}

func sysMenuButtonNew() *sSysMenuButton {
	return &sSysMenuButton{}
}

func init() {
	service.RegisterSysMenuButton(sysMenuButtonNew())
}

// GetList 获取全部菜单按钮数据
func (s *sSysMenuButton) GetList(ctx context.Context, status int, name string, menuId int) (data []*model.UserMenuButtonOut, err error) {
	var menuButton []model.UserMenuButtonOut
	menuButton, err = s.GetData(ctx, status, name, menuId)

	var parentNodeOut []*model.UserMenuButtonOut
	if menuButton != nil {
		//获取所有的根节点
		for _, v := range menuButton {
			var parentNode *model.UserMenuButtonOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeOut = append(parentNodeOut, parentNode)
			}
		}
	}
	data = ButtonTree(parentNodeOut, menuButton)
	return
}

// ButtonTree MenuButtonTree 生成菜单按钮树结构
func ButtonTree(parentNodeOut []*model.UserMenuButtonOut, data []model.UserMenuButtonOut) (dataTree []*model.UserMenuButtonOut) {
	//循环所有一级菜单
	for k, v := range parentNodeOut {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.UserMenuButtonOut
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeOut[k].Children = append(parentNodeOut[k].Children, node)
			}
		}
		ButtonTree(v.Children, data)
	}
	return parentNodeOut
}

// GetData 执行获取数据操作
func (s *sSysMenuButton) GetData(ctx context.Context, status int, name string, menuId int) (data []model.UserMenuButtonOut, err error) {
	m := dao.SysMenuButton.Ctx(ctx)

	if status != -1 {
		m = m.Where(dao.SysMenuButton.Columns().Status, status)
	}
	//模糊查询菜单按钮名称
	if name != "" {
		m = m.WhereLike(dao.SysMenuButton.Columns().Name, "%"+name+"%")
	}
	err = m.Where(g.Map{
		dao.SysMenuButton.Columns().IsDeleted: 0,
		dao.SysMenuButton.Columns().MenuId:    menuId,
	}).Scan(&data)
	return
}

// Add 添加菜单按钮
func (s *sSysMenuButton) Add(ctx context.Context, input *model.AddMenuButtonInput) (err error) {
	var menuButton *entity.SysMenuButton
	//根据名称查看角色是否存在
	menuButton = checkMenuButtonName(ctx, input.MenuId, input.Name, menuButton)
	if menuButton != nil {
		return gerror.New("菜单按钮已存在,无法添加")
	}
	menuButton = checkMenuButtonType(ctx, input.MenuId, input.Types, menuButton)
	if menuButton != nil {
		return gerror.New("菜单按钮类型已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	menuButton = new(entity.SysMenuButton)
	if err := gconv.Scan(input, &menuButton); err != nil {
		return err
	}
	menuButton.IsDeleted = 0
	menuButton.CreatedBy = uint(loginUserId)
	_, err = dao.SysMenuButton.Ctx(ctx).Data(do.SysMenuButton{
		ParentId:    menuButton.ParentId,
		MenuId:      menuButton.MenuId,
		Name:        menuButton.Name,
		Types:       menuButton.Types,
		Description: menuButton.Description,
		Status:      menuButton.Status,
		IsDeleted:   menuButton.IsDeleted,
		CreatedBy:   menuButton.CreatedBy,
		CreatedAt:   gtime.Now(),
	}).Insert()
	if err != nil {
		return err
	}
	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuButton.MenuId)
	if err != nil {
		return
	}
	//所有的菜单按钮
	_, err = s.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

// Detail 菜单按钮详情
func (s *sSysMenuButton) Detail(ctx context.Context, Id int64) (entity *entity.SysMenuButton, err error) {
	_ = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().Id: Id,
	}).Scan(&entity)
	if entity == nil {
		return nil, gerror.New("ID错误")
	}
	return
}

// Edit 修改菜单按钮
func (s *sSysMenuButton) Edit(ctx context.Context, input *model.EditMenuButtonInput) (err error) {
	var menuButton, menuButton2 *entity.SysMenuButton
	//根据ID查看菜单按钮是否存在
	menuButton = checkMenuButtonId(ctx, input.Id, menuButton)
	if menuButton == nil {
		return gerror.New("菜单按钮不存在")
	}
	menuButton2 = checkMenuButtonName(ctx, input.MenuId, input.Name, menuButton)
	if menuButton2 != nil && int(menuButton2.Id) != input.Id {
		return gerror.New("菜单按钮已存在,无法添加")
	}
	menuButton2 = checkMenuButtonType(ctx, input.MenuId, input.Types, menuButton)
	if menuButton2 != nil && int(menuButton2.Id) != input.Id {
		return gerror.New("菜单按钮类型已存在,无法添加")
	}
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if err := gconv.Scan(input, &menuButton); err != nil {
		return err
	}
	menuButton.UpdatedBy = int(uint(loginUserId))
	_, err = dao.SysMenuButton.Ctx(ctx).Data(menuButton).
		Where(dao.SysMenuButton.Columns().Id, input.Id).Update()
	if err != nil {
		return gerror.New("修改失败")
	}
	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuButton.MenuId)
	if err != nil {
		return
	}
	//所有的菜单按钮
	_, err = s.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

// Del 根据ID删除菜单按钮信息
func (s *sSysMenuButton) Del(ctx context.Context, id int64) (err error) {
	var menuButton *entity.SysMenuButton
	_ = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().Id: id,
	}).Scan(&menuButton)
	if menuButton == nil {
		return gerror.New("ID错误")
	}
	//查询是否存在下级
	num, err := dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().ParentId:  id,
		dao.SysMenuButton.Columns().IsDeleted: 0,
	}).Count()
	if err != nil {
		return err
	}
	if num > 0 {
		return gerror.New("请先删除子节点!")
	}
	loginUserId := service.Context().GetUserId(ctx)
	//更新菜单按钮信息
	_, err = dao.SysMenuButton.Ctx(ctx).
		Data(g.Map{
			dao.SysMenuButton.Columns().DeletedBy: uint(loginUserId),
			dao.SysMenuButton.Columns().DeletedAt: gtime.Now(),
			dao.SysMenuButton.Columns().IsDeleted: 1,
		}).Where(dao.SysMenuButton.Columns().Id, id).
		Update()

	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuButton.MenuId)
	if err != nil {
		return
	}
	//所有的菜单按钮
	_, err = s.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

// GetInfoByButtonIds 根据按钮ID数组获取菜单按钮信息
func (s *sSysMenuButton) GetInfoByButtonIds(ctx context.Context, ids []int) (data []*entity.SysMenuButton, err error) {
	var tmpData *gvar.Var
	tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenuButton)

	var tmpSysMenuButton []*entity.SysMenuButton

	var menuButtonInfo []*entity.SysMenuButton
	//根据菜单ID数组获取菜单列表信息
	if tmpData.Val() != nil {
		if err = json.Unmarshal([]byte(tmpData.Val().(string)), &tmpSysMenuButton); err != nil {
			return
		}
		for _, id := range ids {
			for _, tmp := range tmpSysMenuButton {
				if id == int(tmp.Id) {
					menuButtonInfo = append(menuButtonInfo, tmp)
					continue
				}
			}
		}
	}
	if menuButtonInfo != nil && len(menuButtonInfo) > 0 {
		data = menuButtonInfo
		return
	}

	err = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().IsDeleted: 0,
		dao.SysMenuButton.Columns().Status:    1,
	}).WhereIn(dao.SysMenuButton.Columns().Id, ids).Scan(&data)
	return
}

// GetInfoByMenuIds 根据菜单ID数组获取菜单按钮信息
func (s *sSysMenuButton) GetInfoByMenuIds(ctx context.Context, menuIds []int) (data []*entity.SysMenuButton, err error) {
	//获取缓存菜单按钮信息
	for _, v := range menuIds {
		var tmpData *gvar.Var
		tmpData, err = cache.Instance().Get(ctx, consts.CacheSysMenuButton+"_"+gconv.String(v))
		if tmpData.Val() != nil {
			var sysMenuButton []*entity.SysMenuButton
			err = json.Unmarshal([]byte(tmpData.Val().(string)), &sysMenuButton)
			data = append(data, sysMenuButton...)
		}
	}
	if data == nil || len(data) == 0 {
		err = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
			dao.SysMenuButton.Columns().IsDeleted: 0,
			dao.SysMenuButton.Columns().Status:    1,
		}).WhereIn(dao.SysMenuButton.Columns().MenuId, menuIds).Scan(&data)
	}
	return
}

// GetInfoByMenuId 根据菜单ID数组获取菜单按钮信息
func (s *sSysMenuButton) GetInfoByMenuId(ctx context.Context, menuId int) (data []*entity.SysMenuButton, err error) {
	err = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().IsDeleted: 0,
		dao.SysMenuButton.Columns().Status:    1,
		dao.SysMenuButton.Columns().MenuId:    menuId,
	}).Scan(&data)
	if err != nil {
		return
	}
	if data != nil && len(data) > 0 {
		_ = cache.Instance().Set(ctx, consts.CacheSysMenuButton+"_"+gconv.String(menuId), data, 0)
	} else {
		_, err = cache.Instance().Remove(ctx, consts.CacheSysMenuButton+"_"+gconv.String(menuId))
	}

	return
}

// GetAll 获取所有的按钮信息
func (s *sSysMenuButton) GetAll(ctx context.Context) (data []*entity.SysMenuButton, err error) {
	err = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().IsDeleted: 0,
		dao.SysMenuButton.Columns().Status:    1,
	}).Scan(&data)

	if err != nil {
		return
	}
	if data != nil && len(data) > 0 {
		_ = cache.Instance().Set(ctx, consts.CacheSysMenuButton, data, 0)
	} else {
		_, err = cache.Instance().Remove(ctx, consts.CacheSysMenuButton)
	}

	return
}

// EditStatus 修改状态
func (s *sSysMenuButton) EditStatus(ctx context.Context, id int, menuId int, status int) (err error) {
	var menuButton *entity.SysMenuButton
	_ = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().Id: id,
	}).Scan(&menuButton)
	if menuButton == nil {
		return gerror.New("ID错误")
	}
	if menuButton.MenuId != menuId {
		return gerror.New("按钮不属于当前菜单,无法修改")
	}
	if menuButton != nil && menuButton.IsDeleted == 1 {
		return gerror.New("按钮已删除,无法修改")
	}
	if menuButton != nil && menuButton.Status == status {
		return gerror.New("按钮已禁用或启用,无须重复修改")
	}
	loginUserId := service.Context().GetUserId(ctx)
	menuButton.Status = status
	menuButton.UpdatedBy = loginUserId

	_, err = dao.SysMenuButton.Ctx(ctx).Data(menuButton).Where(g.Map{
		dao.SysMenuButton.Columns().Id: id,
	}).Update()

	//获取该菜单下所有的菜单按钮
	_, err = s.GetInfoByMenuId(ctx, menuButton.MenuId)
	//所有的菜单按钮
	_, err = s.GetAll(ctx)
	return
}

// 检查指定ID的数据是否存在
func checkMenuButtonId(ctx context.Context, Id int, menuButton *entity.SysMenuButton) *entity.SysMenuButton {
	_ = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
		dao.SysMenuButton.Columns().Id:        Id,
		dao.SysMenuButton.Columns().IsDeleted: 0,
	}).Scan(&menuButton)
	return menuButton
}

// 检查相同菜单按钮名称的数据是否存在
func checkMenuButtonName(ctx context.Context, menuId int, name string, menuButton *entity.SysMenuButton) *entity.SysMenuButton {
	m := dao.SysMenuButton.Ctx(ctx)
	_ = m.Where(g.Map{
		dao.SysMenuButton.Columns().Name:      name,
		dao.SysMenuButton.Columns().MenuId:    menuId,
		dao.SysMenuButton.Columns().IsDeleted: 0,
	}).Scan(&menuButton)
	return menuButton
}

// 检查相同菜单按钮类型的数据是否存在
func checkMenuButtonType(ctx context.Context, menuId int, types string, menuButton *entity.SysMenuButton) *entity.SysMenuButton {
	m := dao.SysMenuButton.Ctx(ctx)
	_ = m.Where(g.Map{
		dao.SysMenuButton.Columns().Types:     types,
		dao.SysMenuButton.Columns().MenuId:    menuId,
		dao.SysMenuButton.Columns().IsDeleted: 0,
	}).Scan(&menuButton)
	return menuButton
}
