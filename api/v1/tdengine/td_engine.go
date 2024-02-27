package tdengine

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type GetTdEngineAllDbReq struct {
	g.Meta `path:"/getAllDb" method:"get" summary:"获取所有数据库" tags:"时序数据库管理"`
}

type GetTdEngineAllDbRes struct {
	Info []string
}

type GetTdEngineListTableByDatabasesReq struct {
	g.Meta `path:"/getListTablesByDatabases" method:"get" summary:"获取指定数据库下所有的表列表" tags:"时序数据库管理"`
	DbName string `json:"dbName"        description:"数据库名称" v:"required#数据库名称不能为空"`
}

type GetTdEngineListTableByDatabasesRes struct {
	Info []*model.TDEngineTablesList
}

type GetTdEngineTableInfoByTableReq struct {
	g.Meta    `path:"/getTdEngineTableInfoByTable" method:"get" summary:"获取指定数据表结构信息" tags:"时序数据库管理"`
	DbName    string `json:"dbName"        description:"数据库名称" v:"required#数据库名称不能为空"`
	TableName string `json:"tableName"        description:"数据库名称" v:"required#表名称不能为空"`
}

type GetTdEngineTableInfoByTableRes struct {
	Info []*model.TDEngineTableInfo
}

type GetTdEngineTableDataByTableReq struct {
	g.Meta    `path:"/getTdEngineTableDataByTable" method:"get" summary:"获取指定数据表数据信息" tags:"时序数据库管理"`
	DbName    string `json:"dbName"        description:"数据库名称" v:"required#数据库名称不能为空"`
	TableName string `json:"tableName"     description:"数据库名称" v:"required#表名称不能为空"`
}

type GetTdEngineTableDataByTableRes struct {
	*model.TableDataInfo
}
