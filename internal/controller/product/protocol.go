package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

var Protocol = cProtocol{}

type cProtocol struct{}

// ListMessageProtocol 消息协议列表
func (c *cProtocol) ListMessageProtocol(ctx context.Context, req *product.ListMessageProtocolReq) (res *product.ListMessageProtocolRes, err error) {
	res = &product.ListMessageProtocolRes{
		Data: []*model.MessageProtocolRes{
			{Key: "modbus", Name: "modbus"},
		},
	}
	return
}

// ListTrunsportProtocol 传输协议列表
func (c *cProtocol) ListTrunsportProtocol(ctx context.Context, req *product.ListTrunsportProtocolReq) (res *product.ListTrunsportProtocolRes, err error) {
	res = &product.ListTrunsportProtocolRes{
		Data: []*model.TrunsportProtocolRes{
			{Key: "tcp-server", Name: "tcp服务端"},
			{Key: "tcp-client", Name: "tcp客户端"},
			{Key: "udp-server", Name: "udp服务端"},
			{Key: "udp-client", Name: "udp客户端"},
			{Key: "serial", Name: "serial"},
		},
	}
	return
}
