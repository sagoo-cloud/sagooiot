package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"reflect"
	"strings"
)

// GetAllAuthorizeQueryChildrenTree 获取所有的子节点
func GetAllAuthorizeQueryChildrenTree(userMenuTreeRes []*model.AuthorizeQueryTreeOut) (childrenMenuTreeRes []*model.AuthorizeQueryTreeOut) {
	for k := range userMenuTreeRes {
		var isExistChildren = false
		for j := range userMenuTreeRes {
			if userMenuTreeRes[k].Id == uint(userMenuTreeRes[j].ParentId) {
				isExistChildren = true
			}
		}
		if !isExistChildren {
			if len(userMenuTreeRes[k].Children) > 0 {
				childrenMenuTreeRes = append(childrenMenuTreeRes, userMenuTreeRes[k])
			}
		}
	}
	return
}

// GetAllAuthorizeQueryParentTree 获取所有的父节点
func GetAllAuthorizeQueryParentTree(childrenMenuTreeRes []*model.AuthorizeQueryTreeOut, userMenuTreeRes []*model.AuthorizeQueryTreeOut) (data []*model.AuthorizeQueryTreeOut) {
	var parentMenuTreeRes []*model.AuthorizeQueryTreeOut
	for k := range childrenMenuTreeRes {
		for j := range userMenuTreeRes {
			if uint(childrenMenuTreeRes[k].ParentId) == userMenuTreeRes[j].Id {
				//判断父节点是否已存在
				var parentMenuIsExsit = false
				for _, p := range parentMenuTreeRes {
					if p.Id == userMenuTreeRes[j].Id {
						parentMenuIsExsit = true
					}
				}
				if !parentMenuIsExsit {
					parentMenuTreeRes = append(parentMenuTreeRes, userMenuTreeRes[j])
					childrenMenuTreeRes = append(childrenMenuTreeRes, userMenuTreeRes[j])
				}
			}
		}
	}
	if parentMenuTreeRes != nil {
		GetAllAuthorizeQueryParentTree(parentMenuTreeRes, userMenuTreeRes)
	}
	return childrenMenuTreeRes
}

// GetRoleTree 获取角色树
func GetRoleTree(roleInfo []*entity.SysRole) (dataTree []*model.RoleTreeOut, err error) {
	var parentNodeRes []*model.RoleTreeOut
	if roleInfo != nil {
		//获取所有的根节点
		for _, v := range roleInfo {
			var parentNode *model.RoleTreeOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeRes = append(parentNodeRes, parentNode)
			}
		}
	}
	treeData := RoleTree(parentNodeRes, roleInfo)
	return treeData, nil
}

// RoleTree 生成角色树结构
func RoleTree(parentNodeRes []*model.RoleTreeOut, data []*entity.SysRole) (dataTree []*model.RoleTreeOut) {
	//循环所有一级菜单
	for k, v := range parentNodeRes {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.RoleTreeOut
			if j.ParentId == int(v.Id) {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeRes[k].Children = append(parentNodeRes[k].Children, node)
			}
		}
		RoleTree(v.Children, data)
	}
	return parentNodeRes
}

func GetApiTree(apiInfo []*entity.SysApi) (dataTree []*model.SysApiTreeOut, err error) {
	var parentNodeRes []*model.SysApiTreeOut
	if apiInfo != nil {
		//获取所有的根节点
		for _, v := range apiInfo {
			var parentNode *model.SysApiTreeOut
			if v.ParentId == -1 && v.Types == 1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeRes = append(parentNodeRes, parentNode)
			}
		}
	}
	treeData := ApiTree(parentNodeRes, apiInfo)
	if len(parentNodeRes) == 0 {
		if err = gconv.Scan(apiInfo, &treeData); err != nil {
			return
		}
	}
	return treeData, nil
}

// ApiTree 生成接口树结构
func ApiTree(parentNodeRes []*model.SysApiTreeOut, data []*entity.SysApi) (dataTree []*model.SysApiTreeOut) {
	//循环所有一级菜单
	for k, v := range parentNodeRes {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			var node *model.SysApiTreeOut
			if j.ParentId == int(v.Id) {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeRes[k].Children = append(parentNodeRes[k].Children, node)
			}
		}
		ApiTree(v.Children, data)
	}
	return parentNodeRes
}

// GetMenuInfo 根据菜单ID获取指定菜单信息或者获取所有菜单信息
func GetMenuInfo(ctx context.Context, menuIds []int) (userMenuTreeOut []*model.UserMenuTreeOut, err error) {
	cache := common.Cache()
	//查看REDIS是否存在
	tmpData := cache.Get(ctx, consts.CacheSysMenu)
	//将缓存菜单转为struct
	var tmpMenuInfo []*entity.SysMenu
	json.Unmarshal([]byte(tmpData.Val().(string)), &tmpMenuInfo)

	var menuInfo []*entity.SysMenu
	if menuIds != nil {
		//根据菜单ID获取菜单信息
		if tmpData != nil {
			for _, menuId := range menuIds {
				for _, tmp := range tmpMenuInfo {
					if menuId == int(tmp.Id) {
						menuInfo = append(menuInfo, tmp)
						continue
					}
				}
			}
		} else {
			//获取所有的菜单
			menuInfo, err = service.SysMenu().GetInfoByMenuIds(ctx, menuIds)
			if err != nil {
				return
			}
		}

	} else {
		if tmpMenuInfo != nil {
			if err = gconv.Scan(tmpMenuInfo, &menuInfo); err != nil {
				return
			}
		} else {
			//获取所有的菜单
			menuInfo, err = service.SysMenu().GetAll(ctx)
			if err != nil {
				return
			}
		}
	}

	//获取所有的菜单信
	if menuInfo != nil {
		var userMenuTreeInfo []*model.UserMenuTreeOut
		if err = gconv.Scan(menuInfo, &userMenuTreeInfo); err != nil {
			return nil, err
		}
		//封装菜单数据
		//userMenuTreeInfoRes := GetUserMenuTree(userMenuTreeInfo)

		return userMenuTreeInfo, nil
	} else {
		err = gerror.New("无菜单,请先配置菜单")
		return
	}
}

// GetUserItemsTypeTreeOut 根据项目类型 菜单ID封装菜单的按钮，列表字段,API接口
func GetUserItemsTypeTreeOut(ctx context.Context, menuIds []int, itemsType string, userMenuTreeInfo []*model.UserMenuTreeOut) (userMenuTreeRes []*model.UserMenuTreeOut, err error) {
	if strings.EqualFold(itemsType, consts.Button) {
		var menuButtonInfo []*entity.SysMenuButton
		//根据菜单ID获取按钮列表
		menuButtonInfo, err = service.SysMenuButton().GetInfoByMenuIds(ctx, menuIds)
		if err != nil {
			return
		}
		if menuButtonInfo != nil {
			for _, menu := range userMenuTreeInfo {
				menuButtonTreeData, _ := GetUserMenuButton(int(menu.Id), menuButtonInfo)
				menu.Button = append(menu.Button, menuButtonTreeData...)
			}
		}
	} else if strings.EqualFold(itemsType, consts.Column) {

		var menuColumnInfo []*entity.SysMenuColumn
		//根据菜单ID获取列表字段
		menuColumnInfo, err = service.SysMenuColumn().GetInfoByMenuIds(ctx, menuIds)
		if err != nil {
			return
		}
		if menuColumnInfo != nil {
			for _, menu := range userMenuTreeInfo {
				menuColumnTreeData, _ := GetUserMenuColumn(int(menu.Id), menuColumnInfo)
				menu.Column = append(menu.Column, menuColumnTreeData...)
			}
		}

	} else if strings.EqualFold(itemsType, consts.Api) {
		//根据菜单ID获取列表字段
		menuApiInfo, er := service.SysMenuApi().GetInfoByMenuIds(ctx, menuIds)
		if er != nil {
			return
		}

		//获取相关接口ID
		var apiIds []int
		for _, menuApi := range menuApiInfo {
			apiIds = append(apiIds, menuApi.ApiId)
		}

		//获取相关接口信息
		apiInfo, _ := service.SysApi().GetInfoByIds(ctx, apiIds)

		if apiInfo != nil {
			for _, menu := range userMenuTreeInfo {
				var apiId = 0
				var menuApiId = 0
				var apiInfoOut []*model.UserApiOut
				if err = gconv.Scan(apiInfo, &apiInfoOut); err != nil {
					return
				}
				for _, menuApi := range menuApiInfo {
					if menuApi.MenuId == int(menu.Id) {
						apiId = menuApi.ApiId
						menuApiId = int(menuApi.Id)
						break
					}
				}
				if apiId != 0 {
					for _, api := range apiInfoOut {
						if apiId == api.Id {
							api.MenuApiId = menuApiId
							menu.Api = append(menu.Api, api)
							break
						}
					}
				}
			}
		}

	} else {
		err = gerror.New("itemsType参数错误")
		return
	}
	return userMenuTreeInfo, nil
}

// GetAuthorizeMenuTree 获取用户菜单树结构
func GetAuthorizeMenuTree(userMenuTreeOut []*model.AuthorizeQueryTreeOut) (dataTree []*model.AuthorizeQueryTreeOut) {
	var userMenuParentNodeTreeOut []*model.AuthorizeQueryTreeOut
	if userMenuTreeOut != nil {
		//获取所有的根节点
		for _, v := range userMenuTreeOut {
			if v.ParentId == -1 {
				userMenuParentNodeTreeOut = append(userMenuParentNodeTreeOut, v)
			}
		}
	}
	treeData := AuthorizeMenuTree(userMenuParentNodeTreeOut, userMenuTreeOut)
	return treeData
}

// AuthorizeMenuTree 重组菜单子节点
func AuthorizeMenuTree(userMenuParentNodeTreeOut []*model.AuthorizeQueryTreeOut, data []*model.AuthorizeQueryTreeOut) (dataTree []*model.AuthorizeQueryTreeOut) {
	//循环所有一级菜单
	for k, v := range userMenuParentNodeTreeOut {
		var childrenNodeTreeOut []*model.AuthorizeQueryTreeOut
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			if j.ParentId == int(v.Id) {
				//判断有无子节点
				childrenData := CheckChildrenAuthorizeMenuTree(j, data)
				if len(childrenData) > 0 {
					j.Children = append(j.Children, childrenData...)
				}
				var childrenMap g.Map
				if err := gconv.Scan(j, &childrenMap); err != nil {
					return
				}
				childrenNodeTreeOut = append(childrenNodeTreeOut, j)
				userMenuParentNodeTreeOut[k].Children = append(userMenuParentNodeTreeOut[k].Children, childrenMap)
			}
		}
	}
	return userMenuParentNodeTreeOut
}

func CheckChildrenAuthorizeMenuTree(children *model.AuthorizeQueryTreeOut, data []*model.AuthorizeQueryTreeOut) (childrenData []g.Map) {
	//查询所有该菜单下的所有子菜单
	for _, j := range data {
		if j.ParentId == int(children.Id) {
			j.Children = append(j.Children, CheckChildrenAuthorizeMenuTree(j, data)...)
			//判断有无子节点
			var childrenMap g.Map
			if err := gconv.Scan(j, &childrenMap); err != nil {
				return
			}
			childrenData = append(childrenData, childrenMap)

		}
	}
	return
}

// GetAuthorizeItemsTypeTreeOut 根据项目类型 菜单ID封装菜单的按钮，列表字段,API接口
func GetAuthorizeItemsTypeTreeOut(ctx context.Context, menuIds []int, itemsType string, authorizeMenuTreeInfo []*model.AuthorizeQueryTreeOut) (authorizeMenuTreeOut []*model.AuthorizeQueryTreeOut, err error) {
	if strings.EqualFold(itemsType, consts.Button) {
		//根据菜单ID获取按钮列表
		menuButtonInfo, er := service.SysMenuButton().GetInfoByMenuIds(ctx, menuIds)
		if er != nil {
			return
		}
		if menuButtonInfo != nil {
			for _, menu := range authorizeMenuTreeInfo {
				menuButtonTreeData, _ := GetUserMenuButton(int(menu.Id), menuButtonInfo)
				var childrenMaps []g.Map
				if err = gconv.Scan(menuButtonTreeData, &childrenMaps); err != nil {
					return
				}
				menu.Children = append(menu.Children, childrenMaps...)
			}
		}
	} else if strings.EqualFold(itemsType, consts.Column) {
		//根据菜单ID获取列表字段
		menuColumnInfo, er := service.SysMenuColumn().GetInfoByMenuIds(ctx, menuIds)
		if er != nil {
			return
		}
		if menuColumnInfo != nil {
			for _, menu := range authorizeMenuTreeInfo {
				menuColumnTreeData, _ := GetUserMenuColumn(int(menu.Id), menuColumnInfo)
				var childrenMaps []g.Map
				if err = gconv.Scan(menuColumnTreeData, &childrenMaps); err != nil {
					return
				}
				menu.Children = append(menu.Children, childrenMaps...)
			}
		}

	} else if strings.EqualFold(itemsType, consts.Api) {
		//根据菜单ID获取列表字段
		menuApiInfo, er := service.SysMenuApi().GetInfoByMenuIds(ctx, menuIds)
		if er != nil {
			return
		}

		//获取相关接口ID
		var apiIds []int
		for _, menuApi := range menuApiInfo {
			apiIds = append(apiIds, menuApi.ApiId)
		}

		//获取相关接口信息
		apiInfo, _ := service.SysApi().GetInfoByIds(ctx, apiIds)

		if apiInfo != nil {
			for _, menu := range authorizeMenuTreeInfo {
				var apiInfoOut []*model.AuthorizeQueryApiOut
				if err = gconv.Scan(apiInfo, &apiInfoOut); err != nil {
					return
				}
				for _, menuApi := range menuApiInfo {
					if menuApi.MenuId == int(menu.Id) {
						for _, api := range apiInfoOut {
							if menuApi.ApiId == api.Id {
								//菜单与接口绑定ID
								api.Id = int(menuApi.Id)
								//接口ID
								api.ApiId = api.Id
								api.Title = api.Name
								var childrenMap g.Map
								if err = gconv.Scan(api, &childrenMap); err != nil {
									return
								}
								menu.Children = append(menu.Children, childrenMap)
							}
						}
					}
				}
			}
		}

	} else {
		err = gerror.New("itemsType参数错误")
		return
	}
	return authorizeMenuTreeInfo, nil
}

// GetUserMenuTree 获取用户菜单树结构
func GetUserMenuTree(userMenuTreeRes []*model.UserMenuTreeOut) (dataTree []*model.UserMenuTreeOut) {
	var userMenuParentNodeTreeRes []*model.UserMenuTreeOut
	if userMenuTreeRes != nil {
		//获取所有的根节点
		for _, v := range userMenuTreeRes {
			if v.ParentId == -1 {
				userMenuParentNodeTreeRes = append(userMenuParentNodeTreeRes, v)
			}
		}
	}
	treeData := UserMenuTree(userMenuParentNodeTreeRes, userMenuTreeRes)
	return treeData
}

// UserMenuTree 重组菜单子节点
func UserMenuTree(userMenuParentNodeTreeRes []*model.UserMenuTreeOut, data []*model.UserMenuTreeOut) (dataTree []*model.UserMenuTreeOut) {
	//循环所有一级菜单
	for k, v := range userMenuParentNodeTreeRes {
		//查询所有该菜单下的所有子菜单
		for _, j := range data {
			if j.ParentId == int(v.Id) {
				userMenuParentNodeTreeRes[k].Children = append(userMenuParentNodeTreeRes[k].Children, j)
			}
		}
		UserMenuTree(v.Children, data)
	}
	return userMenuParentNodeTreeRes
}

// GetUserMenuButton 获取用户可操作按钮
func GetUserMenuButton(menuId int, menuButtonInfo []*entity.SysMenuButton) (dataTree []*model.UserMenuButtonOut, err error) {
	var menuButton []*entity.SysMenuButton
	for _, button := range menuButtonInfo {
		if menuId == button.MenuId {
			menuButton = append(menuButton, button)
		}
	}
	//获取所有按钮根节点
	var parentMenuButtonNodeRes []*model.UserMenuButtonOut
	if menuButton != nil {
		//获取所有的根节点
		for _, v := range menuButton {
			var parentMenuButtonNode *model.UserMenuButtonOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentMenuButtonNode); err != nil {
					return
				}
				parentMenuButtonNode.Title = v.Name
				parentMenuButtonNodeRes = append(parentMenuButtonNodeRes, parentMenuButtonNode)
			}
		}
	}
	//获取按钮树状结构
	menuButtonTreeData := UserMenuButtonTree(parentMenuButtonNodeRes, menuButton)
	return menuButtonTreeData, nil
}

// UserMenuButtonTree 生成树结构
func UserMenuButtonTree(parentMenuButtonNodeRes []*model.UserMenuButtonOut, data []*entity.SysMenuButton) (dataTree []*model.UserMenuButtonOut) {
	//循环所有一级菜单
	for k, v := range parentMenuButtonNodeRes {
		//查询所有该按钮下的所有子按钮
		for _, j := range data {
			var node *model.UserMenuButtonOut
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				node.Title = node.Name
				parentMenuButtonNodeRes[k].Children = append(parentMenuButtonNodeRes[k].Children, node)
			}
		}
		UserMenuButtonTree(v.Children, data)
	}
	return parentMenuButtonNodeRes
}

// GetUserMenuColumn 获取用户可查看列表
func GetUserMenuColumn(menuId int, menuColumnInfo []*entity.SysMenuColumn) (dataTree []*model.UserMenuColumnOut, err error) {
	var menuColumn []*entity.SysMenuColumn
	for _, column := range menuColumnInfo {
		if menuId == column.MenuId {
			menuColumn = append(menuColumn, column)
		}
	}
	//获取所有按钮根节点
	var parentMenuColumnNodeRes []*model.UserMenuColumnOut
	if menuColumn != nil {
		//获取所有的根节点
		for _, v := range menuColumn {
			var parentColumnButtonNode *model.UserMenuColumnOut
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentColumnButtonNode); err != nil {
					return
				}
				parentColumnButtonNode.Title = parentColumnButtonNode.Name
				parentMenuColumnNodeRes = append(parentMenuColumnNodeRes, parentColumnButtonNode)
			}
		}
	}
	//获取列表树状结构
	menuColumnTreeData := UserMenuColumnTree(parentMenuColumnNodeRes, menuColumn)
	return menuColumnTreeData, nil
}

// UserMenuColumnTree 生成树结构
func UserMenuColumnTree(parentMenuColumnNodeRes []*model.UserMenuColumnOut, data []*entity.SysMenuColumn) (dataTree []*model.UserMenuColumnOut) {
	//循环所有一级菜单
	for k, v := range parentMenuColumnNodeRes {
		//查询所有该按钮下的所有子按钮
		for _, j := range data {
			var node *model.UserMenuColumnOut
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				node.Title = node.Name
				parentMenuColumnNodeRes[k].Children = append(parentMenuColumnNodeRes[k].Children, node)
			}
		}
		UserMenuColumnTree(v.Children, data)
	}
	return parentMenuColumnNodeRes
}

// GetDataWhere 获取数据权限条件查询
func GetDataWhere(ctx context.Context, loginUserId int, entity interface{}) (where g.Map, err error) {
	//获取当前登录用户信息
	userInfo, err := service.SysUser().GetUserById(ctx, uint(loginUserId))
	if err != nil {
		err = gerror.New("登录用户信息错误")
		return
	}
	if userInfo == nil {
		err = gerror.New("登录用户不存在无法访问")
		return
	}
	if userInfo != nil && userInfo.Status == 0 {
		err = gerror.New("登录用户已禁用无法访问")
		return
	}
	if userInfo != nil && userInfo.Status == 2 {
		err = gerror.New("登录用户未验证无法访问")
		return
	}
	t := reflect.TypeOf(entity)
	for i := 0; i < t.Elem().NumField(); i++ {
		if t.Elem().Field(i).Name == "CreatedBy" {
			//若存在用户id的字段，则生成判断数据权限的条件
			//1、获取当前用户所属角色
			userRoleInfo, userRoleErr := service.SysUserRole().GetInfoByUserId(ctx, loginUserId)
			if userRoleErr != nil {
				err = gerror.New("获取用户角色失败")
				return
			}
			if userRoleInfo == nil {
				err = gerror.New("用户无权限访问")
				return
			}
			//判断用户是否为超级管理员
			var isSuperAdmin = false
			var roleIds []int
			for _, userRole := range userRoleInfo {
				if userRole.RoleId == 1 {
					isSuperAdmin = true
				}
				roleIds = append(roleIds, userRole.RoleId)
			}
			if isSuperAdmin {
				//超级管理员可以访问所有的数据
				return
			}
			//不是超级管理员则获取所有角色信息
			roleInfo, roleInfoErr := service.SysRole().GetInfoByIds(ctx, roleIds)
			if roleInfoErr != nil {
				err = gerror.New("获取用户角色失败")
				return
			}
			//2获取角色对应数据权限
			deptIdArr := gset.New()
			for _, role := range roleInfo {
				switch role.DataScope {
				case 1: //全部数据权限
					return
				case 2: //自定数据权限
					//获取角色所有的部门信息
					roleDeptInfo, _ := service.SysRoleDept().GetInfoByRoleId(ctx, int(role.Id))
					if roleDeptInfo == nil {
						err = gerror.New(role.Name + "自定义数据范围,请先配置部门!")
						return
					}
					var deptIds []int64
					for _, roleDept := range roleDeptInfo {
						deptIds = append(deptIds, roleDept.DeptId)
					}
					deptIdArr.Add(gconv.Interfaces(deptIds)...)
				case 3: //本部门数据权限
					deptIdArr.Add(gconv.Int64(userInfo.DeptId))
				case 4: //本部门及以下数据权限
					deptIdArr.Add(gconv.Int64(userInfo.DeptId))
					//获取所有的部门信息
					deptInfo, _ := service.SysDept().GetAll(ctx)
					if deptInfo != nil {
						//获取当前部门所有的下级部门信息
						childrenDeptInfo := GetNextDeptInfoByNowDeptId(int64(userInfo.DeptId), deptInfo)
						if childrenDeptInfo != nil {
							allChildrenDeptInfo := GetAllNextDeptInfoByChildrenDept(childrenDeptInfo, deptInfo, childrenDeptInfo)
							if allChildrenDeptInfo != nil {
								for _, allChildrenDept := range allChildrenDeptInfo {
									deptIdArr.Add(gconv.Int64(allChildrenDept.DeptId))
								}
							}
						}
					}
				case 5: //仅限于自己的数据
					where = g.Map{"created_by": userInfo.Id}
					return
				}
			}
			if deptIdArr.Size() > 0 {
				where = g.Map{"dept_id": deptIdArr.Slice()}
			}
		}
	}
	return
}

// GetNextDeptInfoByNowDeptId 获取当前部门ID下一层级的部门信息
func GetNextDeptInfoByNowDeptId(id int64, deptInfo []*entity.SysDept) (data []*entity.SysDept) {
	//循环所有的部门信息
	var childrenDept []*entity.SysDept
	for _, dept := range deptInfo {
		if dept.ParentId == id {
			//获取子部门信息
			childrenDept = append(childrenDept, dept)
		}
	}
	return childrenDept
}

// GetAllNextDeptInfoByChildrenDept 根据所有的子部门获取所有下级部门信息
func GetAllNextDeptInfoByChildrenDept(childrenDept []*entity.SysDept, deptInfo []*entity.SysDept, resultChildrenAll []*entity.SysDept) (data []*entity.SysDept) {
	var newChildrenDept []*entity.SysDept
	//循环所有的子部门信息
	for _, v := range childrenDept {
		//查询所有该按钮下的所有子按钮
		for _, j := range deptInfo {
			if j.ParentId == v.DeptId {
				newChildrenDept = append(newChildrenDept, j)
				resultChildrenAll = append(resultChildrenAll, j)
			}
		}
	}
	if newChildrenDept != nil {
		GetAllNextDeptInfoByChildrenDept(newChildrenDept, deptInfo, resultChildrenAll)
	}
	return resultChildrenAll
}
