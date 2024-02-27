package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

type cConfigData struct{}

var ConfigData = cConfigData{}

// List 系统参数列表
func (c *cConfigData) List(ctx context.Context, req *common.ConfigSearchReq) (res *common.ConfigSearchRes, err error) {
	var input *model.ConfigDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, out, err := service.ConfigData().List(ctx, input)
	if err != nil {
		return
	}
	res = new(common.ConfigSearchRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.List); err != nil {
			return
		}
	}
	return
}

// Add 添加系统参数
func (c *cConfigData) Add(ctx context.Context, req *common.ConfigAddReq) (res *common.ConfigAddRes, err error) {
	var input *model.AddConfigInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.ConfigData().Add(ctx, input, service.Context().GetUserId(ctx))
	return
}

// Get 获取系统参数
func (c *cConfigData) Get(ctx context.Context, req *common.ConfigGetReq) (res *common.ConfigGetRes, err error) {
	out, err := service.ConfigData().Get(ctx, req.Id)
	if err != nil {
		return
	}
	var data *model.SysConfigRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.ConfigGetRes{
		Data: data,
	}
	return
}

// Edit 修改系统参数
func (c *cConfigData) Edit(ctx context.Context, req *common.ConfigEditReq) (res *common.ConfigEditRes, err error) {
	var input *model.EditConfigInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.ConfigData().Edit(ctx, input, service.Context().GetUserId(ctx))
	return
}

// Delete 删除系统参数
func (c *cConfigData) Delete(ctx context.Context, req *common.ConfigDeleteReq) (res *common.ConfigDeleteRes, err error) {
	err = service.ConfigData().Delete(ctx, req.Ids)
	return
}

// ConfigGetByKey 根据key获取系统参数
func (c *cConfigData) ConfigGetByKey(ctx context.Context, req *common.ConfigGetByKeyReq) (res *common.ConfigGetByKeyRes, err error) {
	out, err := service.ConfigData().GetByKey(ctx, req.ConfigKey)
	if err != nil {
		return
	}
	var data *model.SysConfigRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.ConfigGetByKeyRes{
		Data: data,
	}
	return
}

// ConfigGetByKeys 根据key数组获取系统参数
func (c *cConfigData) ConfigGetByKeys(ctx context.Context, req *common.ConfigGetByKeysReq) (res *common.ConfigGetByKeysRes, err error) {
	out, err := service.ConfigData().GetConfigByKeys(ctx, req.ConfigKey)
	if err != nil {
		return
	}
	var data []*model.SysConfigRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.ConfigGetByKeysRes{
		Data: data,
	}
	return
}

// GetSysConfigSetting 获取系统配置
func (c *cConfigData) GetSysConfigSetting(ctx context.Context, req *common.GetSysConfigSettingReq) (res *common.GetSysConfigSettingRes, err error) {
	out, err := service.ConfigData().GetSysConfigSetting(ctx, req.Types)
	if err != nil {
		return
	}
	var data []*model.SysConfigRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &common.GetSysConfigSettingRes{
		Info: data,
	}
	return
}

// EditSysConfigSetting 修改配置
func (c *cConfigData) EditSysConfigSetting(ctx context.Context, req *common.EditSysConfigSettingReq) (res *common.EditSysConfigSettingRes, err error) {
	var input []*model.EditConfigInput
	if err = gconv.Scan(req.ConfigInfo, &input); err != nil {
		return
	}
	err = service.ConfigData().EditSysConfigSetting(ctx, input)
	if err != nil {
		return
	}
	return
}
