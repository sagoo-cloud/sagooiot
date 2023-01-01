package model

// 消息协议
type MessageProtocolRes struct {
	Key  string `json:"key" dc:"协议标识"`
	Name string `json:"name" dc:"协议名称"`
}

// 传输协议
type TrunsportProtocolRes struct {
	Key  string `json:"key" dc:"协议标识"`
	Name string `json:"name" dc:"协议名称"`
}
