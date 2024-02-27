package model

import "github.com/gogf/gf/v2/os/gtime"

type RoleTreeRes struct {
	Id        uint           `json:"id"        description:""`
	ParentId  int            `json:"parentId"  description:"父ID"`
	ListOrder uint           `json:"listOrder" description:"排序"`
	Name      string         `json:"name"      description:"角色名称"`
	DataScope uint           `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	Remark    string         `json:"remark"    description:"备注"`
	Status    uint           `json:"status"    description:"状态;0:禁用;1:正常"`
	CreatedBy uint           `json:"createdBy"  description:"创建者"`
	CreatedAt *gtime.Time    `json:"createdAt" description:"创建日期"`
	UpdatedBy uint           `json:"updatedBy"  description:"更新者"`
	UpdatedAt *gtime.Time    `json:"updatedAt" description:"修改日期"`
	Children  []*RoleTreeRes `json:"children" description:"子集"`
}

type RoleTreeOut struct {
	Id        uint           `json:"id"        description:""`
	ParentId  int            `json:"parentId"  description:"父ID"`
	ListOrder uint           `json:"listOrder" description:"排序"`
	Name      string         `json:"name"      description:"角色名称"`
	DataScope uint           `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	Remark    string         `json:"remark"    description:"备注"`
	Status    uint           `json:"status"    description:"状态;0:禁用;1:正常"`
	CreatedBy uint           `json:"createdBy"  description:"创建者"`
	CreatedAt *gtime.Time    `json:"createdAt" description:"创建日期"`
	UpdatedBy uint           `json:"updatedBy"  description:"更新者"`
	UpdatedAt *gtime.Time    `json:"updatedAt" description:"修改日期"`
	Children  []*RoleTreeOut `json:"children" description:"子集"`
}

type AddRoleInput struct {
	ParentId  int    `json:"parentId"  description:"父ID"`
	Name      string `json:"name"      description:"角色名称"`
	ListOrder uint   `json:"listOrder" description:"排序"`
	Status    uint   `json:"status"    description:"状态;0:禁用;1:正常"`
	Remark    string `json:"remark"    description:"备注"`
}

type EditRoleInput struct {
	Id        uint   `json:"id"        description:"ID"`
	ParentId  int    `json:"parentId"  description:"父ID"`
	Name      string `json:"name"      description:"角色名称"`
	ListOrder uint   `json:"listOrder" description:"排序"`
	Status    uint   `json:"status"    description:"状态;0:禁用;1:正常"`
	Remark    string `json:"remark"    description:"备注"`
}

type RoleInfoRes struct {
	Id        uint        `json:"id"        description:""`
	ParentId  int         `json:"parentId"  description:"父ID"`
	ListOrder uint        `json:"listOrder" description:"排序"`
	Name      string      `json:"name"      description:"角色名称"`
	DataScope uint        `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	DeptIds   []int64     `json:"deptIds" description:"数据范围为自定义数据权限时返回部门ID数组"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    uint        `json:"status"    description:"状态;0:禁用;1:正常"`
	CreatedBy uint        `json:"createdBy"  description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建日期"`
	UpdatedBy uint        `json:"updatedBy"  description:"更新者"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改日期"`
}

type RoleInfoOut struct {
	Id        uint        `json:"id"        description:""`
	ParentId  int         `json:"parentId"  description:"父ID"`
	ListOrder uint        `json:"listOrder" description:"排序"`
	Name      string      `json:"name"      description:"角色名称"`
	DataScope uint        `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	DeptIds   []int64     `json:"deptIds" description:"数据范围为自定义数据权限时返回部门ID数组"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    uint        `json:"status"    description:"状态;0:禁用;1:正常"`
	CreatedBy uint        `json:"createdBy"  description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建日期"`
	UpdatedBy uint        `json:"updatedBy"  description:"更新者"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改日期"`
}
