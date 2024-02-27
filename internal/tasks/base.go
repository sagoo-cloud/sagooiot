package tasks

type TaskJob struct {
	ID         string        //任务ID
	TaskType   string        //任务类型
	MethodName string        //方法名
	Params     []interface{} //参数
	Explain    string        //任务描述
}

func (t TaskJob) GetFuncNameList() (res map[string]string) {
	res = map[string]string{
		"ClearOperationLogByDays": "清理超过指定天数的操作日志",
		"ClearNoticeLogByDays":    "清理超过指定天数的通知服务日志",
		"ClearAlarmLogByDays":     "清理超过指定天数的告警日志",
		"ClearTDengineLogByDays":  "清理超过指定天数的TD日志",
		"GetAccessURL":            "访问URL",
		"DataSourceSync":          "数据源同步",
		"DataTemplate":            "数据模型聚合数据",
		"DeviceLogClear":          "设备日志清理",
	}
	return
}
