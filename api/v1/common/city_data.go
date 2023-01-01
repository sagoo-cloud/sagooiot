package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

type CityTreeReq struct {
	g.Meta `path:"/city/tree" method:"get" summary:"获取城市列表" tags:"城市管理"`
	Status int    `json:"status" description:"状态:--1为全部,0为禁用,1为正常" `
	Name   string `json:"name" description:"城市名" `
	Code   string `json:"code" description:"城市编码" `
}
type CityTreeRes struct {
	Data []*model.CityTreeRes
}

type AddCityReq struct {
	g.Meta `path:"/city/add" method:"post" summary:"添加城市" tags:"城市管理"`
	*model.AddCityReq
}
type AddCityRes struct {
}

type EditCityReq struct {
	g.Meta `path:"/city/edit" method:"put" summary:"编辑城市" tags:"城市管理"`
	*model.EditCityReq
}
type EditCityRes struct {
}

type GetCityByIdReq struct {
	g.Meta `path:"/city/getInfoById" method:"get" summary:"根据ID获取城市信息" tags:"城市管理"`
	Id     int `json:"id"        description:"" v:"required#ID不能为空"`
}
type GetCityByIdRes struct {
	Data *model.CityRes
}

type DelCityByIdReq struct {
	g.Meta `path:"/city/del" method:"delete" summary:"根据ID删除城市信息" tags:"城市管理"`
	Id     int `json:"id"        description:"" v:"required#ID不能为空"`
}
type DelCityByIdRes struct {
}
