package tdengine

import (
	"context"
	"database/sql"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	//_ "github.com/taosdata/driver-go/v3/taosSql" //原生驱动
)

type sTdEngine struct {
}

func tdEngineNew() *sTdEngine {
	return &sTdEngine{}
}

func init() {
	service.RegisterTdEngine(tdEngineNew())

	name, _ := g.Cfg().Get(context.TODO(), "tdengine.dbName")
	if name.String() != "" {
		dbName = name.String()
	}
}

// 数据库名
var dbName = "sagoo_iot"

// GetConn 获取链接
func (s *sTdEngine) GetConn(ctx context.Context, dbName string) (db *sql.DB, err error) {
	driver, err := g.Cfg().Get(ctx, "tdengine.type")
	if err != nil {
		err = gerror.New("请检查TDengine配置")
		return
	}
	dsn, err := g.Cfg().Get(ctx, "tdengine.dsn")
	if err != nil {
		err = gerror.New("请检查TDengine配置")
		return
	}

	link := dsn.String()
	if dbName != "" {
		link += dbName
	}

	taos, err := sql.Open(driver.String(), link)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("TDengine连接失败")
		return
	}
	return taos, nil
}

// GetTdEngineAllDb 获取所有数据库
func (s *sTdEngine) GetTdEngineAllDb(ctx context.Context) (data []string, err error) {
	taos, err := s.GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	defer taos.Close()

	rows, err := taos.Query("show databases;")
	if err != nil {
		err = gerror.New(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		data = append(data, name)
	}
	return
}

// GetListTableByDatabases 获取指定数据库下所有的表列表
func (s *sTdEngine) GetListTableByDatabases(ctx context.Context, dbName string) (data []*model.TDEngineTablesList, err error) {
	taos, err := s.GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	defer taos.Close()

	rows, err := taos.Query("SELECT table_name AS tableName, db_name AS dbName, create_time AS createTime, stable_name AS stableName FROM information_schema.ins_tables WHERE db_name = '" + dbName + "'")
	if err != nil {
		err = gerror.New(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tableName, db, stableName string
		var createTime *gtime.Time
		err = rows.Scan(&tableName, &db, &createTime, &stableName)
		if err != nil {
			err = gerror.New("获取失败")
			return
		}
		var tDEngineTablesList = new(model.TDEngineTablesList)
		tDEngineTablesList.TableName = tableName
		tDEngineTablesList.DbName = db
		tDEngineTablesList.StableName = stableName
		tDEngineTablesList.CreateTime = createTime
		data = append(data, tDEngineTablesList)
	}
	return
}

// GetTdEngineTableInfoByTable 获取指定数据表结构信息
func (s *sTdEngine) GetTdEngineTableInfoByTable(ctx context.Context, dbName string, tableName string) (data []*model.TDEngineTableInfo, err error) {
	taos, err := s.GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	defer taos.Close()

	rows, err := taos.Query("DESCRIBE " + dbName + "." + tableName + ";")
	if err != nil {
		err = gerror.New(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tDEngineTableInfo = new(model.TDEngineTableInfo)
		err = rows.Scan(&tDEngineTableInfo.Field, &tDEngineTableInfo.Type, &tDEngineTableInfo.Length, &tDEngineTableInfo.Note)
		if err != nil {
			err = gerror.New("获取失败")
			return
		}
		data = append(data, tDEngineTableInfo)
	}
	return
}

// GetTdEngineTableDataByTable 获取指定数据表数据信息
func (s *sTdEngine) GetTdEngineTableDataByTable(ctx context.Context, dbName string, tableName string) (data *model.TableDataInfo, err error) {
	data = new(model.TableDataInfo)
	taos, err := s.GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	defer taos.Close()

	rows, err := taos.Query("SELECT * FROM " + tableName)
	if err != nil {
		err = gerror.New(err.Error())
		return
	}
	defer rows.Close()

	//获取查询结果字段
	columns, _ := rows.Columns()
	//字段数组
	var filed []string
	//封装scanArg
	scanArgs := make([]any, len(columns))
	for i := range columns {
		filed = append(filed, columns[i])
		scanArgs[i] = &columns[i]
	}
	data.Filed = append(data.Filed, filed...)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			err = gerror.New("获取失败")
			return
		}
		//封装返回结果
		var resultMap = make(map[string]interface{})
		for i := range columns {
			resultMap[filed[i]] = columns[i]
		}
		data.Info = append(data.Info, resultMap)
	}

	return
}

// 超级表查询，单条数据
func (s *sTdEngine) GetOne(ctx context.Context, sql string, args ...any) (rs gdb.Record, err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	rows, err := taos.Query(sql, args...)
	if err != nil {
		g.Log().Error(ctx, err, sql, args)
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	values := make([]any, len(columns))
	rs = make(gdb.Record, len(columns))
	for i := range values {
		values[i] = new(any)
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		for i, c := range columns {
			rs[c] = s.Time(gvar.New(values[i]))
		}

		rows.Close()
	}

	return
}

// 超级表查询，多条数据
func (s *sTdEngine) GetAll(ctx context.Context, sql string, args ...any) (rs gdb.Result, err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	rows, err := taos.Query(sql, args...)
	if err != nil {
		g.Log().Error(ctx, err, sql)
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()

	for rows.Next() {
		values := make([]any, len(columns))
		for i := range values {
			values[i] = new(any)
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		m := make(gdb.Record, len(columns))
		for i, c := range columns {
			m[c] = s.Time(gvar.New(values[i]))
		}
		rs = append(rs, m)
	}

	return
}

// REST连接时区处理
func (s *sTdEngine) Time(v *g.Var) (rs *g.Var) {
	driver, _ := g.Cfg().Get(context.TODO(), "tdengine.type")

	if driver.String() == "taosRestful" {
		if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v.String()); err == nil {
			rs = gvar.New(t.Local().Format("2006-01-02 15:04:05"))
			return
		}
	}

	rs = v
	return
}
