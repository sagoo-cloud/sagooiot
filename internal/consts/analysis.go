package consts

const (
	// AnalysisDeviceCountPrefix 设备统计前缀
	AnalysisDeviceCountPrefix  = "analysis:device:count:"
	AnalysisProductCountPrefix = "analysis:product:count:"
	ProductDeviceCount         = "ProductDeviceCount:"

	// 设备统计数据项

	TodayMessageVolume = "TodayMessageVolume:" // 今日消息量
	DeviceTotal        = "DeviceTotal"         // 设备总数
	DeviceDisable      = "DeviceDisable"       // 在线设备数

	AnalysisAlarmCountPrefix = "analysis:alarm:count:"

	AlarmTotal               = "AlarmTotal:"              // 设备总数
	AlarmMonthsMessageVolume = "MonthsMessageVolume:"     // 月度告警
	AlarmTodayMessageVolume  = "TodayMessageVolume"       // 今日告警
	AlarmLevelMessageVolume  = "AlarmLevelMessageVolume:" // 今日告警

)
