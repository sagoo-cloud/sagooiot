package system

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PostDoReq struct {
	g.Meta   `path:"/post/tree" method:"get" summary:"获取岗位列表" tags:"岗位管理"`
	Status   int    `p:"status" description:"状态:-1为全部,0为正常,1为停用" `
	PostName string `p:"postName" description:"岗位名称"`
	PostCode string `p:"postCode" description:"岗位编码"`
}
type PostDoRes struct {
	Data []*model.PostRes
}

type AddPostReq struct {
	g.Meta   `path:"/post/add" method:"post" summary:"添加岗位" tags:"岗位管理"`
	ParentId int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	PostName string `json:"postName"  description:"岗位名称" v:"required#请输入岗位名称"`
	PostSort int    `json:"postSort"  description:"显示顺序"`
	Status   uint   `json:"status"    description:"状态（0正常 1停用）" v:"required#请选择状态"`
	Remark   string `json:"remark"    description:"备注"`
}
type AddPostRes struct {
}

type EditPostReq struct {
	g.Meta   `path:"/post/edit" method:"put" summary:"编辑岗位" tags:"岗位管理"`
	PostId   int64  `json:"postId"    description:"岗位ID" v:"required#岗位ID不能为空"`
	ParentId int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	PostName string `json:"postName"  description:"岗位名称" v:"required#请输入岗位名称"`
	PostSort int    `json:"postSort"  description:"显示顺序"`
	Status   uint   `json:"status"    description:"状态（0正常 1停用）" v:"required#请选择状态"`
	Remark   string `json:"remark"    description:"备注"`
}
type EditPostRes struct {
}

type DetailPostReq struct {
	g.Meta `path:"/post/detail" method:"get" summary:"根据ID获取岗位详情" tags:"岗位管理"`
	PostId int64 `p:"post_id" description:"岗位ID"  v:"required#ID不能为空"`
}
type DetailPostRes struct {
	Data *model.DetailPostRes
}

type DelPostReq struct {
	g.Meta `path:"/post/del" method:"delete" summary:"根据ID删除岗位" tags:"岗位管理"`
	PostId int64 `p:"post_id" description:"岗位ID"  v:"required#ID不能为空"`
}
type DelPostRes struct {
}
