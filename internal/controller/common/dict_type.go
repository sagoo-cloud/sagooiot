package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

type cDictType struct{}

var DictType = cDictType{}

// List 字典类型列表
func (c *cDictType) List(ctx context.Context, req *common.DictTypeSearchReq) (res *common.DictTypeSearchRes, err error) {
	var input *model.DictTypeDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, out, err := service.DictType().List(ctx, input)
	if err != nil {
		return
	}
	res = new(common.DictTypeSearchRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.DictTypeList); err != nil {
			return
		}
	}

	return
}

// Add 添加字典类型
func (c *cDictType) Add(ctx context.Context, req *common.DictTypeAddReq) (res *common.DictTypeAddRes, err error) {
	var input *model.AddDictTypeInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.DictType().Add(ctx, input, service.Context().GetUserId(ctx))
	return
}

// Get 获取字典类型
func (c *cDictType) Get(ctx context.Context, req *common.DictTypeGetReq) (res *common.DictTypeGetRes, err error) {
	out, err := service.DictType().Get(ctx, req)
	if err != nil {
		return
	}
	var data *model.SysDictTypeRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.DictTypeGetRes{
		DictType: data,
	}
	return
}

// Edit 修改字典数据
func (c *cDictType) Edit(ctx context.Context, req *common.DictTypeEditReq) (res *common.DictTypeEditRes, err error) {
	var input *model.EditDictTypeInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.DictType().Edit(ctx, input, service.Context().GetUserId(ctx))
	return
}

func (c *cDictType) Delete(ctx context.Context, req *common.DictTypeDeleteReq) (res *common.DictTypeDeleteRes, err error) {
	err = service.DictType().Delete(ctx, req.DictIds)
	return
}
