package iotModel

import "github.com/gogf/gf/v2/os/gtime"

type TsdTables struct {
	TableName  string      `json:"tableName"        description:"表名"`
	DbName     string      `json:"dbName"        description:"数据库名"`
	StableName string      `json:"stableName"        description:"超级表名"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
}

type TsdTableInfo struct {
	Field  string `json:"field"        description:"字段名"`
	Type   string `json:"type"        description:"类型"`
	Length int    `json:"length"        description:"长度"`
	Note   string `json:"note" description:"note"`
}

type TsdTableDataInfo struct {
	Filed []string                 `json:"filed"        description:"字段"`
	Info  []map[string]interface{} `json:"info"        description:"数据"`
}
