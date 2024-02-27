package model

type GetNoticeTemplateListInput struct {
	ConfigId    string `json:"configId"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	PaginationInput
}
type NoticeTemplateListOutput struct {
	Data []NoticeTemplateOutput
	PaginationOutput
}
type NoticeTemplateOutput struct {
	Id          string `json:"id"          description:""`
	DeptId      int    `json:"deptId"      description:"部门ID"`
	ConfigId    string `json:"configId"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
}
type NoticeTemplateAddInput struct {
	Id          string `json:"id"          description:"ID"`
	DeptId      int    `json:"deptId"      description:"部门ID"`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	ConfigId    string `json:"configId"          description:""`
}
type NoticeTemplateEditInput struct {
	NoticeTemplateAddInput
}
