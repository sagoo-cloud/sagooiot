package tdengine

import (
	"context"
	"fmt"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/tsd/comm"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 物模型 TDengine 表结构维护
type sTSLTable struct{}

func init() {
	service.RegisterTSLTable(tslTableNew())
}

func tslTableNew() *sTSLTable {
	return &sTSLTable{}
}

// Insert 数据入库
func (s *sTSLTable) Insert(ctx context.Context, deviceKey string, data model.ReportPropertyData, subKey ...string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	if len(data) == 0 {
		err = gerror.New("上报数据为空")
		return
	}

	ts := gtime.Now().Format("Y-m-d H:i:s")

	var (
		field = []string{"ts"}
		value = []string{"'" + ts + "'"}
	)
	for k, v := range data {
		k = comm.TsdColumnName(k)

		field = append(field, k)
		value = append(value, "'"+gvar.New(v.Value).String()+"'")
		// 属性上报时间
		field = append(field, k+"_time")
		value = append(value, "'"+gtime.New(v.CreateTime).Format("Y-m-d H:i:s")+"'")
	}

	if len(subKey) == 0 {
		deviceKey = comm.DeviceTableName(deviceKey)
		sql := "INSERT INTO ? (?) VALUES (?)"
		_, err = taos.Exec(sql, deviceKey, strings.Join(field, ","), strings.Join(value, ","))
	} else {
		// 子设备
		skey := comm.DeviceTableName(subKey[0])
		sql := "INSERT INTO ? (?) VALUES (?)"
		_, err = taos.Exec(sql, skey, strings.Join(field, ","), strings.Join(value, ","))
	}
	return
}

// 添加超级表
func (s *sTSLTable) CreateStable(ctx context.Context, tsl *model.TSL) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	// 属性字段
	columns := []string{"ts TIMESTAMP"}
	for _, v := range tsl.Properties {
		maxLength := 0
		if v.ValueType.TSLParamBase.MaxLength != nil {
			maxLength = *v.ValueType.TSLParamBase.MaxLength
		}
		columns = append(columns, s.column(v.ValueType.Type, v.Key, v.Name, maxLength))
		// 属性上报时间
		columns = append(columns, s.column("date", v.Key+"_time", "", 0))
	}

	// 标签字段
	tags := make([]string, len(tsl.Tags)+1)
	tags[0] = "device VARCHAR(255)"
	for i, v := range tsl.Tags {
		maxLength := 0
		if v.ValueType.TSLParamBase.MaxLength != nil {
			maxLength = *v.ValueType.TSLParamBase.MaxLength
		}
		tags[i+1] = s.column(v.ValueType.Type, v.Key, v.Name, maxLength, 1)
	}

	tConent := ""
	if len(tags) > 0 {
		tConent = fmt.Sprintf("TAGS (%s)", strings.Join(tags, ","))
	}
	table := comm.ProductTableName(tsl.Key)
	sql := fmt.Sprintf("CREATE STABLE IF NOT EXISTS %s.%s (%s) %s", dbName, table, strings.Join(columns, ","), tConent)

	_, err = taos.Exec(sql)

	return
}

// 添加子表
func (s *sTSLTable) CreateTable(ctx context.Context, stable, table string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tag := table
	stable = comm.ProductTableName(stable)
	table = comm.DeviceTableName(table)

	sql := fmt.Sprintf("CREATE TABLE  IF NOT EXISTS %s USING %s (device) TAGS ('%s')", table, stable, tag)

	_, err = taos.Exec(sql)

	return
}

func (s *sTSLTable) column(dataType, key, name string, maxLength int, isTag ...int) string {
	column := ""
	tdType := ""
	switch dataType {
	case "int":
		tdType = "INT"
	case "long":
		tdType = "BIGINT"
	case "float":
		tdType = "FLOAT"
	case "double":
		tdType = "DOUBLE"
	case "string":
		if maxLength == 0 {
			maxLength = 150
		}
		tdType = "NCHAR(" + strconv.Itoa(maxLength) + ")"
	case "boolean":
		tdType = "BOOL"
	case "date":
		tdType = "TIMESTAMP"
	default:
		if maxLength == 0 {
			maxLength = 150
		}
		tdType = "NCHAR(" + strconv.Itoa(maxLength) + ")"
	}

	// 属性、tag加前缀
	if len(isTag) > 0 && isTag[0] == 1 {
		key = comm.TsdTagName(key)
	} else {
		key = comm.TsdColumnName(key)
	}

	column = fmt.Sprintf("%s %s", key, tdType)
	return column
}

// 删除超级表
func (s *sTSLTable) DropStable(ctx context.Context, stable string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	stable = comm.ProductTableName(stable)

	sql := fmt.Sprintf("DROP STABLE IF EXISTS %s.%s", dbName, stable)
	_, err = taos.Exec(sql)

	return
}

// 删除子表
func (s *sTSLTable) DropTable(ctx context.Context, table string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	table = comm.DeviceTableName(table)

	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", dbName, table)
	_, err = taos.Exec(sql)

	return
}

// 创建数据库
func (s *sTSLTable) CreateDatabase(ctx context.Context) (err error) {
	// 资源锁
	lockKey := "tdLock:initDb"
	lockVal, err := g.Redis().Do(ctx, "SET", lockKey, gtime.Now().Unix(), "NX", "EX", "3600")
	if err != nil {
		return
	}
	if lockVal.String() != "OK" {
		return
	}
	defer func() {
		_, err = g.Redis().Do(ctx, "DEL", lockKey)
	}()

	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	/*var name string
	if err = taos.QueryRow("SELECT name FROM information_schema.ins_databases WHERE name = '?' LIMIT 1", dbName).Scan(&name); err != nil {
		return
	}

	if name != "" {
		return
	}*/

	_, err = taos.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)

	return
}

// CheckStable 查询超级表是否存在, true=存在
func (s *sTSLTable) CheckStable(ctx context.Context, stable string) (b bool, err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	var name string
	if err = taos.QueryRow("SELECT stable_name FROM information_schema.ins_stables WHERE db_name = '?' AND stable_name = '?' LIMIT 1", dbName, stable).Scan(&name); err != nil {
		return
	}
	if name == "" {
		return
	}

	b = true
	return
}

// CheckTable 查询子表是否存在, true=存在
func (s *sTSLTable) CheckTable(ctx context.Context, table string) (b bool, err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	var name string
	if err = taos.QueryRowContext(ctx, "SELECT table_name FROM information_schema.ins_tables WHERE db_name = '?' AND table_name = '?' LIMIT 1", dbName, table).Scan(&name); err != nil {
		return
	}
	if name == "" {
		return
	}

	b = true
	return
}

// AddDatabaseField 添加数据库字段
func (s *sTSLTable) AddDatabaseField(ctx context.Context, tableName, fieldName string, dataType string, len int) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tableName = comm.ProductTableName(tableName)

	sql := fmt.Sprintf("ALTER STABLE %s.%s ADD COLUMN %s", dbName, tableName, s.column(dataType, fieldName, "", len))
	if _, err = taos.Exec(sql); err != nil {
		return
	}
	// 属性上报时间
	sql = fmt.Sprintf("ALTER STABLE %s.%s ADD COLUMN %s", dbName, tableName, s.column("date", fieldName+"_time", "", 0))
	_, err = taos.Exec(sql)

	return
}

// DelDatabaseField 删除数据库字段
func (s *sTSLTable) DelDatabaseField(ctx context.Context, tableName, fieldName string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tableName = comm.ProductTableName(tableName)
	fieldName = comm.TsdColumnName(fieldName)

	sql := fmt.Sprintf("ALTER STABLE %s.%s DROP COLUMN %s", dbName, tableName, fieldName)
	if _, err = taos.Exec(sql); err != nil {
		return
	}
	// 属性上报时间
	sql = fmt.Sprintf("ALTER STABLE %s.%s DROP COLUMN %s", dbName, tableName, fieldName+"_time")
	_, err = taos.Exec(sql)

	return
}

// ModifyDatabaseField 修改数据库指定字段长度
func (s *sTSLTable) ModifyDatabaseField(ctx context.Context, tableName, fieldName string, dataType string, len int) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tableName = comm.ProductTableName(tableName)

	sql := fmt.Sprintf("ALTER STABLE %s.%s MODIFY COLUMN %s", dbName, tableName, s.column(dataType, fieldName, "", len))
	_, err = taos.Exec(sql)
	if err != nil {
		err = gerror.New("设置字段长度失败,长度只能增大不能缩小")
		return
	}
	return
}

// AddTag 添加标签
func (s *sTSLTable) AddTag(ctx context.Context, tableName, tagName string, dataType string, len int) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tableName = comm.ProductTableName(tableName)

	sql := fmt.Sprintf("ALTER STABLE %s.%s ADD TAG %s", dbName, tableName, s.column(dataType, tagName, "", len, 1))
	_, err = taos.Exec(sql)

	return
}

// DelTag 删除标签
func (s *sTSLTable) DelTag(ctx context.Context, tableName, tagName string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tableName = comm.ProductTableName(tableName)
	tagName = comm.TsdTagName(tagName)

	sql := fmt.Sprintf("ALTER STABLE %s.%s DROP TAG %s", dbName, tableName, tagName)
	_, err = taos.Exec(sql)

	return
}

// ModifyTag 修改标签
func (s *sTSLTable) ModifyTag(ctx context.Context, tableName, tagName string, dataType string, len int) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}

	tableName = comm.ProductTableName(tableName)

	sql := fmt.Sprintf("ALTER STABLE %s.%s MODIFY TAG %s", dbName, tableName, s.column(dataType, tagName, "", len, 1))
	_, err = taos.Exec(sql)
	if err != nil {
		err = gerror.New("设置标签长度失败,长度只能增大不能缩小")
		return
	}
	return
}
