package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type GetDictInput struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	DictType      string `p:"dictType" v:"required#字典类型不能为空"`
	DefaultValue  string `p:"defaultValue"`
}

type GetDictOut struct {
	Data   *DictTypeOut   `json:"info"`
	Values []*DictDataOut `json:"values"`
}

type GetDictRes struct {
	Data   *DictTypeRes   `json:"info"`
	Values []*DictDataRes `json:"values"`
}

type DictTypeOut struct {
	DictName string `json:"name"`
	Remark   string `json:"remark"`
}

type DictTypeRes struct {
	DictName string `json:"name"`
	Remark   string `json:"remark"`
}

type DictDataOut struct {
	DictValue string `json:"key"`
	DictLabel string `json:"value"`
	IsDefault int    `json:"isDefault"`
	Remark    string `json:"remark"`
}

// DictDataRes 字典数据
type DictDataRes struct {
	DictValue string `json:"key"`
	DictLabel string `json:"value"`
	IsDefault int    `json:"isDefault"`
	Remark    string `json:"remark"`
}

type SysDictSearchInput struct {
	DictType  string `p:"dictType"`  //字典类型
	DictLabel string `p:"dictLabel"` //字典标签
	Status    string `p:"status"`    //状态
	PaginationInput
}

type SysDictDataOut struct {
	DictCode  int64       `json:"dictCode"  description:"字典编码"`
	DictSort  int         `json:"dictSort"  description:"字典排序"`
	DictLabel string      `json:"dictLabel" description:"字典标签"`
	DictValue string      `json:"dictValue" description:"字典键值"`
	DictType  string      `json:"dictType"  description:"字典类型"`
	CssClass  string      `json:"cssClass"  description:"样式属性（其他样式扩展）"`
	ListClass string      `json:"listClass" description:"表格回显样式"`
	IsDefault int         `json:"isDefault" description:"是否默认（1是 0否）"`
	Status    int         `json:"status"    description:"状态（0正常 1停用）"`
	CreatedBy uint64      `json:"createdBy"  description:"创建者"`
	UpdatedBy uint64      `json:"updatedBy"  description:"更新者"`
	Remark    string      `json:"remark"    description:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}

type SysDictDataRes struct {
	DictCode  int64       `json:"dictCode"  description:"字典编码"`
	DictSort  int         `json:"dictSort"  description:"字典排序"`
	DictLabel string      `json:"dictLabel" description:"字典标签"`
	DictValue string      `json:"dictValue" description:"字典键值"`
	DictType  string      `json:"dictType"  description:"字典类型"`
	CssClass  string      `json:"cssClass"  description:"样式属性（其他样式扩展）"`
	ListClass string      `json:"listClass" description:"表格回显样式"`
	IsDefault int         `json:"isDefault" description:"是否默认（1是 0否）"`
	Status    int         `json:"status"    description:"状态（0正常 1停用）"`
	CreatedBy uint64      `json:"createdBy"  description:"创建者"`
	UpdatedBy uint64      `json:"updatedBy"  description:"更新者"`
	Remark    string      `json:"remark"    description:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
}

type AddDictDataInput struct {
	DictLabel string `p:"dictLabel"`
	DictValue string `p:"dictValue"`
	DictType  string `p:"dictType"`
	DictSort  int    `p:"dictSort"`
	CssClass  string `p:"cssClass"`
	ListClass string `p:"listClass"`
	IsDefault int    `p:"isDefault"`
	Status    int    `p:"status"`
	Remark    string `p:"remark"`
}

type EditDictDataInput struct {
	DictCode  int    `p:"dictCode"`
	DictLabel string `p:"dictLabel"`
	DictValue string `p:"dictValue"`
	DictType  string `p:"dictType"`
	DictSort  int    `p:"dictSort"`
	CssClass  string `p:"cssClass"`
	ListClass string `p:"listClass"`
	IsDefault int    `p:"isDefault"`
	Status    int    `p:"status"`
	Remark    string `p:"remark"`
}
