package system

import (
	"context"
	systemV1 "sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

var SysOperLog = cSysOperLog{}

type cSysOperLog struct{}

// GetList 获取操作日志列表
func (a *cSysOperLog) GetList(ctx context.Context, req *systemV1.SysOperLogDoReq) (res *systemV1.SysOperLogDoRes, err error) {
	var input *model.SysOperLogDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	total, out, err := service.SysOperLog().GetList(ctx, input)
	if err != nil {
		return
	}
	res = new(systemV1.SysOperLogDoRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}
	return
}

/**
// AddSysOperLog 添加操作日志
func (a *cSysOperLog) AddSysOperLog(ctx context.Context, req *systemV1.AddSysOperLogReq) (res *systemV1.AddSysOperLogRes, err error) {
	err = service.SysOperLog().Add(ctx, req.SysOperLog)
	return
}
*/

// DetailSysOperLog 获取操作日志详情
func (a *cSysOperLog) DetailSysOperLog(ctx context.Context, req *systemV1.DetailSysOperLogReq) (res *systemV1.DetailSysOperLogRes, err error) {
	data, err := service.SysOperLog().Detail(ctx, req.OperId)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *entity.SysOperLog
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &systemV1.DetailSysOperLogRes{
			Data: detailRes,
		}
	}
	return
}

// DelSysOperLog 根据ID删除访问日志
func (a *cSysOperLog) DelSysOperLog(ctx context.Context, req *systemV1.DelSysOperLogReq) (res *systemV1.DelSysOperLogRes, err error) {
	err = service.SysOperLog().Del(ctx, req.OperIds)
	return
}
