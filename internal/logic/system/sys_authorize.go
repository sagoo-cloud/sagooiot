package system

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"strings"
)

type sSysAuthorize struct {
}

func init() {
	service.RegisterSysAuthorize(sysAuthorizeNew())
}

func sysAuthorizeNew() *sSysAuthorize {
	return &sSysAuthorize{}
}

func (s *sSysAuthorize) AuthorizeQuery(ctx context.Context, itemsType string, menuIds []int) (out []*model.AuthorizeQueryTreeOut, err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	if loginUserId == 0 {
		err = gerror.New("未登录或TOKEN失效,请重新登录")
		return
	}
	//查询用户角色信息
	userRoleInfo, err := service.SysUserRole().GetInfoByUserId(ctx, loginUserId)
	if err != nil {
		return
	}
	//判断是否为超级管理员
	if userRoleInfo == nil {
		err = gerror.New("无权限,请联系管理员")
		return
	}
	var isSuperAdmin = false
	var roleIds []int
	for _, userRole := range userRoleInfo {
		if userRole.RoleId == 1 {
			isSuperAdmin = true
		}
		roleIds = append(roleIds, userRole.RoleId)
	}
	if isSuperAdmin {
		var userMenuTreeOut []*model.UserMenuTreeOut
		if strings.EqualFold(itemsType, consts.Menu) {
			userMenuTreeOut, _ = GetMenuInfo(ctx, nil)
		} else {
			userMenuTreeOut, _ = GetMenuInfo(ctx, menuIds)
		}
		if userMenuTreeOut == nil {
			err = gerror.New("无可用菜单")
			return
		}

		var authorizeQueryTreeRes []*model.AuthorizeQueryTreeOut
		if err = gconv.Scan(userMenuTreeOut, &authorizeQueryTreeRes); err != nil {
			return
		}

		if !strings.EqualFold(itemsType, consts.Menu) {
			if len(menuIds) == 0 {
				err = gerror.New("请选择菜单")
				return
			}
			//根据项目类型 菜单ID封装菜单的按钮，列表字段,API接口
			authorizeItemsTypeTreeOut, userItemsTypeTreeErr := GetAuthorizeItemsTypeTreeOut(ctx, menuIds, itemsType, authorizeQueryTreeRes)
			if userItemsTypeTreeErr != nil {
				return nil, userItemsTypeTreeErr
			}

			//获取所有的子节点
			/*childrenMenuTreeOut := GetAllAuthorizeQueryChildrenTree(authorizeItemsTypeTreeOut)
			if childrenMenuTreeOut != nil {
				//查找所有的上级ID
				childrenMenuTreeOut = GetAllAuthorizeQueryParentTree(childrenMenuTreeOut, authorizeItemsTypeTreeOut)
			}*/

			//out = GetAuthorizeMenuTree(childrenMenuTreeOut)
			out = authorizeItemsTypeTreeOut
			return
		}
		out = GetAuthorizeMenuTree(authorizeQueryTreeRes)

		return
	} else {
		var userMenuTreeOut []*model.UserMenuTreeOut
		if strings.EqualFold(itemsType, consts.Menu) {
			//根据角色ID获取角色下配置的菜单信息
			authorizeItemsInfo, authorizeItemsErr := service.SysAuthorize().GetInfoByRoleIdsAndItemsType(ctx, roleIds, consts.Menu)
			if authorizeItemsErr != nil {
				return nil, authorizeItemsErr
			}
			if authorizeItemsInfo == nil {
				err = gerror.New("菜单未配置,请联系管理员")
				return
			}
			//获取菜单ID
			var authorizeItemsIds []int
			for _, authorizeItems := range authorizeItemsInfo {
				authorizeItemsIds = append(authorizeItemsIds, authorizeItems.ItemsId)
			}
			//根据菜单ID获取菜单信息
			userMenuTreeOut, _ = GetMenuInfo(ctx, authorizeItemsIds)
		} else {
			userMenuTreeOut, _ = GetMenuInfo(ctx, menuIds)
		}
		if userMenuTreeOut == nil {
			err = gerror.New("无可用菜单")
			return
		}
		var authorizeQueryTreeOut []*model.AuthorizeQueryTreeOut
		if err = gconv.Scan(userMenuTreeOut, &authorizeQueryTreeOut); err != nil {
			return
		}

		//根据项目类型 菜单ID封装菜单的按钮，列表字段,API接口
		if !strings.EqualFold(itemsType, consts.Menu) {
			if len(menuIds) == 0 {
				err = gerror.New("请选择菜单")
				return
			}
			//根据项目类型 菜单ID封装菜单的按钮，列表字段,API接口
			authorizeItemsTypeTreeOut, userItemsTypeTreeErr := GetAuthorizeItemsTypeTreeOut(ctx, menuIds, itemsType, authorizeQueryTreeOut)
			if userItemsTypeTreeErr != nil {
				return nil, userItemsTypeTreeErr
			}

			//获取所有的子节点
			/*childrenMenuTreeOut := GetAllAuthorizeQueryChildrenTree(authorizeItemsTypeTreeOut)
			if childrenMenuTreeOut != nil {
				//查找所有的上级ID
				childrenMenuTreeOut = GetAllAuthorizeQueryParentTree(childrenMenuTreeOut, authorizeItemsTypeTreeOut)
			}*/

			//out = GetAuthorizeMenuTree(childrenMenuTreeOut)
			out = authorizeItemsTypeTreeOut
			return
		}
		out = GetAuthorizeMenuTree(authorizeQueryTreeOut)
		return
	}
}

// GetInfoByRoleId 根据角色ID获取权限信息
func (s *sSysAuthorize) GetInfoByRoleId(ctx context.Context, roleId int) (data []*entity.SysAuthorize, err error) {
	err = dao.SysAuthorize.Ctx(ctx).Where(g.Map{
		dao.SysAuthorize.Columns().RoleId:    roleId,
		dao.SysAuthorize.Columns().IsDeleted: 0,
	}).Scan(&data)
	return
}

// GetInfoByRoleIds 根据角色ID数组获取权限信息
func (s *sSysAuthorize) GetInfoByRoleIds(ctx context.Context, roleIds []int) (data []*entity.SysAuthorize, err error) {
	//获取缓存菜单按钮信息
	for _, v := range roleIds {
		var tmpData *gvar.Var
		tmpData, err = cache.Instance().Get(ctx, consts.CacheSysAuthorize+"_"+gconv.String(v))
		if err != nil {
			return
		}
		if tmpData.Val() != nil {
			var sysAuthorizeInfo []*entity.SysAuthorize
			err = json.Unmarshal([]byte(tmpData.Val().(string)), &sysAuthorizeInfo)
			data = append(data, sysAuthorizeInfo...)
		}
	}
	if data == nil && len(data) == 0 {
		err = dao.SysAuthorize.Ctx(ctx).Where(g.Map{
			dao.SysAuthorize.Columns().IsDeleted: 0,
		}).WhereIn(dao.SysAuthorize.Columns().RoleId, roleIds).Scan(&data)
	}
	return
}

// GetInfoByRoleIdsAndItemsType 根据角色ID和项目类型获取权限信息
func (s *sSysAuthorize) GetInfoByRoleIdsAndItemsType(ctx context.Context, roleIds []int, itemsType string) (data []*entity.SysAuthorize, err error) {
	err = dao.SysAuthorize.Ctx(ctx).Where(g.Map{
		dao.SysAuthorize.Columns().IsDeleted: 0,
		dao.SysAuthorize.Columns().ItemsType: itemsType,
	}).WhereIn(dao.SysAuthorize.Columns().RoleId, roleIds).Scan(&data)
	return
}

func (s *sSysAuthorize) DelByRoleId(ctx context.Context, roleId int) (err error) {
	loginUserId := service.Context().GetUserId(ctx)
	_, err = dao.SysAuthorize.Ctx(ctx).Data(g.Map{
		dao.SysAuthorize.Columns().DeletedBy: uint(loginUserId),
		dao.SysAuthorize.Columns().IsDeleted: 1,
	}).Where(dao.SysAuthorize.Columns().RoleId, roleId).Update()
	_, err = dao.SysAuthorize.Ctx(ctx).Where(dao.SysAuthorize.Columns().RoleId, roleId).Delete()
	return
}

func (s *sSysAuthorize) Add(ctx context.Context, authorize []*entity.SysAuthorize) (err error) {
	var input []*model.SysAuthorizeInput
	if err = gconv.Scan(authorize, &input); err != nil {
		return
	}
	_, err = dao.SysAuthorize.Ctx(ctx).Data(input).Insert()
	return
}

func (s *sSysAuthorize) AddAuthorize(ctx context.Context, roleId int, menuIds []string, buttonIds []string, columnIds []string, apiIds []string) (err error) {
	//判断是否启用安全控制
	var configDataByIsSecurityControlEnabled *entity.SysConfig
	configDataByIsSecurityControlEnabled, err = service.ConfigData().GetConfigByKey(ctx, consts.SysIsSecurityControlEnabled)
	if err != nil {
		return
	}
	sysColumnSwitch := 0
	sysButtonSwitch := 0
	sysApiSwitch := 0

	if configDataByIsSecurityControlEnabled != nil && strings.EqualFold(configDataByIsSecurityControlEnabled.ConfigValue, "1") {
		//获取系统列表开关参数
		var sysColumnSwitchConfig *entity.SysConfig
		sysColumnSwitchConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.SysColumnSwitch)
		if sysColumnSwitchConfig != nil {
			sysColumnSwitch = gconv.Int(sysColumnSwitchConfig.ConfigValue)
		}
		//获取系统按钮开关参数
		var sysButtonSwitchConfig *entity.SysConfig
		sysButtonSwitchConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.SysButtonSwitch)
		if sysButtonSwitchConfig != nil {
			sysButtonSwitch = gconv.Int(sysButtonSwitchConfig.ConfigValue)
		}

		//获取系统API开关参数
		var sysApiSwitchConfig *entity.SysConfig
		sysApiSwitchConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.SysApiSwitch)
		if sysApiSwitchConfig != nil {
			sysApiSwitch = gconv.Int(sysApiSwitchConfig.ConfigValue)
		}
	}

	err = g.Try(ctx, func(ctx context.Context) {
		//删除原有权限
		err = service.SysAuthorize().DelByRoleId(ctx, roleId)
		var isTrue = true
		//封装菜单权限
		var authorizeInfo []*entity.SysAuthorize
		for _, id := range menuIds {
			var authorize = new(entity.SysAuthorize)
			split := strings.Split(id, "_")
			if len(split) < 2 {
				isTrue = false
				break
			}
			authorize.ItemsId = gconv.Int(split[0])
			authorize.ItemsType = consts.Menu
			authorize.RoleId = roleId
			authorize.IsCheckAll = gconv.Int(split[1])
			authorize.IsDeleted = 0
			authorizeInfo = append(authorizeInfo, authorize)
		}
		if !isTrue {
			err = gerror.New("菜单权限参数错误")
			return
		}
		//封装按钮权限
		if sysButtonSwitch == 1 {
			for _, id := range buttonIds {
				var authorize = new(entity.SysAuthorize)
				split := strings.Split(id, "_")
				if len(split) < 2 {
					isTrue = false
					break
				}
				authorize.ItemsId = gconv.Int(split[0])
				authorize.ItemsType = consts.Button
				authorize.RoleId = roleId
				authorize.IsCheckAll = gconv.Int(split[1])
				authorize.IsDeleted = 0
				authorizeInfo = append(authorizeInfo, authorize)
			}
			if !isTrue {
				err = gerror.New("按钮权限参数错误")
				return
			}
		}

		//封装列表权限
		if sysColumnSwitch == 1 {
			for _, id := range columnIds {
				var authorize = new(entity.SysAuthorize)
				split := strings.Split(id, "_")
				if len(split) < 2 {
					isTrue = false
					break
				}
				authorize.ItemsId = gconv.Int(split[0])
				authorize.ItemsType = consts.Column
				authorize.RoleId = roleId
				authorize.IsCheckAll = gconv.Int(split[1])
				authorize.IsDeleted = 0
				authorizeInfo = append(authorizeInfo, authorize)
			}
			if !isTrue {
				err = gerror.New("列表权限参数错误")
				return
			}
		}

		//封装接口权限
		if sysApiSwitch == 1 {
			for _, id := range apiIds {
				var authorize = new(entity.SysAuthorize)
				split := strings.Split(id, "_")
				if len(split) < 2 {
					isTrue = false
					break
				}
				authorize.ItemsId = gconv.Int(split[0])
				authorize.ItemsType = consts.Api
				authorize.RoleId = roleId
				authorize.IsCheckAll = gconv.Int(split[1])
				authorize.IsDeleted = 0
				authorizeInfo = append(authorizeInfo, authorize)
			}
			if !isTrue {
				err = gerror.New("接口权限参数错误")
				return
			}
		}

		err = s.Add(ctx, authorizeInfo)
		if err != nil {
			err = gerror.New("添加权限失败")
			return
		}
		//添加缓存信息
		err := cache.Instance().Set(ctx, consts.CacheSysAuthorize+"_"+gconv.String(roleId), authorizeInfo, 0)
		if err != nil {
			return
		}
	})
	return
}
func (s *sSysAuthorize) IsAllowAuthorize(ctx context.Context, roleId int) (isAllow bool, err error) {
	//判断角色ID是否为1
	if roleId == 1 {
		err = gerror.New("无法更改超级管理员权限")
		return
	}

	isAllow = false
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)
	//查看当前登录用户所有的角色
	//查询用户角色信息
	userRoleInfo, err := service.SysUserRole().GetInfoByUserId(ctx, loginUserId)
	if err != nil {
		return
	}
	//判断是否为超级管理员
	if userRoleInfo == nil {
		err = gerror.New("无权限,请联系管理员")
		return
	}
	var isSuperAdmin = false
	var roleIds []int
	for _, userRole := range userRoleInfo {
		if userRole.RoleId == 1 {
			isSuperAdmin = true
		}
		roleIds = append(roleIds, userRole.RoleId)
	}
	//判断当前用户是否为超级管理员
	if isSuperAdmin {
		//如果是超级管理员则可以对所有角色授权
		isAllow = true
	} else {
		//根据角色ID获取菜单配置
		authorizeInfo, authorizeErr := s.GetInfoByRoleId(ctx, roleId)
		if authorizeErr != nil {
			return
		}
		//如果需要授权的角色ID无任何权限，则当前用户可以对其授权
		if authorizeInfo == nil {
			isAllow = true
		} else {
			//菜单Ids
			var menuIds []int
			//按钮Ids
			var menuButtonIds []int
			//列表Ids
			var menuColumnIds []int
			//API Ids
			var menuApiIds []int
			for _, authorize := range authorizeInfo {
				if strings.EqualFold(authorize.ItemsType, consts.Menu) {
					menuIds = append(menuIds, authorize.ItemsId)
				} else if strings.EqualFold(authorize.ItemsType, consts.Button) {
					menuButtonIds = append(menuButtonIds, authorize.ItemsId)
				} else if strings.EqualFold(authorize.ItemsType, consts.Column) {
					menuColumnIds = append(menuColumnIds, authorize.ItemsId)
				} else if strings.EqualFold(authorize.ItemsType, consts.Api) {
					menuApiIds = append(menuApiIds, authorize.ItemsId)
				}
			}
			//获取当前用户所有的权限
			nowUserAuthorizeInfo, nowUserAuthorizeErr := s.GetInfoByRoleIds(ctx, roleIds)
			if nowUserAuthorizeErr != nil {
				return
			}

			//获取系统列表开关参数
			var sysColumnSwitchConfig *entity.SysConfig
			sysColumnSwitchConfig, err = service.ConfigData().GetConfigByKey(ctx, "sys.column.switch")
			sysColumnSwitch := 0
			if sysColumnSwitchConfig != nil {
				sysColumnSwitch = gconv.Int(sysColumnSwitchConfig.ConfigValue)
			}
			//获取系统按钮开关参数
			var sysButtonSwitchConfig *entity.SysConfig
			sysButtonSwitchConfig, err = service.ConfigData().GetConfigByKey(ctx, "sys.button.switch")
			sysButtonSwitch := 0
			if sysButtonSwitchConfig != nil {
				sysButtonSwitch = gconv.Int(sysButtonSwitchConfig.ConfigValue)
			}

			//获取系统API开关参数
			var sysApiSwitchConfig *entity.SysConfig
			sysApiSwitchConfig, err = service.ConfigData().GetConfigByKey(ctx, "sys.api.switch")
			sysApiSwitch := 0
			if sysApiSwitchConfig != nil {
				sysApiSwitch = gconv.Int(sysApiSwitchConfig.ConfigValue)
			}

			//菜单Ids
			var nowUserMenuIds []int
			//按钮Ids
			var nowUserMenuButtonIds []int
			//列表Ids
			var nowUserMenuColumnIds []int
			//API Ids
			var nowUserMenuApiIds []int
			for _, authorize := range nowUserAuthorizeInfo {
				if strings.EqualFold(authorize.ItemsType, consts.Menu) {
					nowUserMenuIds = append(nowUserMenuIds, authorize.ItemsId)
				} else if strings.EqualFold(authorize.ItemsType, consts.Button) && sysButtonSwitch == 1 {
					nowUserMenuButtonIds = append(nowUserMenuButtonIds, authorize.ItemsId)
				} else if strings.EqualFold(authorize.ItemsType, consts.Column) && sysColumnSwitch == 1 {
					nowUserMenuColumnIds = append(nowUserMenuColumnIds, authorize.ItemsId)
				} else if strings.EqualFold(authorize.ItemsType, consts.Api) && sysApiSwitch == 1 {
					nowUserMenuApiIds = append(nowUserMenuApiIds, authorize.ItemsId)
				}
			}

			//判断按钮、列表、API开关状态，如果关闭则获取菜单对应的所有信息
			//判断按钮开关
			if sysButtonSwitch == 0 {
				//获取所有按钮
				var menuButtons []*entity.SysMenuButton
				menuButtons, err = service.SysMenuButton().GetInfoByMenuIds(ctx, menuIds)
				if err != nil {
					return
				}
				if len(menuButtons) > 0 {
					for _, menuButton := range menuButtons {
						nowUserMenuButtonIds = append(nowUserMenuButtonIds, int(menuButton.Id))
					}
				}
			}

			//判断列表开关
			if sysColumnSwitch == 0 {
				//获取所有列表
				var menuColumns []*entity.SysMenuColumn
				menuColumns, err = service.SysMenuColumn().GetInfoByMenuIds(ctx, menuIds)
				if err != nil {
					return
				}
				if len(menuColumns) > 0 {
					for _, menuColumn := range menuColumns {
						nowUserMenuColumnIds = append(nowUserMenuColumnIds, int(menuColumn.Id))
					}
				}
			}

			//判断API开关
			if sysApiSwitch == 0 {
				//获取所有API
				var menuApis []*entity.SysMenuApi
				menuApis, err = service.SysMenuApi().GetInfoByMenuIds(ctx, menuIds)
				if err != nil {
					return
				}
				if len(menuApis) > 0 {
					for _, menuApi := range menuApis {
						nowUserMenuApiIds = append(nowUserMenuApiIds, int(menuApi.Id))
					}
				}
			}

			//判断当前登录用户是否大于授权角色的权限
			//获取当前登录用户的菜单信息
			nowUserMenuInfo, _ := service.SysMenu().GetInfoByMenuIds(ctx, nowUserMenuIds)
			//获取授权角色的菜单信息
			menuInfo, _ := service.SysMenu().GetInfoByMenuIds(ctx, menuIds)
			if len(menuInfo) == 0 {
				isAllow = true
			} else {
				var menuInfoIsAllow = true
				//判断当前登录用户所拥有的菜单是包含授权角色的菜单
				for _, menu := range menuInfo {
					var nowMenuIsAllow = false
					for _, nowUserMenu := range nowUserMenuInfo {
						if menu.Id == nowUserMenu.Id {
							nowMenuIsAllow = true
							break
						}
					}
					if !nowMenuIsAllow {
						menuInfoIsAllow = false
						break
					}
				}
				//判断是否都包含
				if menuInfoIsAllow {
					//判断按钮是否都包含
					//获取当前登录用户按钮单信息
					nowUserMenuButtonInfo, _ := service.SysMenuButton().GetInfoByMenuIds(ctx, nowUserMenuButtonIds)
					//获取授权角色的按钮信息
					menuButtonInfo, _ := service.SysMenuButton().GetInfoByMenuIds(ctx, menuButtonIds)
					var menuButtonInfoIsAllow = true
					if len(menuButtonInfo) != 0 {
						//判断当前登录用户所拥有的按钮是包含授权角色的按钮
						for _, menuButton := range menuButtonInfo {
							var nowMenuButtonIsAllow = false
							for _, nowUserMenuButton := range nowUserMenuButtonInfo {
								if menuButton.Id == nowUserMenuButton.Id {
									nowMenuButtonIsAllow = true
									break
								}
							}
							if !nowMenuButtonIsAllow {
								menuButtonInfoIsAllow = false
								break
							}
						}
						if menuButtonInfoIsAllow {
							//判断列表是否都包含
							//获取当前登录用户的列表信息
							nowUserMenuColumnInfo, _ := service.SysMenuColumn().GetInfoByMenuIds(ctx, nowUserMenuColumnIds)
							//获取授权角色的列表信息
							menuColumnInfo, _ := service.SysMenuColumn().GetInfoByMenuIds(ctx, menuColumnIds)
							var menuColumnInfoIsAllow = true
							if len(menuColumnInfo) == 0 {
								//判断当前登录用户所拥有的列表是包含授权角色的列表
								for _, menuColumn := range menuColumnInfo {
									var nowMenuColumnIsAllow = false
									for _, nowUserMenuColumn := range nowUserMenuColumnInfo {
										if menuColumn.Id == nowUserMenuColumn.Id {
											nowMenuColumnIsAllow = true
											break
										}
									}
									if !nowMenuColumnIsAllow {
										menuColumnInfoIsAllow = false
										break
									}
								}

								if menuColumnInfoIsAllow {
									isAllow = true
								}
							} else {
								isAllow = true
							}
						}

					} else {
						isAllow = true
					}

				}
			}
		}
	}
	return
}

// InitAuthorize 初始化系统权限
func (s *sSysAuthorize) InitAuthorize(ctx context.Context) (err error) {
	//获取所有菜单信息
	menuInfos, err := service.SysMenu().GetAll(ctx)
	if err != nil {
		return
	}
	if menuInfos != nil && len(menuInfos) > 0 {
		//根据菜单信息初始化对应按钮信息
		var menuIds []uint
		for _, menuInfo := range menuInfos {
			menuIds = append(menuIds, menuInfo.Id)
		}
		//获取所有的按钮列表
		var menuButtonInfos []*entity.SysMenuButton
		err = dao.SysMenuButton.Ctx(ctx).Where(g.Map{
			dao.SysMenuButton.Columns().IsDeleted: 0,
			dao.SysMenuButton.Columns().Status:    1,
		}).Scan(&menuButtonInfos)
		if err != nil {
			g.Log().Debug(ctx, "获取菜单按钮信息失败", err.Error())
			return
		}
		if menuButtonInfos != nil && len(menuButtonInfos) > 0 {
			err = cache.Instance().Set(ctx, consts.CacheSysMenuButton, menuButtonInfos, 0)
			if err != nil {
				return
			}
		}
		//获取所有的菜单列表
		var menuColumnInfos []*entity.SysMenuColumn
		err = dao.SysMenuColumn.Ctx(ctx).Where(g.Map{
			dao.SysMenuColumn.Columns().IsDeleted: 0,
			dao.SysMenuColumn.Columns().Status:    1,
		}).Scan(&menuColumnInfos)
		if err != nil {
			return
		}
		if menuColumnInfos != nil && len(menuColumnInfos) > 0 {
			err = cache.Instance().Set(ctx, consts.CacheSysMenuColumn, menuColumnInfos, 0)
			if err != nil {
				return
			}
		}
		//获取所有的菜单接口
		var menuApiInfos []*entity.SysMenuApi
		err = dao.SysMenuApi.Ctx(ctx).Where(g.Map{
			dao.SysMenuApi.Columns().IsDeleted: 0,
		}).Scan(&menuApiInfos)
		if err != nil {
			return
		}
		if menuApiInfos != nil && len(menuApiInfos) > 0 {
			err = cache.Instance().Set(ctx, consts.CacheSysMenuApi, menuApiInfos, 0)
			if err != nil {
				return
			}
		}
		//添加缓存信息
		for _, menuId := range menuIds {
			var tmpMenuButton []*entity.SysMenuButton
			for _, menuButtonInfo := range menuButtonInfos {
				if int(menuId) == menuButtonInfo.MenuId {
					tmpMenuButton = append(tmpMenuButton, menuButtonInfo)
				}
			}
			//添加按钮缓存
			if tmpMenuButton != nil && len(tmpMenuButton) > 0 {
				err = cache.Instance().Set(ctx, consts.CacheSysMenuButton+"_"+gconv.String(menuId), tmpMenuButton, 0)
				if err != nil {
					return
				}
			}

			var tmpMenuColumn []*entity.SysMenuColumn
			for _, menuColumnInfo := range menuColumnInfos {
				if int(menuId) == menuColumnInfo.MenuId {
					tmpMenuColumn = append(tmpMenuColumn, menuColumnInfo)
				}
			}
			//添加列表缓存
			if tmpMenuColumn != nil && len(tmpMenuColumn) > 0 {
				err = cache.Instance().Set(ctx, consts.CacheSysMenuColumn+"_"+gconv.String(menuId), tmpMenuColumn, 0)
				if err != nil {
					return
				}
			}

			var tmpMenuApi []*entity.SysMenuApi
			for _, menuApiInfo := range menuApiInfos {
				if int(menuId) == menuApiInfo.MenuId {
					tmpMenuApi = append(tmpMenuApi, menuApiInfo)
				}
			}
			//添加菜单与接口绑定关系缓存
			if tmpMenuApi != nil && len(tmpMenuApi) > 0 {
				err = cache.Instance().Set(ctx, consts.CacheSysMenuApi+"_"+gconv.String(menuId), tmpMenuApi, 0)
				if err != nil {
					return
				}
			}
		}
		//获取所有的接口信息
		var sysApiInfos []*entity.SysApi
		err = dao.SysApi.Ctx(ctx).Where(g.Map{
			dao.SysApi.Columns().IsDeleted: 0,
			dao.SysApi.Columns().Status:    1,
		}).Scan(&sysApiInfos)
		if err != nil {
			return
		}
		if sysApiInfos != nil && len(sysApiInfos) > 0 {
			err = cache.Instance().Set(ctx, consts.CacheSysApi, sysApiInfos, 0)
			if err != nil {
				return
			}
		}

		//获取所有的角色ID
		var roleInfos []*entity.SysRole
		err = dao.SysRole.Ctx(ctx).Where(g.Map{
			dao.SysRole.Columns().IsDeleted: 0,
			dao.SysRole.Columns().Status:    1,
		}).Scan(&roleInfos)
		if roleInfos != nil && len(roleInfos) > 0 {
			//获取所有的权限配置
			var authorizeInfos []*entity.SysAuthorize
			err = dao.SysAuthorize.Ctx(ctx).Where(g.Map{
				dao.SysAuthorize.Columns().IsDeleted: 0,
			}).Scan(&authorizeInfos)

			for _, roleInfo := range roleInfos {
				var tmpAuthorizeInfos []*entity.SysAuthorize
				for _, authorizeInfo := range authorizeInfos {
					if int(roleInfo.Id) == authorizeInfo.RoleId {
						tmpAuthorizeInfos = append(tmpAuthorizeInfos, authorizeInfo)
					}
				}
				if tmpAuthorizeInfos != nil && len(tmpAuthorizeInfos) > 0 {
					err = cache.Instance().Set(ctx, consts.CacheSysAuthorize+"_"+gconv.String(roleInfo.Id), tmpAuthorizeInfos, 0)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
