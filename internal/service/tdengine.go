// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"database/sql"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	ITdEngine interface {
		// GetConn 获取数据库连接
		GetConn(ctx context.Context, dbName string) (*sql.DB, error)
		// Time REST连接时区处理
		Time(v *g.Var) (rs *g.Var)
		// ClearLogByDays 删除指定天数的设备日志数据
		ClearLogByDays(ctx context.Context, days int) (err error)
	}
	ITdLogTable interface {
		// 添加超级表
		CreateStable(ctx context.Context) (err error)
		// 写入数据
		Insert(ctx context.Context, log *model.TdLogAddInput) (err error)
		// 清理过期数据
		Clear(ctx context.Context) (err error)
		// 超级表查询，多条数据
		GetAll(ctx context.Context, sql string, args ...any) (list []model.TdLog, err error)
	}
	ITSLTable interface {
		// Insert 数据入库
		Insert(ctx context.Context, deviceKey string, data model.ReportPropertyData, subKey ...string) (err error)
		// 添加超级表
		CreateStable(ctx context.Context, tsl *model.TSL) (err error)
		// 添加子表
		CreateTable(ctx context.Context, stable, table string) (err error)
		// 删除超级表
		DropStable(ctx context.Context, stable string) (err error)
		// 删除子表
		DropTable(ctx context.Context, table string) (err error)
		// 创建数据库
		CreateDatabase(ctx context.Context) (err error)
		// CheckStable 查询超级表是否存在, true=存在
		CheckStable(ctx context.Context, stable string) (b bool, err error)
		// CheckTable 查询子表是否存在, true=存在
		CheckTable(ctx context.Context, table string) (b bool, err error)
		// AddDatabaseField 添加数据库字段
		AddDatabaseField(ctx context.Context, tableName, fieldName string, dataType string, len int) (err error)
		// DelDatabaseField 删除数据库字段
		DelDatabaseField(ctx context.Context, tableName, fieldName string) (err error)
		// ModifyDatabaseField 修改数据库指定字段长度
		ModifyDatabaseField(ctx context.Context, tableName, fieldName string, dataType string, len int) (err error)
		// AddTag 添加标签
		AddTag(ctx context.Context, tableName, tagName string, dataType string, len int) (err error)
		// DelTag 删除标签
		DelTag(ctx context.Context, tableName, tagName string) (err error)
		// ModifyTag 修改标签
		ModifyTag(ctx context.Context, tableName, tagName string, dataType string, len int) (err error)
	}
)

var (
	localTdEngine   ITdEngine
	localTdLogTable ITdLogTable
	localTSLTable   ITSLTable
)

func TdEngine() ITdEngine {
	if localTdEngine == nil {
		panic("implement not found for interface ITdEngine, forgot register?")
	}
	return localTdEngine
}

func RegisterTdEngine(i ITdEngine) {
	localTdEngine = i
}

func TdLogTable() ITdLogTable {
	if localTdLogTable == nil {
		panic("implement not found for interface ITdLogTable, forgot register?")
	}
	return localTdLogTable
}

func RegisterTdLogTable(i ITdLogTable) {
	localTdLogTable = i
}

func TSLTable() ITSLTable {
	if localTSLTable == nil {
		panic("implement not found for interface ITSLTable, forgot register?")
	}
	return localTSLTable
}

func RegisterTSLTable(i ITSLTable) {
	localTSLTable = i
}
