package model

import "github.com/gogf/gf/v2/os/gtime"

type ConfigDoInput struct {
	ConfigName    string `p:"configName"`    //参数名称
	ConfigKey     string `p:"configKey"`     //参数键名
	ConfigType    string `p:"configType"`    //状态
	DictClassCode string `p:"dictClassCode"` //字典类型编码
	*PaginationInput
}

type SysConfigRes struct {
	ConfigId      uint        `json:"configId"    description:"参数主键"`
	ConfigName    string      `json:"configName"  description:"参数名称"`
	ConfigKey     string      `json:"configKey"   description:"参数键名"`
	ConfigValue   string      `json:"configValue" description:"参数键值"`
	ConfigType    int         `json:"configType"  description:"系统内置（Y是 N否）"`
	DictClassCode string      `json:"dictClassCode" description:"字典类型编码"`
	CreateBy      uint        `json:"createBy"    description:"创建者"`
	UpdateBy      uint        `json:"updateBy"    description:"更新者"`
	Remark        string      `json:"remark"      description:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"   description:"修改时间"`
}

type SysConfigOut struct {
	ConfigId      uint        `json:"configId"    description:"参数主键"`
	ConfigName    string      `json:"configName"  description:"参数名称"`
	ConfigKey     string      `json:"configKey"   description:"参数键名"`
	ConfigValue   string      `json:"configValue" description:"参数键值"`
	ConfigType    int         `json:"configType"  description:"系统内置（Y是 N否）"`
	DictClassCode string      `json:"dictClassCode" description:"字典类型编码"`
	CreateBy      uint        `json:"createBy"    description:"创建者"`
	UpdateBy      uint        `json:"updateBy"    description:"更新者"`
	Remark        string      `json:"remark"      description:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"   description:"修改时间"`
}

type AddConfigInput struct {
	ConfigName    string `p:"configName"`
	ConfigKey     string `p:"configKey"`
	ConfigValue   string `p:"configValue"`
	ConfigType    int    `p:"configType"`
	Remark        string `p:"remark"`
	DictClassCode string `p:"dictClassCode"`
}

type EditConfigInput struct {
	ConfigId      int    `p:"configId"`
	ConfigName    string `p:"configName"`
	ConfigKey     string `p:"configKey"`
	ConfigValue   string `p:"configValue"`
	ConfigType    int    `p:"configType"`
	Remark        string `p:"remark"`
	DictClassCode string `p:"dictClassCode"`
}
