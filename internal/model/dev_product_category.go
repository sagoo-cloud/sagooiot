package model

import "sagooiot/internal/model/entity"

type ProductCategoryTreeOutput struct {
	*entity.DevProductCategory
	Children []*ProductCategoryTreeOutput `json:"children" description:"子分类"`
}

type ProductCategoryOutput struct {
	*entity.DevProductCategory
}

type AddProductCategoryInput struct {
	ParentId uint   `json:"parentId" description:"父级分类ID"`
	Key      string `json:"key" description:"分类标识" v:"required#请输入标识"`
	Name     string `json:"name" description:"分类名称" v:"required#请输入名称"`
	Desc     string `json:"desc" description:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Sort     int    `json:"sort" dc:"排序"`
}

type EditProductCategoryInput struct {
	Id       uint   `json:"id" description:"分类ID" v:"required#分类ID不能为空"`
	ParentId uint   `json:"parentId" description:"父级分类ID"`
	Key      string `json:"key" description:"分类标识" v:"required#请输入标识"`
	Name     string `json:"name" description:"分类名称" v:"required#请输入名称"`
	Desc     string `json:"desc" description:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Sort     int    `json:"sort" dc:"排序"`
}
