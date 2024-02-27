package model

import "github.com/gogf/gf/v2/os/gtime"

type PostRes struct {
	PostId    int64       `json:"postId"    description:"岗位ID"`
	ParentId  int64       `json:"parentId"  description:"父ID"`
	PostCode  string      `json:"postCode"  description:"岗位编码"`
	PostName  string      `json:"postName"  description:"岗位名称"`
	PostSort  int         `json:"postSort"  description:"显示顺序"`
	Status    uint        `json:"status"    description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"    description:"备注"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	Children  []*PostRes  `json:"children" description:"子集"`
}

type PostOut struct {
	PostId    int64       `json:"postId"    description:"岗位ID"`
	ParentId  int64       `json:"parentId"  description:"父ID"`
	PostCode  string      `json:"postCode"  description:"岗位编码"`
	PostName  string      `json:"postName"  description:"岗位名称"`
	PostSort  int         `json:"postSort"  description:"显示顺序"`
	Status    uint        `json:"status"    description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"    description:"备注"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	Children  []*PostOut  `json:"children" description:"子集"`
}

type AddPostInput struct {
	ParentId int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	PostName string `json:"postName"  description:"岗位名称" v:"required#请输入岗位名称"`
	PostSort int    `json:"postSort"  description:"显示顺序"`
	Status   uint   `json:"status"    description:"状态（0正常 1停用）" v:"required#请选择状态"`
	Remark   string `json:"remark"    description:"备注"`
}

type EditPostInput struct {
	PostId   int64  `json:"postId"    description:"岗位ID" v:"required#岗位ID不能为空"`
	ParentId int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	PostName string `json:"postName"  description:"岗位名称" v:"required#请输入岗位名称"`
	PostSort int    `json:"postSort"  description:"显示顺序"`
	Status   uint   `json:"status"    description:"状态（0正常 1停用）" v:"required#请选择状态"`
	Remark   string `json:"remark"    description:"备注"`
}

type DetailPostRes struct {
	PostId    int64       `json:"postId"    description:"岗位ID"`
	ParentId  int64       `json:"parentId"  description:"父ID"`
	PostCode  string      `json:"postCode"  description:"岗位编码"`
	PostName  string      `json:"postName"  description:"岗位名称"`
	PostSort  int         `json:"postSort"  description:"显示顺序"`
	Status    uint        `json:"status"    description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"    description:"备注"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
}

type DetailPostOut struct {
	PostId    int64       `json:"postId"    description:"岗位ID"`
	ParentId  int64       `json:"parentId"  description:"父ID"`
	PostCode  string      `json:"postCode"  description:"岗位编码"`
	PostName  string      `json:"postName"  description:"岗位名称"`
	PostSort  int         `json:"postSort"  description:"显示顺序"`
	Status    uint        `json:"status"    description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"    description:"备注"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
}
