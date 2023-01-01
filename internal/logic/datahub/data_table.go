package datahub

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 数据表名称
func getTableName(sourceId uint64) string {
	return "data_source_" + strconv.FormatUint(sourceId, 10)
}

// 生成数据表结构
func createTable(ctx context.Context, sourceId uint64) (table string, err error) {
	table = getTableName(sourceId)

	isExsit := checkTableExsit(ctx, table)
	if isExsit {
		err = gerror.New("数据表" + table + "已存在")
		return
	}

	// 获取节点
	nodeList, err := service.DataNode().List(ctx, sourceId)
	if err != nil {
		return
	}
	if len(nodeList) == 0 {
		err = gerror.New("该数据源还未创建数据节点")
		return
	}

	pk := ""
	columns := make([]string, len(nodeList)+1)
	for i, v := range nodeList {
		switch v.DataType {
		case "int":
			columns[i] = "`" + v.Key + "` int(11) DEFAULT 0 COMMENT '" + v.Name + "'"
		case "long":
			columns[i] = "`" + v.Key + "` bigint(20) DEFAULT 0 COMMENT '" + v.Name + "'"
		case "float":
			columns[i] = "`" + v.Key + "` float DEFAULT 0 COMMENT '" + v.Name + "'"
		case "double":
			columns[i] = "`" + v.Key + "` double DEFAULT 0 COMMENT '" + v.Name + "'"
		case "string":
			columns[i] = "`" + v.Key + "` varchar(255) DEFAULT '' COMMENT '" + v.Name + "'"
		case "boolean":
			columns[i] = "`" + v.Key + "` tinyint DEFAULT 0 COMMENT '" + v.Name + "'"
		case "date":
			columns[i] = "`" + v.Key + "` datetime DEFAULT NULL COMMENT '" + v.Name + "'"
		default:
			columns[i] = "`" + v.Key + "` varchar(255) DEFAULT '' COMMENT '" + v.Name + "'"
		}

		// 主键
		if v.IsPk == 1 {
			pk = ", primary key (`" + v.Key + "`)"
		}
	}
	columns[len(nodeList)] = "`created_at` datetime DEFAULT NULL COMMENT '创建时间'"

	sql := "CREATE TABLE " + table + " ( " + strings.Join(columns, ",") + pk + " );"
	_, err = g.DB(DataCenter()).Exec(ctx, sql)
	if err != nil {
		return
	}

	return
}

// 表删除
func dropTable(ctx context.Context, sourceId uint64) error {
	table := getTableName(sourceId)

	isExsit := checkTableExsit(ctx, table)
	if !isExsit {
		return nil
	}

	sql := "DROP TABLE " + table
	_, err := g.DB(DataCenter()).Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

// 字段增加
func addColumn(ctx context.Context, nodeId uint64) (err error) {
	var p *entity.DataNode
	err = dao.DataNode.Ctx(ctx).Where(dao.DataNode.Columns().NodeId, nodeId).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("数据节点不存在")
	}

	table := getTableName(p.SourceId)

	isExsit := checkTableExsit(ctx, table)
	if !isExsit {
		return gerror.New("数据表" + table + "不存在")
	}

	column := ""

	switch p.DataType {
	case "int":
		column = "`" + p.Key + "` int(11) DEFAULT 0 COMMENT '" + p.Name + "'"
	case "long":
		column = "`" + p.Key + "` bigint(20) DEFAULT 0 COMMENT '" + p.Name + "'"
	case "float":
		column = "`" + p.Key + "` float DEFAULT 0 COMMENT '" + p.Name + "'"
	case "double":
		column = "`" + p.Key + "` double DEFAULT 0 COMMENT '" + p.Name + "'"
	case "string":
		column = "`" + p.Key + "` varchar(255) DEFAULT '' COMMENT '" + p.Name + "'"
	case "boolean":
		column = "`" + p.Key + "` tinyint DEFAULT 0 COMMENT '" + p.Name + "'"
	case "date":
		column = "`" + p.Key + "` datetime DEFAULT NULL COMMENT '" + p.Name + "'"
	default:
		column = "`" + p.Key + "` varchar(255) DEFAULT '' COMMENT '" + p.Name + "'"
	}

	sql := "ALTER TABLE " + table + " ADD COLUMN " + column
	_, err = g.DB(DataCenter()).Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

// 字段删除
func dropColumn(ctx context.Context, nodeId uint64) (err error) {
	var p *entity.DataNode
	err = dao.DataNode.Ctx(ctx).Where(dao.DataNode.Columns().NodeId, nodeId).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("数据节点不存在")
	}

	table := getTableName(p.SourceId)

	isExsit := checkTableExsit(ctx, table)
	if !isExsit {
		return gerror.New("数据表" + table + "不存在")
	}

	sql := "ALTER TABLE " + table + " DROP `" + p.Key + "`"
	_, err = g.DB(DataCenter()).Exec(ctx, sql)

	return err
}

// 检查表是否存在, true:已存在，false:不存在
func checkTableExsit(ctx context.Context, table string) bool {
	sql := "select * from information_schema.tables where table_name = ? and table_schema = (select database()) limit 1"
	one, _ := g.DB(DataCenter()).GetOne(ctx, sql, table)
	return one != nil
}

// 获取数据库分组
func DataCenter() string {
	name, _ := g.Cfg().Get(context.TODO(), "database.dataCenter.link")
	if name.String() != "" {
		return "dataCenter"
	}
	return ""
}
