package model

import "github.com/gogf/gf/v2/os/gtime"

type GenTableColumnRes struct {
	Id               int64       `json:"id"               description:"编号"`
	TableId          int64       `json:"tableId"          description:"归属表编号"`
	ColumnName       string      `json:"columnName"       description:"列名称"`
	ColumnComment    string      `json:"columnComment"    description:"列描述"`
	ColumnType       string      `json:"columnType"       description:"列类型"`
	GoType           string      `json:"goType"           description:"Go类型"`
	GoField          string      `json:"goField"          description:"Go字段名"`
	HtmlField        string      `json:"htmlField"        description:"html字段名"`
	IsPk             string      `json:"isPk"             description:"是否主键（1是）"`
	IsIncrement      string      `json:"isIncrement"      description:"是否自增（1是）"`
	IsRequired       string      `json:"isRequired"       description:"是否必填（1是）"`
	IsInsert         string      `json:"isInsert"         description:"是否为插入字段（1是）"`
	IsEdit           string      `json:"isEdit"           description:"是否编辑字段（1是）"`
	IsList           string      `json:"isList"           description:"是否列表字段（1是）"`
	IsQuery          string      `json:"isQuery"          description:"是否查询字段（1是）"`
	QueryType        string      `json:"queryType"        description:"查询方式（等于、不等于、大于、小于、范围）"`
	HtmlType         string      `json:"htmlType"         description:"显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）"`
	DictType         string      `json:"dictType"         description:"字典类型"`
	Sort             int         `json:"sort"             description:"排序"`
	LinkTableName    string      `json:"linkTableName"    description:"关联表名"`
	LinkTableClass   string      `json:"linkTableClass"   description:"关联表类名"`
	LinkTablePackage string      `json:"linkTablePackage" description:"关联表包名"`
	LinkLabelId      string      `json:"linkLabelId"      description:"关联表键名"`
	LinkLabelName    string      `json:"linkLabelName"    description:"关联表字段值"`
	CreateBy         int         `json:"createBy"         description:"创建者"`
	UpdateBy         int         `json:"updateBy"         description:"更新者"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
}
type GenTableColumnAddInput struct {
	TableId          int64       `json:"tableId"          description:"归属表编号"`
	ColumnName       string      `json:"columnName"       description:"列名称"`
	ColumnComment    string      `json:"columnComment"    description:"列描述"`
	ColumnType       string      `json:"columnType"       description:"列类型"`
	GoType           string      `json:"goType"           description:"Go类型"`
	GoField          string      `json:"goField"          description:"Go字段名"`
	HtmlField        string      `json:"htmlField"        description:"html字段名"`
	IsPk             string      `json:"isPk"             description:"是否主键（1是）"`
	IsIncrement      string      `json:"isIncrement"      description:"是否自增（1是）"`
	IsRequired       string      `json:"isRequired"       description:"是否必填（1是）"`
	IsInsert         string      `json:"isInsert"         description:"是否为插入字段（1是）"`
	IsEdit           string      `json:"isEdit"           description:"是否编辑字段（1是）"`
	IsList           string      `json:"isList"           description:"是否列表字段（1是）"`
	IsQuery          string      `json:"isQuery"          description:"是否查询字段（1是）"`
	QueryType        string      `json:"queryType"        description:"查询方式（等于、不等于、大于、小于、范围）"`
	HtmlType         string      `json:"htmlType"         description:"显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）"`
	DictType         string      `json:"dictType"         description:"字典类型"`
	Sort             int         `json:"sort"             description:"排序"`
	LinkTableName    string      `json:"linkTableName"    description:"关联表名"`
	LinkTableClass   string      `json:"linkTableClass"   description:"关联表类名"`
	LinkTablePackage string      `json:"linkTablePackage" description:"关联表包名"`
	LinkLabelId      string      `json:"linkLabelId"      description:"关联表键名"`
	LinkLabelName    string      `json:"linkLabelName"    description:"关联表字段值"`
	CreateBy         int         `json:"createBy"         description:"创建者"`
	UpdateBy         int         `json:"updateBy"         description:"更新者"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
}
type GenTableColumnEditInput struct {
	Id int `json:"id"          description:"ID"`
	GenTableColumnAddInput
}

// GenTableAndColumnsRes 表与字段组合数据
type GenTableAndColumnsRes struct {
	TableInfo *GenTableRes
	Columns   []*GenTableColumnRes `json:"columns"`
}
