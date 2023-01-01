package product

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 消息协议
type ListMessageProtocolReq struct {
	g.Meta `path:"/protocol/message_protocol_list" method:"get" summary:"消息协议" tags:"协议"`
}
type ListMessageProtocolRes struct {
	Data []*model.MessageProtocolRes `json:"data" dc:"消息协议"`
}

// 传输协议
type ListTrunsportProtocolReq struct {
	g.Meta `path:"/protocol/trunsport_protocol_list" method:"get" summary:"传输协议" tags:"协议"`
}
type ListTrunsportProtocolRes struct {
	Data []*model.TrunsportProtocolRes `json:"data" dc:"传输协议"`
}
