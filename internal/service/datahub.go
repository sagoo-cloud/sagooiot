// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"

	"github.com/go-gota/gota/dataframe"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IDataSource interface {
		Add(ctx context.Context, in *model.DataSourceApiAddInput) (sourceId uint64, err error)
		Edit(ctx context.Context, in *model.DataSourceApiEditInput) (err error)
		Del(ctx context.Context, ids []uint64) (err error)
		Search(ctx context.Context, in *model.DataSourceSearchInput) (out *model.DataSourceSearchOutput, err error)
		List(ctx context.Context) (list []*entity.DataSource, err error)
		Detail(ctx context.Context, sourceId uint64) (out *model.DataSourceOutput, err error)
		Deploy(ctx context.Context, sourceId uint64) (err error)
		Undeploy(ctx context.Context, sourceId uint64) (err error)
		UpdateData(ctx context.Context, sourceId uint64) (err error)
		GetData(ctx context.Context, in *model.DataSourceDataInput) (out *model.DataSourceDataOutput, err error)
		GetAllData(ctx context.Context, in *model.SourceDataAllInput) (out *model.SourceDataAllOutput, err error)
		AllSource(ctx context.Context) (out []*model.AllSourceOut, err error)
		CopeSource(ctx context.Context, sourceId uint64) (err error)
		UpdateInterval(ctx context.Context, sourceId uint64, cronExpression string) (err error)
		GetApiData(ctx context.Context, sourceId uint64) (apiData []string, err error)
		AddDb(ctx context.Context, in *model.DataSourceDbAddInput) (sourceId uint64, err error)
		EditDb(ctx context.Context, in *model.DataSourceDbEditInput) (err error)
		GetDbFields(ctx context.Context, sourceId uint64) (g.MapStrAny, error)
		GetDbData(ctx context.Context, sourceId uint64) (string, error)
		AddDevice(ctx context.Context, in *model.DataSourceDeviceAddInput) (sourceId uint64, err error)
		EditDevice(ctx context.Context, in *model.DataSourceDeviceEditInput) (err error)
		GetDeviceData(ctx context.Context, sourceId uint64) (string, error)
	}
	IDataSourceRecord interface {
		GetForTpl(ctx context.Context, sourceId uint64, tid uint64) (rs gdb.Result, err error)
	}
	IDataTemplate interface {
		Add(ctx context.Context, in *model.DataTemplateAddInput) (id uint64, err error)
		Edit(ctx context.Context, in *model.DataTemplateEditInput) (err error)
		Del(ctx context.Context, ids []uint64) (err error)
		Search(ctx context.Context, in *model.DataTemplateSearchInput) (out *model.DataTemplateSearchOutput, err error)
		List(ctx context.Context) (list []*entity.DataTemplate, err error)
		Detail(ctx context.Context, id uint64) (out *model.DataTemplateOutput, err error)
		Deploy(ctx context.Context, id uint64) (err error)
		Undeploy(ctx context.Context, id uint64) (err error)
		GetData(ctx context.Context, in *model.DataTemplateDataInput) (out *model.DataTemplateDataOutput, err error)
		GetAllData(ctx context.Context, in *model.TemplateDataAllInput) (out *model.TemplateDataAllOutput, err error)
		GetDataBySql(ctx context.Context, sql string) (df dataframe.DataFrame, err error)
		GetDataByTableName(ctx context.Context, tableName string) (df dataframe.DataFrame, err error)
		GetLastData(ctx context.Context, in *model.TemplateDataLastInput) (out *model.TemplateDataLastOutput, err error)
		UpdateData(ctx context.Context, id uint64) error
		GetInfoByIds(ctx context.Context, ids []uint64) (data []*entity.DataTemplate, err error)
		AllTemplate(ctx context.Context) (out []*model.AllTemplateOut, err error)
		UpdateInterval(ctx context.Context, id uint64, cronExpression string) (err error)
		CopeTemplate(ctx context.Context, id uint64) (err error)
		CheckRelation(ctx context.Context, id uint64) (yes bool, err error)
		SetRelation(ctx context.Context, in *model.TemplateDataRelationInput) (err error)
		SourceList(ctx context.Context, id uint64) (list []*model.DataSourceOutput, err error)
	}
	IDataTemplateBusi interface {
		Add(ctx context.Context, in *model.DataTemplateBusiAddInput) (err error)
		GetInfos(ctx context.Context, busiTypes int) (data *entity.DataTemplateBusi, err error)
		GetInfo(ctx context.Context, busiTypes int) (data *entity.DataTemplateBusi, err error)
		GetTable(ctx context.Context, busiTypes int) (table string, err error)
		Del(ctx context.Context, tid uint64) error
	}
	IDataTemplateNode interface {
		Add(ctx context.Context, in *model.DataTemplateNodeAddInput) (err error)
		Edit(ctx context.Context, in *model.DataTemplateNodeEditInput) (err error)
		Del(ctx context.Context, id uint64) (err error)
		List(ctx context.Context, tid uint64) (list []*model.DataTemplateNodeOutput, err error)
	}
	IDataTemplateRecord interface {
		UpdateData(ctx context.Context, tid uint64) error
	}
	IDataNode interface {
		Add(ctx context.Context, in *model.DataNodeAddInput) (err error)
		Edit(ctx context.Context, in *model.DataNodeEditInput) (err error)
		Del(ctx context.Context, nodeId uint64) (err error)
		List(ctx context.Context, sourceId uint64) (list []*model.DataNodeOutput, err error)
		Detail(ctx context.Context, nodeId uint64) (out *model.DataNodeOutput, err error)
	}
)

var (
	localDataNode           IDataNode
	localDataSource         IDataSource
	localDataSourceRecord   IDataSourceRecord
	localDataTemplate       IDataTemplate
	localDataTemplateBusi   IDataTemplateBusi
	localDataTemplateNode   IDataTemplateNode
	localDataTemplateRecord IDataTemplateRecord
)

func DataTemplateBusi() IDataTemplateBusi {
	if localDataTemplateBusi == nil {
		panic("implement not found for interface IDataTemplateBusi, forgot register?")
	}
	return localDataTemplateBusi
}

func RegisterDataTemplateBusi(i IDataTemplateBusi) {
	localDataTemplateBusi = i
}

func DataTemplateNode() IDataTemplateNode {
	if localDataTemplateNode == nil {
		panic("implement not found for interface IDataTemplateNode, forgot register?")
	}
	return localDataTemplateNode
}

func RegisterDataTemplateNode(i IDataTemplateNode) {
	localDataTemplateNode = i
}

func DataTemplateRecord() IDataTemplateRecord {
	if localDataTemplateRecord == nil {
		panic("implement not found for interface IDataTemplateRecord, forgot register?")
	}
	return localDataTemplateRecord
}

func RegisterDataTemplateRecord(i IDataTemplateRecord) {
	localDataTemplateRecord = i
}

func DataNode() IDataNode {
	if localDataNode == nil {
		panic("implement not found for interface IDataNode, forgot register?")
	}
	return localDataNode
}

func RegisterDataNode(i IDataNode) {
	localDataNode = i
}

func DataSource() IDataSource {
	if localDataSource == nil {
		panic("implement not found for interface IDataSource, forgot register?")
	}
	return localDataSource
}

func RegisterDataSource(i IDataSource) {
	localDataSource = i
}

func DataSourceRecord() IDataSourceRecord {
	if localDataSourceRecord == nil {
		panic("implement not found for interface IDataSourceRecord, forgot register?")
	}
	return localDataSourceRecord
}

func RegisterDataSourceRecord(i IDataSourceRecord) {
	localDataSourceRecord = i
}

func DataTemplate() IDataTemplate {
	if localDataTemplate == nil {
		panic("implement not found for interface IDataTemplate, forgot register?")
	}
	return localDataTemplate
}

func RegisterDataTemplate(i IDataTemplate) {
	localDataTemplate = i
}
