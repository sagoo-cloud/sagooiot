package core

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type messageRouterHandler interface {
	DeviceDataHandle()
}

type messageRouter struct {
	ctx          context.Context
	msg          []byte
	data         map[string]interface{}
	msgType      string
	deviceDetail *model.DeviceOutput
}

func (m messageRouter) router() {
	var h messageRouterHandler
	switch m.msgType {
	case consts.PropertyReport:
		h = MessagePropertyReporter{m}
	}
	if h == nil {
		g.Log().Errorf(m.ctx, "message type:%s not supported, message:%s, message ignored", m.msgType, string(m.msg))
		return
	}
	h.DeviceDataHandle()
}
