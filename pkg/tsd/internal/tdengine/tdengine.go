package tdengine

import (
	"context"
	"database/sql"
	"fmt"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/tsd/comm"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
)

type TdEngine struct {
	Option comm.Option
	db     *sql.DB
}

// Connect 连接数据库
func (m *TdEngine) connect() (db *sql.DB, err error) {
	tdDb, err := sql.Open(m.Option.DriverName, m.Option.Link+m.Option.Database)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return
	}
	tdDb.SetMaxIdleConns(m.Option.MaxIdleConns)
	tdDb.SetMaxOpenConns(m.Option.MaxOpenConns)
	m.db = tdDb
	//fmt.Println("Connected to TdEngine database.", m.Option)
	return m.db, nil
}
func (m *TdEngine) Close() {
	if m.db == nil {
		return
	}
	err := m.db.Close()
	if err != nil {
		fmt.Println("failed to close TDengine, err:", err)
		return
	}
	//fmt.Println("Closed TdEngine database connection.")
}

// Query 查询
func (m *TdEngine) Query(sql string) (rows *sql.Rows, err error) {
	if m.db == nil {
		_, err = m.connect()
		if err != nil {
			return
		}
	}

	rows, err = m.db.Query(sql)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}
	return
}

// Count 统计表中数据条数
func (m *TdEngine) Count(table string) (int, error) {
	if m.db == nil {
		_, err := m.connect()
		if err != nil {
			return 0, err
		}
	}
	sqlStr := "select count(*) as num from " + table
	rows, err := m.db.Query(sqlStr)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return 0, err
	}
	var num int
	for rows.Next() {
		err = rows.Scan(&num)
		if err != nil {
			fmt.Println("failed to query TDengine, err:", err)
			return 0, err
		}
	}
	return num, nil
}

// GetAllDatabaseName 获取所有数据库名称
func (m *TdEngine) GetAllDatabaseName(ctx context.Context) (names []string, err error) {
	if m.db == nil {
		if _, err = m.connect(); err != nil {
			return
		}
	}

	sqlStr := "show databases"
	rows, err := m.db.Query(sqlStr)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}

	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			fmt.Println("failed to query TDengine, err:", err)
			return
		}
		names = append(names, name)
	}
	return
}

// GetTableListByDatabase 获取指定的数据库下所有的表
func (m *TdEngine) GetTableListByDatabase(ctx context.Context, dbName string) (tableList []iotModel.TsdTables, err error) {
	if m.db == nil {
		if _, err = m.connect(); err != nil {
			return
		}
	}

	sqlStr := fmt.Sprintf("select table_name as tableName, db_name as dbName, create_time as createTime, stable_name as stableName from information_schema.ins_tables where db_name='%s'", dbName)
	rows, err := m.db.Query(sqlStr)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}

	for rows.Next() {
		var (
			dbName, stableName, tableName string
			createTime                    *gtime.Time
		)
		if err = rows.Scan(&tableName, &dbName, &createTime, &stableName); err != nil {
			fmt.Println("failed to query TDengine, err:", err)
			return
		}
		tableList = append(tableList, iotModel.TsdTables{
			DbName:     dbName,
			StableName: stableName,
			TableName:  tableName,
			CreateTime: createTime,
		})
	}
	return
}

// GetTableInfo 获取指定数据表结构信息
func (m *TdEngine) GetTableInfo(ctx context.Context, tableName string) (table []*iotModel.TsdTableInfo, err error) {
	if m.db == nil {
		if _, err = m.connect(); err != nil {
			return
		}
	}

	sqlStr := fmt.Sprintf("describe %s", tableName)
	rows, err := m.db.Query(sqlStr)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}

	for rows.Next() {
		tableInfo := new(iotModel.TsdTableInfo)
		if err = rows.Scan(&tableInfo.Field, &tableInfo.Type, &tableInfo.Length, &tableInfo.Note); err != nil {
			fmt.Println("failed to query TDengine, err:", err)
			return
		}
		table = append(table, tableInfo)
	}
	return
}

// GetTableData 获取指定数据表数据信息
func (m *TdEngine) GetTableData(ctx context.Context, tableName string) (table *iotModel.TsdTableDataInfo, err error) {
	if m.db == nil {
		if _, err = m.connect(); err != nil {
			return
		}
	}

	sqlStr := fmt.Sprintf("select * from %s", tableName)
	rows, err := m.db.Query(sqlStr)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}

	var (
		columns []string
		filed   []string
	)
	//获取查询结果字段
	if columns, err = rows.Columns(); err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}
	//封装scanArg
	scanArgs := make([]any, len(columns))
	for i := range columns {
		filed = append(filed, columns[i])
		scanArgs[i] = &columns[i]
	}
	table = new(iotModel.TsdTableDataInfo)
	table.Filed = append(table.Filed, filed...)
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			fmt.Println("failed to query TDengine, err:", err)
			return
		}
		//封装返回结果
		var resultMap = make(map[string]interface{})
		for i := range columns {
			resultMap[filed[i]] = columns[i]
		}
		table.Info = append(table.Info, resultMap)
	}
	return
}

// GetTableDataOne 获取超级表的单条数据
func (m *TdEngine) GetTableDataOne(ctx context.Context, sqlStr string, args ...any) (rs gdb.Record, err error) {
	if m.db == nil {
		if _, err = m.connect(); err != nil {
			return
		}
	}

	rows, err := m.db.Query(sqlStr, args...)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}

	list, err := g.DB().GetCore().RowsToResult(ctx, rows)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}
	if !list.IsEmpty() {
		rs = make(gdb.Record)
		for k, v := range list[0] {
			// 去除前缀
			prefixLen := len(comm.TdPropertyPrefix)
			if k[:prefixLen] == comm.TdPropertyPrefix {
				k = k[prefixLen:]
			}
			rs[k] = v
		}
	}
	return
}

// GetTableDataAll 获取超级表的多条数据
func (m *TdEngine) GetTableDataAll(ctx context.Context, sqlStr string, args ...any) (rs gdb.Result, err error) {
	if m.db == nil {
		if _, err = m.connect(); err != nil {
			return
		}
	}

	rows, err := m.db.Query(sqlStr, args...)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}

	list, err := g.DB().GetCore().RowsToResult(ctx, rows)
	if err != nil {
		fmt.Println("failed to query TDengine, err:", err)
		return
	}
	if !list.IsEmpty() {
		for _, rc := range list {
			newRc := make(gdb.Record)
			for k, v := range rc {
				// 去除前缀
				prefixLen := len(comm.TdPropertyPrefix)
				if k[:prefixLen] == comm.TdPropertyPrefix {
					k = k[prefixLen:]
				}
				newRc[k] = v
			}
			rs = append(rs, newRc)
		}
	}
	return
}
