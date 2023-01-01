package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	systemV1 "github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility"
	"github.com/sagoo-cloud/sagooiot/utility/response"

	"github.com/gogf/gf/v2/util/gconv"
)

var SysLoginLog = cSysLoginLog{}

type cSysLoginLog struct{}

// GetList 获取访问日志列表
func (a *cSysLoginLog) GetList(ctx context.Context, req *systemV1.SysLoginLogDoReq) (res *systemV1.SysLoginLogDoRes, err error) {
	var reqData = new(model.SysLoginLogInput)
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return
	}
	total, page, outData, err := service.SysLoginLog().GetList(ctx, reqData)
	if err != nil {
		return
	}
	res = new(systemV1.SysLoginLogDoRes)
	res.Total = total
	res.CurrentPage = page
	if outData != nil {
		if err = gconv.Scan(outData, &res.Data); err != nil {
			return
		}
	}
	return
}

// Export 导出登录访问日志
func (a *cSysLoginLog) Export(ctx context.Context, req *systemV1.SysLoginLogDoExportReq) (res *systemV1.SysLoginLogDoExportRes, err error) {
	var reqData = new(model.SysLoginLogInput)
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return
	}
	_, _, outList, err := service.SysLoginLog().GetList(ctx, reqData)
	if err != nil {
		return
	}

	//处理数据并导出
	var resData []interface{}
	for _, d := range outList {
		resData = append(resData, d)
	}
	data := utility.ToExcel(resData)
	var request = g.RequestFromCtx(ctx)
	response.ToXls(request, data, "SysLoginLog")

	return
}

// DetailSysLoginLog 获取访问日志详情
func (a *cSysLoginLog) DetailSysLoginLog(ctx context.Context, req *systemV1.DetailSysLoginLogReq) (res *systemV1.DetailSysLoginLogRes, err error) {
	data, err := service.SysLoginLog().Detail(ctx, req.InfoId)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *entity.SysLoginLog
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &systemV1.DetailSysLoginLogRes{
			Data: detailRes,
		}
	}
	return
}

// DelSysLoginLog 根据ID删除访问日志
func (a *cSysLoginLog) DelSysLoginLog(ctx context.Context, req *systemV1.DelSysLoginLogReq) (res *systemV1.DelSysLoginLogRes, err error) {
	err = service.SysLoginLog().Del(ctx, req.InfoIds)
	return
}
