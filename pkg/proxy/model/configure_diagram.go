package model

import "github.com/gogf/gf/v2/os/gtime"

type ConfigureDiagramOut struct {
	Id        int                 `json:"id"        description:""`
	Code      string              `json:"code"      description:""`
	DeptId    int                 `json:"deptId"    description:"部门ID"`
	FolderId  int                 `json:"folderId"  description:"文件夹ID"`
	Name      string              `json:"name"      description:"名字"`
	Types     int                 `json:"types"     description:"类型"`
	Images    string              `json:"images"    description:"图片"`
	PointIds  []map[string]string `json:"pointIds"  description:""`
	JsonData  string              `json:"jsonData"  description:""`
	Remark    string              `json:"remark"    description:""`
	Status    int                 `json:"status"    description:"状态 0 停用 1启用"`
	IsDeleted int                 `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint                `json:"createdBy" description:"创建人"`
	CreatedAt *gtime.Time         `json:"createdAt" description:"创建时间"`
	UpdatedBy int                 `json:"updatedBy" description:"修改人"`
	UpdatedAt *gtime.Time         `json:"updatedAt" description:"更新时间"`
	DeletedBy int                 `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time         `json:"deletedAt" description:"删除时间"`
}
