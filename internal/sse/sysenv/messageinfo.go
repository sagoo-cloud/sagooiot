package sysenv

import (
	"context"
	"encoding/json"
	"sagooiot/internal/service"
)

// GetUnReadMessageLast 获取用户最后一条未读消息
func GetUnReadMessageLast(userId int) (data []byte) {
	tmpData, err := service.SysMessage().GetUnReadMessageLast(context.TODO(), userId)
	if err != nil {
		return
	}
	if tmpData != nil && len(tmpData) > 0 {
		data, _ = json.Marshal(tmpData)
	}
	return
}
