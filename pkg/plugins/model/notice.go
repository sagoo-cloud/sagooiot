package model

type NoticeSendObject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NoticeInfoData struct {
	ConfigId     string             `orm:"config_id"   json:"config_id"`             //
	ComeFrom     string             `orm:"come_from"   json:"come_from"`             //
	Method       string             `orm:"method"      json:"method"`                //
	MethodCron   string             `orm:"method_cron" json:"method_cron"`           //
	MethodNum    int                `orm:"method_num"  json:"method_num"`            //
	MsgTitle     string             `orm:"msg_title"   json:"msg_title"`             //
	MsgBody      string             `orm:"msg_body"    json:"msg_body"`              //
	MsgUrl       string             `orm:"msg_url"     json:"msg_url"`               //
	UserIds      string             `orm:"user_ids"    json:"user_ids"`              //
	PartyIds     string             `orm:"party_ids"   json:"party_ids"`             //
	Totag        []NoticeSendObject `orm:"totag"       json:"totag"`                 //
	TemplateCode string             `orm:"template_code"       json:"template_code"` //
}

type NoticeData struct {
	Config    map[interface{}]interface{}
	SendParam map[string]interface{}
	Msg       NoticeInfoData
}
