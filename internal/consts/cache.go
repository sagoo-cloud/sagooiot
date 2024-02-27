/*
* @desc:缓存相关
 */

package consts

const (
	CacheModelMem   = "memory"
	CacheModelRedis = "redis"

	// CacheSysDict 字典缓存菜单KEY
	CacheSysDict = "SystemCache:sysDict"

	// CacheSysRole 角色缓存key
	CacheSysRole = "SystemCache:sysRole"

	// CacheSysDept 部门缓存key
	CacheSysDept = "SystemCache:sysDept"

	// CacheSysAuthTag 权限缓存TAG标签
	CacheSysAuthTag = "SystemCache:sysAuthTag"

	// CacheSysDictTag 字典缓存标签
	CacheSysDictTag = "SystemCache:sysDictTag"
	//CacheSysMenu 系统菜单
	CacheSysMenu = "SystemCache:sysMenu"
	//CacheSysMenuButton 系统菜单按钮
	CacheSysMenuButton = "SystemCache:sysMenuButton:"
	//CacheSysMenuColumn 系统菜单按钮
	CacheSysMenuColumn = "SystemCache:sysMenuColumn:"
	//CacheSysAuthorize 系统权限
	CacheSysAuthorize = "SystemCache:sysAuthorize:"
	//CacheSysMenuApi 系统API与菜单绑定关系表
	CacheSysMenuApi = "SystemCache:sysMenuApi:"
	//CacheSysApi 系统API
	CacheSysApi = "SystemCache:sysApi"
	//CacheUserAuthorize 用户权限
	CacheUserAuthorize = "SystemCache:userAuthorize"
	//CacheUserInfo 用户信息
	CacheUserInfo = "SystemCache:userInfo"

	//CacheIpBlackList IP访问黑名单
	CacheIpBlackList = "SystemCache:sysIpBlackList"

	//CacheDeviceOnline 下面的是网络部分用到的
	CacheDeviceOnline = "networkDeviceOnline"

	// 告警规则
	CacheAlarmRule = "AlarmRule:rule"
	// 服务器信息
	CacheServerInfo = "SystemCache:server_info"

	CacheSysErrorPrefix = "SysErrorPwdNum:"

	// 插件配置缓存
	PluginsTypeName = "plugins:%s:%s"

	// GoFrame ORM缓存前缀
	CacheGfOrmPrefix = "SelectCache:"
)
