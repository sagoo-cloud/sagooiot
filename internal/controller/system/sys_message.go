package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysMessage = cSysMessage{}

type cSysMessage struct{}

// GetMessageList 获取消息列表
func (a *cSysMessage) GetMessageList(ctx context.Context, req *system.GetMessageListReq) (res *system.GetMessageListRes, err error) {
	var input *model.MessageListDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	total, out, err := service.SysMessage().GetList(ctx, input)
	if err != nil {
		return
	}
	res = new(system.GetMessageListRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Info); err != nil {
			return
		}
	}
	return
}

/*// AddMessage 添加消息
func (a *cSysMessage) AddMessage(ctx context.Context, req *system.AddMessageReq) (res *system.AddMessageRes, err error) {
	var message *model.AddMessageReq
	if err = gconv.Scan(req.AddMessageReq, &message); err != nil {
		return
	}
	err = service.SysMessage().Add(ctx, message)
	return
}*/

// GetUnReadMessageAll 获取所有未读消息列表
func (a *cSysMessage) GetUnReadMessageAll(ctx context.Context, req *system.GetUnReadMessageAllReq) (res *system.GetUnReadMessageAllRes, err error) {
	var input *model.MessageListDoInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	total, out, err := service.SysMessage().GetUnReadMessageAll(ctx, input)
	if err != nil {
		return
	}
	res = new(system.GetUnReadMessageAllRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Info); err != nil {
			return
		}
	}
	return
}

// GetUnReadMessageCount 获取所有未读消息数量
func (a *cSysMessage) GetUnReadMessageCount(ctx context.Context, req *system.GetUnReadMessageCountReq) (res *system.GetUnReadMessageCountRes, err error) {
	out, err := service.SysMessage().GetUnReadMessageCount(ctx)
	if err != nil {
		return
	}
	res = &system.GetUnReadMessageCountRes{
		Count: out,
	}
	return
}

// DelMessage 删除消息
func (a *cSysMessage) DelMessage(ctx context.Context, req *system.DelMessageReq) (res *system.DelMessageRes, err error) {
	err = service.SysMessage().DelMessage(ctx, req.Ids)
	return
}

// ClearMessage 一键清空消息
func (a *cSysMessage) ClearMessage(ctx context.Context, req *system.ClearMessageReq) (res *system.ClearMessageRes, err error) {
	err = service.SysMessage().ClearMessage(ctx)
	return
}

// ReadMessage 阅读消息
func (a *cSysMessage) ReadMessage(ctx context.Context, req *system.ReadMessageReq) (res *system.ReadMessageRes, err error) {
	err = service.SysMessage().ReadMessage(ctx, req.Id)
	return
}

// ReadMessageAll 全部阅读消息
func (a *cSysMessage) ReadMessageAll(ctx context.Context, req *system.ReadMessageAllReq) (res *system.ReadMessageAllRes, err error) {
	err = service.SysMessage().ReadMessageAll(ctx)
	return
}
