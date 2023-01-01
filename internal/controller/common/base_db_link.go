package common

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

var BaseDbLink = cBaseDbLink{}

type cBaseDbLink struct{}

// GetList 获取数据源列表
func (a *cBaseDbLink) GetList(ctx context.Context, req *common.BaseDbLinkDoReq) (res *common.BaseDbLinkDoRes, err error) {
	var input *model.BaseDbLinkDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	total, out, err := service.BaseDbLink().GetList(ctx, input)
	if err != nil {
		return
	}
	res = new(common.BaseDbLinkDoRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}
	return
}

// AddBaseDbLink 添加数据源
func (a *cBaseDbLink) AddBaseDbLink(ctx context.Context, req *common.AddBaseDbLinkReq) (res *common.AddBaseDbLinkRes, err error) {
	var input *model.AddBaseDbLinkInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.BaseDbLink().Add(ctx, input)
	return
}

// DetailBaseDbLink 获取数据源详情
func (a *cBaseDbLink) DetailBaseDbLink(ctx context.Context, req *common.DetailBaseDbLinkReq) (res *common.DetailBaseDbLinkRes, err error) {
	data, err := service.BaseDbLink().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailBaseDbLinkRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &common.DetailBaseDbLinkRes{
			Data: detailRes,
		}
	}
	return
}

// EditBaseDbLink 编辑数据源
func (a *cBaseDbLink) EditBaseDbLink(ctx context.Context, req *common.EditBaseDbLinkReq) (res *common.EditBaseDbLinkRes, err error) {
	var input *model.EditBaseDbLinkInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.BaseDbLink().Edit(ctx, input)
	return
}

// DelBaseDbLink 根据ID删除数据源
func (a *cBaseDbLink) DelBaseDbLink(ctx context.Context, req *common.DelBaseDbLinkReq) (res *common.DelBaseDbLinkRes, err error) {
	err = service.BaseDbLink().Del(ctx, req.Id)
	return
}
