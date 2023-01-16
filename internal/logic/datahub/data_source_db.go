package datahub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"reflect"
	"strings"
	"sync"

	_ "github.com/gogf/gf/contrib/drivers/mssql/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

// 添加数据库数据源
func (s *sDataSource) AddDb(ctx context.Context, in *model.DataSourceDbAddInput) (sourceId uint64, err error) {
	id, _ := dao.DataSource.Ctx(ctx).
		Fields(dao.DataSource.Columns().SourceId).
		Where(dao.DataSource.Columns().Key, in.Key).
		Value()
	if id.Int64() > 0 {
		err = gerror.New("数据源标识重复")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataSource
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	param.Status = 0
	param.LockKey = 0

	conf, err := json.Marshal(in.Config)
	if err != nil {
		err = gerror.New("数据源配置格式错误")
		return
	}
	param.Config = conf

	if in.Rule != nil {
		rule, err := json.Marshal(in.Rule)
		if err != nil {
			return 0, gerror.New("规则配置格式错误")
		}
		param.Rule = rule
	}

	rs, err := dao.DataSource.Ctx(ctx).Data(param).Insert()
	if err != nil {
		return
	}

	newId, _ := rs.LastInsertId()
	sourceId = uint64(newId)

	return
}

// 编辑数据库数据源
func (s *sDataSource) EditDb(ctx context.Context, in *model.DataSourceDbEditInput) (err error) {
	out, err := s.Detail(ctx, in.SourceId)
	if err != nil {
		return err
	}
	if out == nil {
		return gerror.New("数据源不存在")
	}
	if out.Status == model.DataSourceStatusOn {
		return gerror.New("数据源已发布")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataSource
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.SourceId = nil
	if out.LockKey == 1 {
		param.Key = nil
	} else {
		id, _ := dao.DataSource.Ctx(ctx).
			Fields(dao.DataSource.Columns().SourceId).
			Where(dao.DataSource.Columns().Key, in.Key).
			WhereNot(dao.DataSource.Columns().SourceId, in.SourceId).
			Value()
		if id.Int64() > 0 {
			err = gerror.New("数据源标识重复")
			return
		}
	}

	conf, err := json.Marshal(in.Config)
	if err != nil {
		return gerror.New("数据源配置格式错误")
	}
	param.Config = conf

	if in.Rule != nil {
		rule, err := json.Marshal(in.Rule)
		if err != nil {
			return gerror.New("规则配置格式错误")
		}
		param.Rule = rule
	}

	_, err = dao.DataSource.Ctx(ctx).Data(param).Where(dao.DataSource.Columns().SourceId, in.SourceId).Update()

	return
}

// 获取数据库表结构
func (s *sDataSource) GetDbFields(ctx context.Context, sourceId uint64) (g.MapStrAny, error) {
	p, _ := s.Detail(ctx, sourceId)
	if p == nil || p.DbConfig == nil {
		return nil, gerror.New("数据源不存在或未进行配置")
	}

	// 数据库配置
	conf := p.DbConfig

	db, err := gdb.New(gdb.ConfigNode{
		Host: conf.Host,
		Port: gconv.String(conf.Port),
		User: conf.User,
		Pass: conf.Passwd,
		Name: conf.DbName,
		Type: conf.Type,
	})
	if err != nil {
		return nil, err
	}

	if conf.QueryType == model.DataSourceDbQueryTypeSql {
		rs, err := db.GetOne(ctx, conf.TableName)
		if err != nil {
			return nil, err
		}

		tmp := rs.GMap()
		data := make(g.MapStrAny)
		for k := range tmp.Map() {
			tf := new(gdb.TableField)
			tf.Name = k
			tf.Comment = k

			switch (tmp.GetVar(k).Interface()).(type) {
			case int:
				tf.Type = "int"
			case string:
				tf.Type = "string"
			case float64:
				tf.Type = "double"
			case *gtime.Time:
				tf.Type = "date"
			default:
				tf.Type = "string"
			}
			data[k] = tf
		}
		return data, nil
	}

	fields, err := db.TableFields(ctx, conf.TableName)
	if err != nil {
		return nil, err
	}

	// 字段类型转换
	data := make(g.MapStrAny, len(fields))
	for k, v := range fields {
		p := *v
		vtype := strings.Split(p.Type, "(")
		p.Type = mappingFieldType(vtype[0])
		data[k] = p
	}
	return data, nil
}

// mssql字段类型映射
func mappingFieldType(t string) (netT string) {
	switch t {
	case "int", "smallint", "tinyint", "timestamp":
		netT = "int"
	case "bigint":
		netT = "long"
	case "float", "decimal", "numeric":
		netT = "float"
	case "real":
		netT = "double"
	case "varchar", "char", "nvarchar", "nchar", "bit":
		netT = "string"
	case "text", "ntext":
		netT = "text"
	case "datetime", "smalldatetime":
		netT = "date"
	default:
		netT = "string"
	}
	return
}

// 获取数据库单条数据
func (s *sDataSource) GetDbData(ctx context.Context, sourceId uint64) (string, error) {
	rs, _, err := s.getDbData(ctx, sourceId, 1)
	if err != nil || rs.Len() == 0 {
		return "", err
	}
	return rs[0].Json(), nil
}

// 获取数据源配置的数据库数据
// 数据源配置时，注意控制数据获取的数量，数量过大可能造成内存溢出
func (s *sDataSource) getDbData(ctx context.Context, sourceId uint64, limit int) (rs gdb.Result, ds *model.DataSourceOutput, err error) {
	p, _ := s.Detail(ctx, sourceId)
	if p == nil || p.DbConfig == nil {
		err = gerror.New("数据源不存在或未进行配置")
		return
	}
	ds = p

	// 数据库配置
	conf := p.DbConfig

	dbDebug, _ := g.Cfg().Get(context.TODO(), "database.default.debug")

	db, err := gdb.New(gdb.ConfigNode{
		Host:  conf.Host,
		Port:  gconv.String(conf.Port),
		User:  conf.User,
		Pass:  conf.Passwd,
		Name:  conf.DbName,
		Type:  conf.Type,
		Debug: dbDebug.Bool(),
	})
	if err != nil {
		return
	}

	if conf.QueryType == model.DataSourceDbQueryTypeSql {
		sql := conf.TableName
		if strings.Contains(sql, "%d") {
			sql = fmt.Sprintf(conf.TableName, conf.PkMax)
		}
		rs, err = db.GetAll(ctx, sql)
		return
	}

	where := ""
	if limit <= 0 {
		limit = conf.Num

		if conf.Pk != "" {
			// 判断主键类型
			var tmap g.MapStrAny
			tmap, err = s.GetDbFields(ctx, sourceId)
			if err != nil {
				return
			}
			if t, ok := tmap[conf.Pk]; ok {
				tf := t.(gdb.TableField)
				switch tf.Type {
				case "int", "long":
					where = fmt.Sprintf("where %s > %d", conf.Pk, conf.PkMax)
				}
			}
		}
	}
	order := ""
	if conf.Pk != "" {
		order = fmt.Sprintf("order by %s asc", conf.Pk)
	}
	sql := fmt.Sprintf("select * from %s %s %s limit %d", conf.TableName, where, order, limit)
	if conf.Type == "mssql" {
		sql = fmt.Sprintf("select top %d * from %s %s %s", limit, conf.TableName, where, order)
	}
	rs, err = db.GetAll(ctx, sql)
	return
}

// 数据库源数据记录列表
func (s *sDataSource) getDbDataRecord(ctx context.Context, in *model.DataSourceDataInput, ds *model.DataSourceOutput) (out *model.DataSourceDataOutput, err error) {
	out = new(model.DataSourceDataOutput)

	table := getTableName(in.SourceId)

	// 搜索条件
	var exp []string
	var value []any
	if in.Param != nil {
		for k, v := range in.Param {
			exp = append(exp, k)
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					ele := s.Index(i)
					value = append(value, ele.Interface())
				}
			} else {
				value = append(value, v)
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " where " + strings.Join(exp, " and ")
	}
	sql := "select * from " + table + where + " order by created_at desc"

	out.Total, _ = g.DB(DataCenter()).GetCount(ctx, sql, value...)
	out.CurrentPage = in.PageNum

	sql += fmt.Sprintf(" limit %d, %d", (in.PageNum-1)*in.PageSize, in.PageSize)
	rs, err := g.DB(DataCenter()).GetAll(ctx, sql, value...)
	if err != nil {
		return
	}
	out.List = rs.Json()

	return
}

// 数据库源数据记录列表，非分页
func (s *sDataSource) getDbDataAllRecord(ctx context.Context, in *model.SourceDataAllInput, ds *model.DataSourceOutput) (out *model.SourceDataAllOutput, err error) {
	out = new(model.SourceDataAllOutput)

	table := getTableName(in.SourceId)

	// 搜索条件
	var exp []string
	var value []any
	if in.Param != nil {
		for k, v := range in.Param {
			exp = append(exp, k)
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				for i := 0; i < s.Len(); i++ {
					ele := s.Index(i)
					value = append(value, ele.Interface())
				}
			} else {
				value = append(value, v)
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " where " + strings.Join(exp, " and ")
	}
	sql := "select * from " + table + where + " order by created_at desc"
	rs, err := g.DB(DataCenter()).GetAll(ctx, sql, value...)
	if err != nil {
		return
	}
	out.List = rs.Json()

	return
}

// 复制数据库数据源
func (s *sDataSource) copeDbSource(ctx context.Context, ds *model.DataSourceOutput) (newSourceId uint64, err error) {
	err = dao.DataSource.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 复制源
		in := new(model.DataSourceDbAddInput)
		in.DataSource = model.DataSource{}
		in.Config = model.DataSourceConfigDb{}

		err = gconv.Scan(ds.DataSource, &in.DataSource)
		if err != nil {
			return err
		}
		in.DataSource.Key += "_" + gtime.Now().Format("YmdHis")

		err = gconv.Scan(ds.DbConfig, &in.Config)
		if err != nil {
			return err
		}

		newSourceId, err = s.AddDb(ctx, in)
		if err != nil {
			return err
		}

		// 复制节点
		err = s.copeNode(ctx, ds.SourceId, newSourceId)

		return err
	})

	return
}

func (s *sDataSource) updateDataForDb(ctx context.Context, sourceId uint64) error {
	// 数据节点
	nodeList, err := service.DataNode().List(ctx, sourceId)
	if err != nil {
		return err
	}
	if len(nodeList) == 0 {
		return gerror.New("数据源未添加数据节点")
	}

	// 获取数据库数据
	rs, ds, err := s.getDbData(ctx, sourceId, -1)
	if err != nil {
		return err
	}

	var pkMax uint64 = 0

	// 数据映射
	var insertData []*gmap.AnyAnyMap
	for _, row := range rs {
		pk := row[ds.DbConfig.Pk]
		if pk.Uint64() > pkMax {
			pkMax = pk.Uint64()
		}

		m := gmap.New(true)

		var wg sync.WaitGroup
		for _, v := range nodeList {
			wg.Add(1)
			go func(v *model.DataNodeOutput) {
				defer wg.Done()

				// 规则过滤，启用节点规则
				var rule *model.DataSourceRule
				if len(v.NodeRule) > 0 {
					rule = v.NodeRule[0]
				}

				sv, ok := row[v.Value]
				if !ok {
					return
				}

				rs := sv.String()
				if rule != nil && rule.Expression != "" {
					// 正则过滤数据
					rs, err = gregex.ReplaceString(rule.Expression, rule.Replace, rs)
					if err != nil {
						return
					}
				}
				m.Set(v.Key, rs)
			}(v)
		}
		wg.Wait()
		insertData = append(insertData, m)
	}

	// 入库
	if len(insertData) > 0 {
		table := getTableName(sourceId)
		err = g.DB(DataCenter()).GetCore().ClearTableFields(ctx, table)
		if err != nil {
			return err
		}
		_, err = g.DB(DataCenter()).Save(ctx, table, insertData)
		if err != nil {
			return err
		}
	}

	// 主键最大值存储
	if pkMax > 0 {
		conf := ds.DbConfig
		conf.PkMax = pkMax
		c, _ := json.Marshal(conf)
		_, err = dao.DataSource.Ctx(ctx).
			Data(g.Map{
				dao.DataSource.Columns().Config: c,
			}).
			Where(dao.DataSource.Columns().SourceId, sourceId).
			Update()
	}
	return err
}
