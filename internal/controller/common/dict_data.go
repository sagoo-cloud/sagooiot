package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

type cDictData struct{}

var DictData = cDictData{}

// GetDictData 获取字典数据
func (c *cDictData) GetDictData(ctx context.Context, req *common.GetDictReq) (res *common.GetDictRes, err error) {
	var input *model.GetDictInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	out, err := service.DictData().GetDictWithDataByType(ctx, input)
	if err != nil {
		return
	}
	var data *model.GetDictRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.GetDictRes{
		Data:   data.Data,
		Values: data.Values,
	}
	return
}

// List 获取字典数据列表
func (c *cDictData) List(ctx context.Context, req *common.DictDataSearchReq) (res *common.DictDataSearchRes, err error) {
	var input *model.SysDictSearchInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, out, err := service.DictData().List(ctx, input)
	if err != nil {
		return
	}
	res = new(common.DictDataSearchRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.List); err != nil {
			return
		}
	}
	return
}

// Add 添加字典数据
func (c *cDictData) Add(ctx context.Context, req *common.DictDataAddReq) (res *common.DictDataAddRes, err error) {
	var input *model.AddDictDataInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.DictData().Add(ctx, input, service.Context().GetUserId(ctx))
	return
}

// Get 获取对应的字典数据
func (c *cDictData) Get(ctx context.Context, req *common.DictDataGetReq) (res *common.DictDataGetRes, err error) {
	out, err := service.DictData().Get(ctx, req.DictCode)
	var data *model.SysDictDataRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.DictDataGetRes{
		Dict: data,
	}
	return
}

// Edit 修改字典数据
func (c *cDictData) Edit(ctx context.Context, req *common.DictDataEditReq) (res *common.DictDataEditRes, err error) {
	var input *model.EditDictDataInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.DictData().Edit(ctx, input, service.Context().GetUserId(ctx))
	return
}

func (c *cDictData) Delete(ctx context.Context, req *common.DictDataDeleteReq) (res *common.DictDataDeleteRes, err error) {
	err = service.DictData().Delete(ctx, req.Ids)
	return
}
