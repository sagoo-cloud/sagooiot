package model

import "github.com/gogf/gf/v2/os/gtime"

type SysAuthorizeInput struct {
	RoleId     int         `json:"roleId"     description:"角色ID"`
	ItemsType  string      `json:"itemsType"  description:"项目类型 menu菜单 button按钮 column列表字段 api接口"`
	ItemsId    int         `json:"itemsId"    description:"项目ID"`
	IsCheckAll int         `json:"isCheckAll" description:"是否全选 1是 0否"`
	IsDeleted  int         `json:"isDeleted"  description:"是否删除 0未删除 1已删除"`
	CreatedBy  uint        `json:"createdBy"  description:"创建人"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	DeletedBy  int         `json:"deletedBy"  description:"删除人"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除时间"`
}
