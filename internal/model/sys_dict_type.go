package model

import "github.com/gogf/gf/v2/os/gtime"

type DictTypeDoInput struct {
	DictName       string `p:"dictName"`       //字典名称
	DictType       string `p:"dictType"`       //字典类型
	Status         string `p:"status"`         //字典状态
	ModuleClassify string `p:"moduleClassify"` //字典模块分类
	*PaginationInput
}
type SysDictTypeInfoOut struct {
	DictId         uint64      `orm:"dict_id,primary"  json:"dictId"`              // 字典主键
	DictName       string      `orm:"dict_name"        json:"dictName"`            // 字典名称
	DictType       string      `orm:"dict_type,unique" json:"dictType"`            // 字典类型
	Status         uint        `orm:"status"           json:"status"`              // 状态（0正常 1停用）
	Remark         string      `orm:"remark"           json:"remark"`              // 备注
	CreatedAt      *gtime.Time `orm:"created_at"       json:"createdAt"`           // 创建日期
	ModuleClassify string      `orm:"module_classify"       json:"moduleClassify"` // 字典模块分类
}

type AddDictTypeInput struct {
	DictName       string `p:"dictName"`
	DictType       string `p:"dictType"`
	Status         uint   `p:"status"`
	Remark         string `p:"remark"`
	ModuleClassify string `p:"moduleClassify"`
}

type EditDictTypeInput struct {
	DictId         int    `p:"dictId"`
	DictName       string `p:"dictName"`
	DictType       string `p:"dictType"`
	Status         uint   `p:"status"`
	Remark         string `p:"remark"`
	ModuleClassify string `p:"moduleClassify"`
}

type SysDictTypeInfoRes struct {
	DictId         uint64      `orm:"dict_id,primary"  json:"dictId"`              // 字典主键
	DictName       string      `orm:"dict_name"        json:"dictName"`            // 字典名称
	DictType       string      `orm:"dict_type,unique" json:"dictType"`            // 字典类型
	Status         uint        `orm:"status"           json:"status"`              // 状态（0正常 1停用）
	Remark         string      `orm:"remark"           json:"remark"`              // 备注
	CreatedAt      *gtime.Time `orm:"created_at"       json:"createdAt"`           // 创建日期
	ModuleClassify string      `orm:"module_classify"       json:"moduleClassify"` // 字典模块分类
}

type SysDictTypeOut struct {
	DictId         uint64      `json:"dictId"    description:"字典主键"`
	DictName       string      `json:"dictName"  description:"字典名称"`
	DictType       string      `json:"dictType"  description:"字典类型"`
	Status         uint        `json:"status"    description:"状态（0正常 1停用）"`
	CreateBy       uint        `json:"createBy"  description:"创建者"`
	UpdateBy       uint        `json:"updateBy"  description:"更新者"`
	Remark         string      `json:"remark"    description:"备注"`
	CreatedAt      *gtime.Time `json:"createdAt" description:"创建日期"`
	UpdatedAt      *gtime.Time `json:"updatedAt" description:"修改日期"`
	ModuleClassify string      `json:"moduleClassify" description:"字典模块分类"`
}

type SysDictTypeRes struct {
	DictId         uint64      `json:"dictId"    description:"字典主键"`
	DictName       string      `json:"dictName"  description:"字典名称"`
	DictType       string      `json:"dictType"  description:"字典类型"`
	Status         uint        `json:"status"    description:"状态（0正常 1停用）"`
	CreateBy       uint        `json:"createBy"  description:"创建者"`
	UpdateBy       uint        `json:"updateBy"  description:"更新者"`
	Remark         string      `json:"remark"    description:"备注"`
	CreatedAt      *gtime.Time `json:"createdAt" description:"创建日期"`
	UpdatedAt      *gtime.Time `json:"updatedAt" description:"修改日期"`
	ModuleClassify string      `json:"moduleClassify" description:"字典模块分类"`
}
