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

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// 添加设备数据源
func (s *sDataSource) AddDevice(ctx context.Context, in *model.DataSourceDeviceAddInput) (sourceId uint64, err error) {
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

// 编辑设备数据源
func (s *sDataSource) EditDevice(ctx context.Context, in *model.DataSourceDeviceEditInput) (err error) {
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

// 获取设备源数据，源配置测试使用
func (s *sDataSource) GetDeviceData(ctx context.Context, sourceId uint64) (string, error) {
	p, _ := s.Detail(ctx, sourceId)
	if p == nil || p.DeviceConfig == nil {
		return "", gerror.New("数据源不存在或未进行配置")
	}

	// 设备配置
	conf := p.DeviceConfig

	// TDengine
	sql := "select * from ? where device='?' order by ts desc limit 1"
	rs, err := service.TdEngine().GetOne(ctx, sql, conf.ProductKey, conf.DeviceKey)
	if err != nil {
		return "", err
	}
	data, _ := json.Marshal(rs)

	return string(data), nil
}

// 设备源数据记录列表
func (s *sDataSource) getDeviceDataRecord(ctx context.Context, in *model.DataSourceDataInput, ds *model.DataSourceOutput) (out *model.DataSourceDataOutput, err error) {
	out = new(model.DataSourceDataOutput)

	// 设备配置
	conf := ds.DeviceConfig

	// 搜索条件
	var exp []string
	if in.Param != nil {
		for k, v := range in.Param {
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				tmp := k
				for i := 0; i < s.Len(); i++ {
					s := fmt.Sprintf("'%v'", v)
					tmp = strings.Replace(tmp, "?", s, 1)
				}
				exp = append(exp, tmp)
			} else {
				s := fmt.Sprintf("'%v'", v)
				exp = append(exp, strings.Replace(k, "?", s, 1))
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " and " + strings.Join(exp, " and ")
	}

	// TDengine
	sql := fmt.Sprintf("select count(*) as num from %s where device='%s' %s limit 100", conf.ProductKey, conf.DeviceKey, where)
	rs, err := service.TdEngine().GetOne(ctx, sql)
	if err != nil {
		return
	}
	out.Total = rs["num"].Int()
	out.CurrentPage = in.PageNum
	if in.PageNum*in.PageSize > 100 {
		return
	}

	// 获取节点字段
	var fileds []string
	nodeList, err := service.DataNode().List(ctx, ds.SourceId)
	if err != nil {
		return
	}
	for _, v := range nodeList {
		fileds = append(fileds, v.Value)
	}

	sql = fmt.Sprintf("select %s from %s where device='%s' %s order by ts desc limit %d, %d", strings.Join(fileds, ","), conf.ProductKey, conf.DeviceKey, where, (in.PageNum-1)*in.PageSize, in.PageSize)
	data, err := service.TdEngine().GetAll(ctx, sql)
	if err != nil {
		return
	}

	var nr gdb.Result
	r := make(gdb.Record, len(nodeList))
	for _, row := range data {
		for _, v := range nodeList {
			r[v.Key] = row[v.Value]
		}
		nr = append(nr, r)
	}

	out.List = nr.Json()

	return
}

// 设备源数据记录列表，非分页
func (s *sDataSource) getDeviceDataAllRecord(ctx context.Context, in *model.SourceDataAllInput, ds *model.DataSourceOutput) (out *model.SourceDataAllOutput, err error) {
	out = new(model.SourceDataAllOutput)

	// 设备配置
	conf := ds.DeviceConfig

	// 搜索条件
	var exp []string
	if in.Param != nil {
		for k, v := range in.Param {
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				s := reflect.ValueOf(v)
				tmp := k
				for i := 0; i < s.Len(); i++ {
					s := fmt.Sprintf("'%v'", v)
					tmp = strings.Replace(tmp, "?", s, 1)
				}
				exp = append(exp, tmp)
			} else {
				s := fmt.Sprintf("'%v'", v)
				exp = append(exp, strings.Replace(k, "?", s, 1))
			}
		}
	}
	where := ""
	if len(exp) > 0 {
		where = " and " + strings.Join(exp, " and ")
	}

	// 获取节点字段
	var fileds []string
	nodeList, err := service.DataNode().List(ctx, ds.SourceId)
	if err != nil {
		return
	}
	for _, v := range nodeList {
		fileds = append(fileds, v.Value)
	}

	// TDengine
	sql := fmt.Sprintf("select %s from %s where device='%s' %s order by ts desc", strings.Join(fileds, ","), conf.ProductKey, conf.DeviceKey, where)
	data, err := service.TdEngine().GetAll(ctx, sql)
	if err != nil {
		return
	}
	out.List = data.Json()

	return
}

// 复制设备数据源
func (s *sDataSource) copeDeviceSource(ctx context.Context, ds *model.DataSourceOutput) (newSourcId uint64, err error) {
	err = dao.DataSource.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 复制源
		in := new(model.DataSourceDeviceAddInput)
		in.DataSource = model.DataSource{}
		in.Config = model.DataSourceConfigDevice{}

		err = gconv.Scan(ds.DataSource, &in.DataSource)
		if err != nil {
			return err
		}
		in.DataSource.Key += "_" + gtime.Now().Format("YmdHis")

		err = gconv.Scan(ds.DeviceConfig, &in.Config)
		if err != nil {
			return err
		}

		newSourcId, err = s.AddDevice(ctx, in)
		if err != nil {
			return err
		}

		// 复制节点
		err = s.copeNode(ctx, ds.SourceId, newSourcId)

		return err
	})

	return
}
