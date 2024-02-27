package consts

// 消息队列
const (
	QueueRequestLogTopic        = "sagooiot_request_log"           // 访问日志
	QueueDeviceAlarmLogTopic    = "device_alarm_log"               // 设备日志
	QueueDeviceDataSaveTopic    = "task.device.data.save"          // 设备数据保存
	QueueDeviceStatusInfoUpdate = "task.device.status.info.update" // 设备信息更新
)
