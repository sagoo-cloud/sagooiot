package tdengine

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/tdengine"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var TdEngine = cTdEngine{}

type cTdEngine struct{}

// GetTdEngineAllDb 获取所有数据库
func (a *cTdEngine) GetTdEngineAllDb(ctx context.Context, req *tdengine.GetTdEngineAllDbReq) (res *tdengine.GetTdEngineAllDbRes, err error) {
	data, err := service.TdEngine().GetTdEngineAllDb(ctx)
	if data != nil {
		res = &tdengine.GetTdEngineAllDbRes{
			Info: data,
		}
	}
	return
}

// GetListTableByDatabases 获取指定数据库下所有的表列表
func (a *cTdEngine) GetListTableByDatabases(ctx context.Context, req *tdengine.GetTdEngineListTableByDatabasesReq) (res *tdengine.GetTdEngineListTableByDatabasesRes, err error) {
	data, err := service.TdEngine().GetListTableByDatabases(ctx, req.DbName)
	if data != nil {
		res = &tdengine.GetTdEngineListTableByDatabasesRes{
			Info: data,
		}
	}
	return
}

// GetTdEngineTableInfoByTable 获取指定数据表结构信息
func (a *cTdEngine) GetTdEngineTableInfoByTable(ctx context.Context, req *tdengine.GetTdEngineTableInfoByTableReq) (res *tdengine.GetTdEngineTableInfoByTableRes, err error) {
	data, err := service.TdEngine().GetTdEngineTableInfoByTable(ctx, req.DbName, req.TableName)
	if data != nil {
		res = &tdengine.GetTdEngineTableInfoByTableRes{
			Info: data,
		}
	}
	return
}

// GetTdEngineTableDataByTable 获取指定表数据信息
func (a *cTdEngine) GetTdEngineTableDataByTable(ctx context.Context, req *tdengine.GetTdEngineTableDataByTableReq) (res *tdengine.GetTdEngineTableDataByTableRes, err error) {
	data, err := service.TdEngine().GetTdEngineTableDataByTable(ctx, req.DbName, req.TableName)
	if data != nil {
		res = &tdengine.GetTdEngineTableDataByTableRes{
			TableDataInfo: data,
		}
	}
	return
}
