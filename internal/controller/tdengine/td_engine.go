package tdengine

import (
	"context"
	"sagooiot/api/v1/tdengine"
	"sagooiot/internal/model"
	"sagooiot/pkg/tsd"
)

var TdEngine = cTdEngine{}

type cTdEngine struct{}

// GetTdEngineAllDb 获取所有数据库
func (a *cTdEngine) GetTdEngineAllDb(ctx context.Context, req *tdengine.GetTdEngineAllDbReq) (res *tdengine.GetTdEngineAllDbRes, err error) {
	tsdDb := tsd.DB()
	defer tsdDb.Close()

	data, err := tsdDb.GetAllDatabaseName(ctx)
	if data != nil {
		res = &tdengine.GetTdEngineAllDbRes{
			Info: data,
		}
	}
	return
}

// GetListTableByDatabases 获取指定数据库下所有的表列表
func (a *cTdEngine) GetListTableByDatabases(ctx context.Context, req *tdengine.GetTdEngineListTableByDatabasesReq) (res *tdengine.GetTdEngineListTableByDatabasesRes, err error) {
	tsdDb := tsd.DB()
	defer tsdDb.Close()

	rs, err := tsdDb.GetTableListByDatabase(ctx, req.DbName)
	if err != nil {
		return
	}
	var data []*model.TDEngineTablesList
	if len(rs) > 0 {
		for _, v := range rs {
			data = append(data, &model.TDEngineTablesList{
				DbName:     v.DbName,
				StableName: v.StableName,
				TableName:  v.TableName,
				CreateTime: v.CreateTime,
			})
		}
		res = &tdengine.GetTdEngineListTableByDatabasesRes{
			Info: data,
		}
	}
	return
}

// GetTdEngineTableInfoByTable 获取指定数据表结构信息
func (a *cTdEngine) GetTdEngineTableInfoByTable(ctx context.Context, req *tdengine.GetTdEngineTableInfoByTableReq) (res *tdengine.GetTdEngineTableInfoByTableRes, err error) {
	tsdDb := tsd.DB()
	defer tsdDb.Close()

	rs, err := tsdDb.GetTableInfo(ctx, req.TableName)
	if err != nil {
		return
	}
	var data []*model.TDEngineTableInfo
	if len(rs) > 0 {
		for _, v := range rs {
			data = append(data, &model.TDEngineTableInfo{
				Field:  v.Field,
				Type:   v.Type,
				Length: v.Length,
				Note:   v.Note,
			})
		}
		res = &tdengine.GetTdEngineTableInfoByTableRes{
			Info: data,
		}
	}
	return
}

// GetTdEngineTableDataByTable 获取指定表数据信息
func (a *cTdEngine) GetTdEngineTableDataByTable(ctx context.Context, req *tdengine.GetTdEngineTableDataByTableReq) (res *tdengine.GetTdEngineTableDataByTableRes, err error) {
	tsdDb := tsd.DB()
	defer tsdDb.Close()

	table, err := tsdDb.GetTableData(ctx, req.TableName)
	if err != nil {
		return
	}
	var data *model.TableDataInfo
	if table != nil {
		data = &model.TableDataInfo{
			Filed: table.Filed,
			Info:  table.Info,
		}
		res = &tdengine.GetTdEngineTableDataByTableRes{
			TableDataInfo: data,
		}
	}
	return
}
