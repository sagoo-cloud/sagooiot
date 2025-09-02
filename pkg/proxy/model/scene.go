package model

type SceneOutput struct {
	Id          string `json:"id"          description:"场景ID"`
	Code        string `json:"code"          description:"场景编码"`
	DeptId      int    `json:"deptId"      description:"部门ID"`
	Name        string `json:"name"          description:"场景名称"`
	Description string `json:"description"          description:"场景描述"`
	SceneType   string `json:"sceneType"          description:"场景类型"`
	Status      string `json:"status"          description:"状态：0=未启用，1=已启用"`
	StartTime   string `json:"startTime"          description:"开始时间"`
	CreateBy    string `json:"createBy"          description:"创建者"`
	CreatedAt   string `json:"createdAt"          description:"创建时间"`
	UpdateBy    string `json:"updateBy"          description:"更新者"`
	UpdatedAt   string `json:"updatedAt"          description:"更新时间"`
	DeletedBy   string `json:"deletedBy"          description:"删除者"`
	DeletedAt   string `json:"deletedAt"          description:"删除时间"`
}
