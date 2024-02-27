package model

import "sagooiot/internal/model/entity"

type AlarmLevelOutput struct {
	*entity.AlarmLevel
}

type AlarmLevelListOutput struct {
	List []*entity.AlarmLevel `json:"list"`
}

type AlarmLevelEditInput struct {
	Level uint   `json:"level" dc:"告警级别" v:"required|in:1,2,3,4,5#告警级别不能为空|告警级别不正确"`
	Name  string `json:"name" dc:"告警名称" v:"required#请输入告警名称"`
}
