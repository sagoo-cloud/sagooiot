package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var CityData = cCityData{}

type cCityData struct{}

// CityTree 获取列表
func (a *cCityData) CityTree(ctx context.Context, req *common.CityTreeReq) (res *common.CityTreeRes, err error) {
	info, err := service.CityData().GetList(ctx, req.Status, req.Name, req.Code)
	if err != nil {
		return
	}
	if info != nil {
		var dataTree []*model.CityTreeRes
		if err = gconv.Scan(info, &dataTree); err != nil {
			return
		}
		treeData, er := GetCityTreeRes(dataTree)
		if er != nil {
			return
		}

		res = &common.CityTreeRes{
			Data: treeData,
		}
	}

	return
}

func GetCityTreeRes(heatStationInfo []*model.CityTreeRes) (dataTree []*model.CityTreeRes, err error) {
	var parentNodeRes []*model.CityTreeRes
	if heatStationInfo != nil {
		//获取所有的根节点
		for _, v := range heatStationInfo {
			var parentNode *model.CityTreeRes
			if v.ParentId == -1 {
				if err = gconv.Scan(v, &parentNode); err != nil {
					return
				}
				parentNodeRes = append(parentNodeRes, parentNode)
			}
		}
	}
	treeData := GetCityChildrenTree(parentNodeRes, heatStationInfo)
	return treeData, nil
}

func GetCityChildrenTree(parentNodeRes []*model.CityTreeRes, data []*model.CityTreeRes) (dataTree []*model.CityTreeRes) {
	//循环所有一级节点
	for k, v := range parentNodeRes {
		//查询所有该节点下的所有子节点
		for _, j := range data {
			var node *model.CityTreeRes
			if j.ParentId == v.Id {
				if err := gconv.Scan(j, &node); err != nil {
					return
				}
				parentNodeRes[k].Children = append(parentNodeRes[k].Children, node)
			}
		}
		GetCityChildrenTree(v.Children, data)
	}
	return parentNodeRes
}

// AddCity 添加城市
func (a *cCityData) AddCity(ctx context.Context, req *common.AddCityReq) (res *common.AddCityRes, err error) {
	var city *entity.CityData
	if err = gconv.Scan(req.AddCityReq, &city); err != nil {
		return
	}
	err = service.CityData().Add(ctx, city)
	if err != nil {
		return
	}
	return
}

// EditCity 编辑城市
func (a *cCityData) EditCity(ctx context.Context, req *common.EditCityReq) (res *common.EditCityRes, err error) {
	var city *entity.CityData
	if err = gconv.Scan(req.EditCityReq, &city); err != nil {
		return
	}
	err = service.CityData().Edit(ctx, city)
	if err != nil {
		return
	}
	return
}

// GetCityById 根据ID获取城市
func (a *cCityData) GetCityById(ctx context.Context, req *common.GetCityByIdReq) (res *common.GetCityByIdRes, err error) {
	data, err := service.CityData().GetInfoById(ctx, req.Id)
	if err != nil {
		return
	}
	var city *model.CityRes
	if data != nil {
		if err = gconv.Scan(data, &city); err != nil {
			return
		}
	}
	res = &common.GetCityByIdRes{
		Data: city,
	}
	return
}

// DelCityById 根据ID删除城市
func (a *cCityData) DelCityById(ctx context.Context, req *common.DelCityByIdReq) (res *common.DelCityByIdRes, err error) {
	err = service.CityData().DelById(ctx, req.Id)
	if err != nil {
		return
	}
	return
}
