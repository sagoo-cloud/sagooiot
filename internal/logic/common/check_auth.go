package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
	"strings"
)

type sCheckAuth struct {
}

func CheckAuth() *sCheckAuth {
	return &sCheckAuth{}
}

func init() {
	service.RegisterCheckAuth(CheckAuth())
}

// IsToken 验证TOKEN是否正确
func (s *sCheckAuth) IsToken(ctx context.Context) (isToken bool, expiresAt int64, isAuth string, err error) {
	authorization := ghttp.RequestFromCtx(ctx).Header.Get("Authorization")
	if authorization == "" {
		err = gerror.New("请先登录!")
		return
	}
	isToken = false
	expiresAt = 0
	//验证TOKEN是否正确
	data, _ := service.SysToken().ParseToken(ghttp.RequestFromCtx(ctx))
	if data != nil {
		isToken = true
		expiresAt = gconv.Int64(data.ExpiresAt)
	}
	//获取当前登录用户账户
	userName := service.Context().GetUserName(ctx)

	dict, err := service.DictData().GetDictDataByType(ctx, "rule_engine_user_blacklist")
	if err != nil {
		return
	}
	isAuth = "save"
	if dict != nil && dict.Values != nil && len(dict.Values) > 0 {
		for _, v := range dict.Values {
			if strings.EqualFold(v.DictValue, userName) {
				isAuth = "read"
				break
			}
		}
	}
	return
}

// CheckAccessAuth 验证访问权限
func (s *sCheckAuth) CheckAccessAuth(ctx context.Context, address string) (isAllow bool, err error) {
	isAllow = false
	//查询API开关是否打开
	sysApiSwitchConfig, _ := service.ConfigData().GetConfigByKey(ctx, "sys.api.switch")
	sysApiSwitch := 0
	if sysApiSwitchConfig != nil {
		sysApiSwitch = gconv.Int(sysApiSwitchConfig.ConfigValue)
	}
	if sysApiSwitch == 0 {
		isAllow = true
		return
	}

	//获取用户角色信息
	userRoleInfo, err := service.SysUserRole().GetInfoByUserId(ctx, service.Context().GetUserId(ctx))
	if err != nil {
		err = gerror.New("获取用户角色失败")
		return
	}
	if userRoleInfo == nil {
		err = gerror.New("用户未配置角色信息,请联系管理员")
		return
	}

	var roleIds []int
	//判断是否为超级管理员
	var isSuperAdmin = false
	for _, userRole := range userRoleInfo {
		//获取角色ID
		if userRole.RoleId == 1 {
			isSuperAdmin = true
		}
		roleIds = append(roleIds, userRole.RoleId)
	}

	//超级管理员拥有所有访问权限
	if isSuperAdmin {
		isAllow = true
		return
	}

	//获取角色ID下所有的请求API
	authorizeInfo, authorizeErr := service.SysAuthorize().GetInfoByRoleIdsAndItemsType(ctx, roleIds, consts.Api)
	if authorizeErr != nil {
		err = gerror.New("获取用户权限失败")
		return
	}

	if authorizeInfo == nil || len(authorizeInfo) == 0 {
		err = gerror.New("未授权接口,无访问权限!")
		return
	}

	//判断是否与当前访问接口一致
	var menuApiIds []int
	for _, authorize := range authorizeInfo {
		menuApiIds = append(menuApiIds, authorize.ItemsId)
	}
	//获取所有的接口API
	menuApiInfo, menuApiErr := service.SysMenuApi().GetInfoByIds(ctx, menuApiIds)
	if menuApiErr != nil {
		err = gerror.New("相关接口未配置")
		return
	}
	if menuApiInfo == nil || len(menuApiInfo) == 0 {
		err = gerror.New("接口未绑定菜单,请联系管理员!")
		return
	}
	var apiIds []int
	for _, menuApi := range menuApiInfo {
		apiIds = append(apiIds, menuApi.ApiId)
	}
	//获取所有的接口
	apiInfo, apiErr := service.SysApi().GetInfoByIds(ctx, apiIds)
	if apiErr != nil {
		err = gerror.New("获取接口失败")
		return
	}
	if apiInfo == nil || len(apiInfo) == 0 {
		err = gerror.New("相关接口未配置")
		return
	}

	var isExist = false
	//获取请求路径
	for _, api := range apiInfo {
		if strings.EqualFold(address, api.Address) {
			isExist = true
			break
		}
	}
	if isExist {
		isAllow = true
		return
	}
	return
}
