package model

import "github.com/sagoo-cloud/sagooiot/internal/model/entity"

const (
	AlarmLevel_1 uint = iota + 1 // 告警级别：超紧急
	AlarmLevel_2                 // 告警级别：紧急
	AlarmLevel_3                 // 告警级别：严重
	AlarmLevel_4                 // 告警级别：一般
	AlarmLevel_5                 // 告警级别：提醒
)

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
