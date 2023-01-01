package datahub

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sDataSourceRecord struct{}

func init() {
	service.RegisterDataSourceRecord(dataSourceRecordNew())
}

func dataSourceRecordNew() *sDataSourceRecord {
	return &sDataSourceRecord{}
}

func (s *sDataSourceRecord) GetForTpl(ctx context.Context, sourceId uint64, tid uint64) (rs gdb.Result, err error) {
	ds, err := service.DataSource().Detail(ctx, sourceId)
	if err != nil {
		return
	}
	if ds == nil {
		err = gerror.New("数据源不存在")
		return
	}

	dsNodes, err := service.DataNode().List(ctx, sourceId)
	if err != nil {
		return
	}
	if len(dsNodes) == 0 {
		err = gerror.New("数据源未创建节点")
		return
	}

	dt, err := service.DataTemplate().Detail(ctx, tid)
	if err != nil {
		return
	}
	if dt == nil {
		err = gerror.New("数据模型不存在")
		return
	}

	tplNodes, err := service.DataTemplateNode().List(ctx, tid)
	if err != nil {
		return
	}
	if len(tplNodes) == 0 {
		err = gerror.New("数据模型未创建模型节点")
		return
	}

	switch ds.From {
	case model.DataSourceFromApi:
		rs, err = s.getFromApi(ctx, ds, dsNodes, dt, tplNodes)
	case model.DataSourceFromDb:
		rs, err = s.getFromDb(ctx, ds, dsNodes, dt, tplNodes)
	case model.DataSourceFromDevice:
		rs, err = s.getFromDevice(ctx, ds, dsNodes, dt, tplNodes)
	}

	return
}

func (s *sDataSourceRecord) getFromDevice(ctx context.Context, ds *model.DataSourceOutput, dsNodes []*model.DataNodeOutput, dt *model.DataTemplateOutput, tplNodes []*model.DataTemplateNodeOutput) (rs gdb.Result, err error) {
	table := ds.DeviceConfig.ProductKey
	where := "where device='" + ds.DeviceConfig.DeviceKey + "'"

	// TDengine
	sql := fmt.Sprintf("select * from %s %s order by ts desc limit 1", table, where)
	data, err := service.TdEngine().GetAll(ctx, sql)
	if err != nil || data.Len() == 0 {
		return
	}

	r := make(gdb.Record, len(dsNodes))
	for _, v := range dsNodes {
		r[v.Key] = data[0][v.Value]
	}
	rs = append(rs, r)

	return
}

func (s *sDataSourceRecord) getFromApi(ctx context.Context, ds *model.DataSourceOutput, dsNodes []*model.DataNodeOutput, dt *model.DataTemplateOutput, tplNodes []*model.DataTemplateNodeOutput) (rs gdb.Result, err error) {
	table := getTableName(ds.SourceId)

	groupNum := len(ds.ApiConfig.RequestParams)
	if groupNum > 1 {
		// 多组数据
		sql := fmt.Sprintf("select * from %s order by created_at desc limit %d", table, groupNum)
		rs, err = g.DB(DataCenter()).GetAll(ctx, sql)
		return
	}

	sql := fmt.Sprintf("select * from %s order by created_at desc limit 1", table)
	rs, err = g.DB(DataCenter()).GetAll(ctx, sql)
	return
}

func (s *sDataSourceRecord) getFromDb(ctx context.Context, ds *model.DataSourceOutput, dsNodes []*model.DataNodeOutput, dt *model.DataTemplateOutput, tplNodes []*model.DataTemplateNodeOutput) (rs gdb.Result, err error) {
	table := getTableName(ds.SourceId)

	sql := fmt.Sprintf("select * from %s order by created_at desc limit %d", table, ds.DbConfig.Num)
	rs, err = g.DB(DataCenter()).GetAll(ctx, sql)
	return
}
