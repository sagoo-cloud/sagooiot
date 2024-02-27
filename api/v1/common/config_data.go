package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type ConfigSearchReq struct {
	g.Meta         `path:"/config/list" tags:"系统参数管理" method:"get" summary:"系统参数列表"`
	ConfigName     string `p:"configName"`     //参数名称
	ConfigKey      string `p:"configKey"`      //参数键名
	ConfigType     string `p:"configType"`     //状态
	ModuleClassify string `p:"moduleClassify"` //字典分类编码
	*PaginationReq
}

type ConfigSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.SysConfigRes `json:"list"`
	PaginationRes
}

type ConfigAddReq struct {
	g.Meta         `path:"/config/add" tags:"系统参数管理" method:"post" summary:"添加系统参数"`
	ConfigName     string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey      string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue    string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType     int    `p:"configType"    v:"required|in:0,1#系统内置不能为空|系统内置类型只能为0或1"`
	ModuleClassify string `p:"moduleClassify" v:"required#字典分类不能为空"`
	Remark         string `p:"remark"`
}

type ConfigAddRes struct {
}

type ConfigGetReq struct {
	g.Meta `path:"/config/get" tags:"系统参数管理" method:"get" summary:"获取系统参数"`
	Id     int `p:"id" v:"required#ID不能为空"`
}

type ConfigGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *model.SysConfigRes `json:"data"`
}

type ConfigEditReq struct {
	g.Meta         `path:"/config/edit" tags:"系统参数管理" method:"put" summary:"修改系统参数"`
	ConfigId       int    `p:"configId" v:"required|min:1#主键ID不能为空|主键ID参数错误"`
	ConfigName     string `p:"configName"  v:"required#参数名称不能为空"`
	ConfigKey      string `p:"configKey"  v:"required#参数键名不能为空"`
	ConfigValue    string `p:"configValue"  v:"required#参数键值不能为空"`
	ConfigType     int    `p:"configType"    v:"required|in:0,1#系统内置不能为空|系统内置类型只能为0或1"`
	ModuleClassify string `p:"moduleClassify" v:"required#字典分类不能为空"`
	Remark         string `p:"remark"`
}

type ConfigEditRes struct {
}

type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete" tags:"系统参数管理" method:"delete" summary:"删除系统参数"`
	Ids    []int `p:"ids"`
}

type ConfigDeleteRes struct {
}

type ConfigGetByKeyReq struct {
	g.Meta    `path:"/config/getInfoByKey" tags:"系统参数管理" method:"get" summary:"根据KEY获取系统参数"`
	ConfigKey string `p:"configKey"   description:"参数键名" v:"required#参数键名不能为空"`
}

type ConfigGetByKeyRes struct {
	g.Meta `mime:"application/json"`
	Data   *model.SysConfigRes `json:"data"`
}

type ConfigGetByKeysReq struct {
	g.Meta    `path:"/config/getInfoByKeys" tags:"系统参数管理" method:"get" summary:"根据KEY数组获取系统参数"`
	ConfigKey []string `p:"configKey"   description:"参数键名" v:"required#参数键名不能为空"`
}

type ConfigGetByKeysRes struct {
	g.Meta `mime:"application/json"`
	Data   []*model.SysConfigRes `json:"data"`
}

type GetSysConfigSettingReq struct {
	g.Meta `path:"/getSysConfigSetting" method:"get" summary:"获取系统配置" tags:"系统参数管理"`
	Types  int `p:"types"   description:"类型  0 基础配置 1 安全配置"`
}
type GetSysConfigSettingRes struct {
	Info []*model.SysConfigRes `json:"data"`
}

type EditSysConfigSettingReq struct {
	g.Meta     `path:"/editSysConfigSetting" method:"put" summary:"修改系统配置" tags:"系统参数管理"`
	ConfigInfo []*model.EditConfigReq
}
type EditSysConfigSettingRes struct {
}
