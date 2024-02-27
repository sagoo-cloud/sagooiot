package model

type GetNoticeConfigListInput struct {
	SendGateway string `json:"sendGateway"          description:"sendGateway"`
	Types       string `json:"types"          description:"types"`
	PaginationInput
}
type NoticeConfigListOutput struct {
	Data []NoticeConfigOutput
	PaginationOutput
}
type NoticeConfigOutput struct {
	Id          string `json:"id"          description:""`
	DeptId      int    `json:"deptId"      description:"部门ID"`
	Title       string `json:"title"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Types       string `json:"types"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
}
type NoticeConfigAddInput struct {
	Id          string `json:"id"          description:"ID"`
	Title       string `json:"title"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Types       string `json:"types"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
}
type NoticeConfigEditInput struct {
	NoticeConfigAddInput
}
