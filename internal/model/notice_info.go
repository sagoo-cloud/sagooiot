package model

type GetNoticeInfoListInput struct {
	ConfigId string `json:"configId"   description:""`
	ComeFrom string `json:"comeFrom"   description:""`
	Method   string `json:"method"     description:""`
	Status   int    `json:"status"     description:""`
	PaginationInput
}
type NoticeInfoListOutput struct {
	Data []NoticeInfoOutput
	PaginationOutput
}
type NoticeInfoOutput struct {
	Status     string `json:"status"          description:""`
	MethodCron string `json:"methodCron"          description:""`
	Id         string `json:"id"          description:""`
	Totag      string `json:"totag"          description:""`
	Method     string `json:"method"          description:""`
	MsgBody    string `json:"msgBody"          description:""`
	MsgUrl     string `json:"msgUrl"          description:""`
	ConfigId   string `json:"configId"          description:""`
	ComeFrom   string `json:"comeFrom"          description:""`
	UserIds    string `json:"userIds"          description:""`
	MethodNum  string `json:"methodNum"          description:""`
	CreatedAt  string `json:"createdAt"          description:""`
	MsgTitle   string `json:"msgTitle"          description:""`
	OrgIds     string `json:"orgIds"          description:""`
}
type NoticeInfoAddInput struct {
	Totag      string `json:"totag"          description:""`
	Status     string `json:"status"          description:""`
	MethodCron string `json:"methodCron"          description:""`
	ConfigId   string `json:"configId"          description:""`
	ComeFrom   string `json:"comeFrom"          description:""`
	Method     string `json:"method"          description:""`
	MsgBody    string `json:"msgBody"          description:""`
	MsgUrl     string `json:"msgUrl"          description:""`
	UserIds    string `json:"userIds"          description:""`
	MsgTitle   string `json:"msgTitle"          description:""`
	OrgIds     string `json:"orgIds"          description:""`
	MethodNum  string `json:"methodNum"          description:""`
	CreatedAt  string `json:"createdAt"          description:""`
}
type NoticeInfoEditInput struct {
	Id int `json:"id"          description:"ID"`
	NoticeInfoAddInput
}
