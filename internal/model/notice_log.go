package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/model/entity"
)

type NoticeLogAddInput struct {
	TemplateId  string      `json:"templateId" dc:"通知模板ID"`
	SendGateway string      `json:"sendGateway" dc:"通知发送通道：sms、work_weixin、dingding"`
	Addressee   string      `json:"addressee" dc:"收信人"`
	Title       string      `json:"title" dc:"通知标题"`
	Content     string      `json:"content" dc:"通知内容"`
	Status      int         `json:"status" dc:"发送状态：0=失败，1=成功"`
	FailMsg     string      `json:"failMsg" dc:"失败信息"`
	SendTime    *gtime.Time `json:"sendTime" dc:"发送时间"`
}

type NoticeLogSearchInput struct {
	Status int `json:"status" dc:"发送状态：0=失败，1=成功"`
	PaginationInput
}
type NoticeLogSearchOutput struct {
	List []NoticeLogList `json:"list" dc:"通知日志列表"`
	PaginationOutput
}

type NoticeLogList struct {
	entity.NoticeLog

	Gateway string `json:"gateway" dc:"发送通道"`
}
