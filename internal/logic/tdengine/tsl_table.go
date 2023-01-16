package tdengine

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// 物模型 TDengine 表结构维护
type sTSLTable struct {
}

func init() {
	service.RegisterTSLTable(tslTableNew())
}

func tslTableNew() *sTSLTable {
	return &sTSLTable{}
}

// 数据入库
func (s *sTSLTable) Insert(ctx context.Context, deviceKey string, data map[string]any) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	if len(data) == 0 {
		return
	}

	ts := gtime.Now().Format("Y-m-d H:i:s")

	var (
		field = []string{"ts"}
		value = []string{"'" + ts + "'"}
	)
	for k, v := range data {
		field = append(field, k)
		value = append(value, "'"+gvar.New(v).String()+"'")
	}

	sql := "INSERT INTO ? (?) VALUES (?)"
	_, err = taos.Exec(sql, deviceKey, strings.Join(field, ","), strings.Join(value, ","))

	return
}

// 添加超级表
func (s *sTSLTable) CreateStable(ctx context.Context, tsl *model.TSL) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, dbName)
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	// 属性字段
	columns := []string{"ts TIMESTAMP"}
	for _, v := range tsl.Properties {
		maxLength := 0
		if v.ValueType.TSLParamBase.MaxLength != nil {
			maxLength = *v.ValueType.TSLParamBase.MaxLength
		}
		columns = append(columns, s.column(v.ValueType.Type, v.Key, v.Name, maxLength))
	}

	// 标签字段
	tags := make([]string, len(tsl.Tags)+1)
	tags[0] = "device VARCHAR(255) COMMENT '设备标识'"
	for i, v := range tsl.Tags {
		maxLength := 0
		if v.ValueType.TSLParamBase.MaxLength != nil {
			maxLength = *v.ValueType.TSLParamBase.MaxLength
		}
		tags[i+1] = s.column(v.ValueType.Type, v.Key, v.Name, maxLength)
	}

	tConent := ""
	if len(tags) > 0 {
		tConent = fmt.Sprintf("TAGS (%s)", strings.Join(tags, ","))
	}
	sql := fmt.Sprintf("CREATE STABLE %s.%s (%s) %s", dbName, tsl.Key, strings.Join(columns, ","), tConent)

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
	defer taos.Close()

	sql := fmt.Sprintf("CREATE TABLE %s USING %s (device) TAGS ('%s')", table, stable, table)
	println(sql)

	_, err = taos.Exec(sql)

	return
}

func (s *sTSLTable) column(dataType, key, name string, maxLength int) string {
	column := ""
	comment := ""
	if name != "" {
		comment = "COMMENT '" + name + "'"
	}
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
			maxLength = 255
		}
		tdType = "NCHAR(" + strconv.Itoa(maxLength) + ")"
	case "boolean":
		tdType = "BOOL"
	case "date":
		tdType = "TIMESTAMP"
	default:
		if maxLength == 0 {
			maxLength = 255
		}
		tdType = "NCHAR(" + strconv.Itoa(maxLength) + ")"
	}
	column = fmt.Sprintf("%s %s %s", key, tdType, comment)
	return column
}

// 删除超级表
func (s *sTSLTable) DropStable(ctx context.Context, table string) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	sql := fmt.Sprintf("DROP STABLE IF EXISTS %s.%s", dbName, table)
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
	defer taos.Close()

	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", dbName, table)
	_, err = taos.Exec(sql)

	return
}

// 创建数据库
func (s *sTSLTable) CreateDatabase(ctx context.Context) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	var name string
	taos.QueryRow("SELECT name FROM information_schema.ins_databases WHERE name = '?' LIMIT 1", dbName).Scan(&name)
	if name != "" {
		return
	}

	_, err = taos.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)

	return
}

// AddDatabaseField 添加数据库字段
func (s *sTSLTable) AddDatabaseField(ctx context.Context, tableName, fieldName string, dataType string, len int) (err error) {
	taos, err := service.TdEngine().GetConn(ctx, "")
	if err != nil {
		err = gerror.New("获取链接失败")
		return
	}
	defer taos.Close()

	sql := fmt.Sprintf("ALTER STABLE %s.%s ADD COLUMN %s", dbName, tableName, s.column(dataType, fieldName, "", len))
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
	defer taos.Close()

	sql := fmt.Sprintf("ALTER STABLE %s.%s DROP COLUMN %s", dbName, tableName, fieldName)
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
	defer taos.Close()

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
	defer taos.Close()

	sql := fmt.Sprintf("ALTER STABLE %s.%s ADD TAG %s", dbName, tableName, s.column(dataType, tagName, "", len))
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
	defer taos.Close()

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
	defer taos.Close()

	sql := fmt.Sprintf("ALTER STABLE %s.%s MODIFY TAG %s", dbName, tableName, s.column(dataType, tagName, "", len))
	_, err = taos.Exec(sql)
	if err != nil {
		err = gerror.New("设置标签长度失败,长度只能增大不能缩小")
		return
	}
	return
}
