package datahub

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"reflect"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

func (s *sDataSource) updateDataForApi(ctx context.Context, sourceId uint64) error {
	// 数据节点
	nodeList, err := service.DataNode().List(ctx, sourceId)
	if err != nil {
		return err
	}
	if len(nodeList) == 0 {
		return gerror.New("数据源未添加数据节点")
	}

	// 获取api数据
	apiDataArr, err := s.GetApiData(ctx, sourceId)
	if err != nil {
		return err
	}

	// 数据映射
	var insertData []*gmap.AnyAnyMap
	for _, apiData := range apiDataArr {
		j := gjson.New(apiData)
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

				rs := getValue(v.Value, j, rule)
				if rs != "" {
					m.Set(v.Key, rs)
				}
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

	return nil
}

// 获取api数据
func (s *sDataSource) GetApiData(ctx context.Context, sourceId uint64) (apiData []string, err error) {
	p, _ := s.Detail(ctx, sourceId)
	if p == nil || p.ApiConfig == nil {
		err = gerror.New("数据源不存在或未进行配置")
		return
	}

	// api配置
	config := p.ApiConfig

	// api数据获取
	get := func(header, data g.MapStrStr) (string, error) {
		client := g.Client()
		client.Header(header)
		res, err := client.DoRequest(ctx, config.Method, config.Url, data)
		if err != nil {
			return "", err
		}
		defer res.Close()
		return res.ReadAllString(), nil
	}

	if len(config.RequestParams) > 0 {
		for _, param := range config.RequestParams {
			header := g.MapStrStr{}
			data := g.MapStrStr{}
			for _, v := range param {
				if v.Type == "header" {
					header[v.Key] = v.Value
				} else {
					data[v.Key] = v.Value
				}
			}
			rs, err := get(header, data)
			if err != nil {
				return nil, err
			}
			apiData = append(apiData, rs)
		}
	} else {
		rs, err := get(nil, nil)
		if err != nil {
			return nil, err
		}
		apiData = append(apiData, rs)
	}

	return
}

// api数据提取值
func getValue(value string, apiData *gjson.Json, rule *model.DataSourceRule) string {
	sv := apiData.Get(value)
	if sv == nil {
		return ""
	}

	s := sv.String()

	if rule != nil && rule.Expression != "" {
		// 正则过滤数据
		rs, _ := gregex.ReplaceString(rule.Expression, rule.Replace, s)
		return gconv.String(rs)
	}

	return s
}

// api源数据记录列表
func (s *sDataSource) getApiDataRecord(ctx context.Context, in *model.DataSourceDataInput, ds *model.DataSourceOutput) (out *model.DataSourceDataOutput, err error) {
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

// api源数据记录列表，非分页
func (s *sDataSource) getApiDataAllRecord(ctx context.Context, in *model.SourceDataAllInput, ds *model.DataSourceOutput) (out *model.SourceDataAllOutput, err error) {
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

// 复制API数据源
func (s *sDataSource) copeApiSource(ctx context.Context, ds *model.DataSourceOutput) (newSourceId uint64, err error) {
	err = dao.DataSource.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 复制源
		in := new(model.DataSourceApiAddInput)
		in.DataSource = model.DataSource{}
		in.Config = model.DataSourceConfigApi{}

		err = gconv.Scan(ds.DataSource, &in.DataSource)
		if err != nil {
			return err
		}
		in.DataSource.Key += "_" + gtime.Now().Format("YmdHis")

		err = gconv.Scan(ds.ApiConfig, &in.Config)
		if err != nil {
			return err
		}

		newSourceId, err = s.Add(ctx, in)
		if err != nil {
			return err
		}

		// 复制节点
		err = s.copeNode(ctx, ds.SourceId, newSourceId)

		return err
	})

	return
}
