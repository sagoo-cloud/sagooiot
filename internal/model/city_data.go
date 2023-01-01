package model

import "github.com/gogf/gf/v2/os/gtime"

type CityTreeRes struct {
	Id        int            `json:"id"        description:""`
	Name      string         `json:"name"      description:"名字"`
	Code      string         `json:"code"      description:"编码"`
	ParentId  int            `json:"parentId"   description:"父ID"`
	Sort      int            `json:"sort"      description:"排序"`
	Status    uint           `json:"status"    description:"状态;0:禁用;1:正常"`
	IsDeleted int            `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreateBy  uint           `json:"createBy"  description:"创建者"`
	CreatedAt *gtime.Time    `json:"createdAt" description:"创建日期"`
	UpdateBy  uint           `json:"updateBy"  description:"更新者"`
	UpdatedAt *gtime.Time    `json:"updatedAt" description:"修改日期"`
	DeletedBy int            `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time    `json:"deletedAt" description:"删除时间"`
	Children  []*CityTreeRes `json:"children" description:"子集"`
}

type AddCityReq struct {
	Name     string `json:"name"      description:"名字"`
	Code     string `json:"code"      description:"编码"`
	ParentId int    `json:"parentId"   description:"父ID"`
	Sort     int    `json:"sort"      description:"排序"`
	Status   uint   `json:"status"    description:"状态;0:禁用;1:正常"`
}

type EditCityReq struct {
	Id       int    `json:"id"        description:""`
	Name     string `json:"name"      description:"名字"`
	Code     string `json:"code"      description:"编码"`
	ParentId int    `json:"parentId"   description:"父ID"`
	Sort     int    `json:"sort"      description:"排序"`
	Status   uint   `json:"status"    description:"状态;0:禁用;1:正常"`
}

type CityRes struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:"名字"`
	Code      string      `json:"code"      description:"编码"`
	ParentId  int         `json:"parentId"   description:"父ID"`
	Sort      int         `json:"sort"      description:"排序"`
	Status    uint        `json:"status"    description:"状态;0:禁用;1:正常"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreateBy  uint        `json:"createBy"  description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建日期"`
	UpdateBy  uint        `json:"updateBy"  description:"更新者"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改日期"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
