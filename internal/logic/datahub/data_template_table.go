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
func getTplTableName(id uint64) string {
	return "data_template_" + strconv.FormatUint(id, 10)
}

// 生成数据表结构
func createTplTable(ctx context.Context, id uint64) (table string, err error) {
	table = getTplTableName(id)

	isExsit := checkTableExsit(ctx, table)
	if isExsit {
		err = gerror.New("数据表" + table + "已存在")
		return
	}

	// 获取节点
	tplNodes, err := service.DataTemplateNode().List(ctx, id)
	if err != nil {
		return
	}
	if len(tplNodes) == 0 {
		err = gerror.New("该数据模型还未创建模型节点")
		return
	}

	pk := ""
	columns := make([]string, len(tplNodes)+1)
	for i, v := range tplNodes {
		// 节点默认值，主要处理自动生成的节点
		df := "DEFAULT 0"
		if v.DataType == "string" {
			df = "DEFAULT ''"
		}
		if v.From == 1 && v.Default != "" {
			df = "DEFAULT '" + v.Default + "'"
		}

		switch v.DataType {
		case "int":
			columns[i] = "`" + v.Key + "` int(11) " + df + " COMMENT '" + v.Name + "'"
		case "long":
			columns[i] = "`" + v.Key + "` bigint(20) " + df + " COMMENT '" + v.Name + "'"
		case "float":
			columns[i] = "`" + v.Key + "` float " + df + " COMMENT '" + v.Name + "'"
		case "double":
			columns[i] = "`" + v.Key + "` double " + df + " COMMENT '" + v.Name + "'"
		case "string":
			columns[i] = "`" + v.Key + "` varchar(255) " + df + " COMMENT '" + v.Name + "'"
		case "boolean":
			columns[i] = "`" + v.Key + "` tinyint " + df + " COMMENT '" + v.Name + "'"
		case "date":
			columns[i] = "`" + v.Key + "` datetime DEFAULT NULL COMMENT '" + v.Name + "'"
		default:
			columns[i] = "`" + v.Key + "` varchar(255) " + df + " COMMENT '" + v.Name + "'"
		}

		// 主键
		if v.IsPk == 1 {
			pk = ", primary key (`" + v.Key + "`)"
		}
	}
	columns[len(tplNodes)] = "`created_at` datetime DEFAULT NULL COMMENT '创建时间'"

	sql := "CREATE TABLE " + table + " ( " + strings.Join(columns, ",") + pk + " );"
	_, err = g.DB(DataCenter()).Exec(ctx, sql)
	if err != nil {
		return
	}

	return
}

// 表删除
func dropTplTable(ctx context.Context, id uint64) error {
	table := getTplTableName(id)

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
func addTplColumn(ctx context.Context, nodeId uint64) (err error) {
	var p *entity.DataTemplateNode
	err = dao.DataTemplateNode.Ctx(ctx).Where(dao.DataTemplateNode.Columns().Id, nodeId).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("模型节点不存在")
	}

	table := getTplTableName(p.Tid)

	isExsit := checkTableExsit(ctx, table)
	if !isExsit {
		return gerror.New("数据表" + table + "不存在")
	}

	// 节点默认值，主要处理自动生成的节点
	df := "DEFAULT 0"
	if p.DataType == "string" {
		df = "DEFAULT ''"
	}
	if p.From == 1 && p.Default != "" {
		df = "DEFAULT '" + p.Default + "'"
	}

	column := ""
	switch p.DataType {
	case "int":
		column = "`" + p.Key + "` int(11) " + df + " COMMENT '" + p.Name + "'"
	case "long":
		column = "`" + p.Key + "` bigint(20) " + df + " COMMENT '" + p.Name + "'"
	case "float":
		column = "`" + p.Key + "` float " + df + " COMMENT '" + p.Name + "'"
	case "double":
		column = "`" + p.Key + "` double " + df + " COMMENT '" + p.Name + "'"
	case "string":
		column = "`" + p.Key + "` varchar(255) " + df + " COMMENT '" + p.Name + "'"
	case "boolean":
		column = "`" + p.Key + "` tinyint " + df + " COMMENT '" + p.Name + "'"
	case "date":
		column = "`" + p.Key + "` datetime DEFAULT NULL COMMENT '" + p.Name + "'"
	default:
		column = "`" + p.Key + "` varchar(255) " + df + " COMMENT '" + p.Name + "'"
	}

	sql := "ALTER TABLE " + table + " ADD COLUMN " + column
	_, err = g.DB(DataCenter()).Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

// 字段删除
func dropTplColumn(ctx context.Context, nodeId uint64) (err error) {
	var p *entity.DataTemplateNode
	err = dao.DataTemplateNode.Ctx(ctx).Where(dao.DataTemplateNode.Columns().Id, nodeId).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("数据节点不存在")
	}

	table := getTplTableName(p.Tid)

	isExsit := checkTableExsit(ctx, table)
	if !isExsit {
		return gerror.New("数据表" + table + "不存在")
	}

	sql := "ALTER TABLE " + table + " DROP `" + p.Key + "`"
	_, err = g.DB(DataCenter()).Exec(ctx, sql)

	return err
}
