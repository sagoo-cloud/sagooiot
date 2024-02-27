package north

const (
	// 设备上线
	DeviceOnlineMessageTopic = "/message/device/online"
	// 设备下线
	DeviceOfflineMessageTopic = "/message/device/offline"
	// 设备添加
	DeviceAddMessageTopic = "/message/device/add"
	// 设备删除
	DeviceDeleteMessageTopic = "/message/device/delete"
	// 设备上报属性
	PropertyReportMessageTopic = "/message/tsl/receive/property/report"
	// 设备上报事件
	EventReportMessageTopic = "/message/tsl/receive/event/report"
	// 平台调用设备服务请求
	ServiceCallMessageTopic = "/message/tsl/send/service/call"
	// 平台接收到设备服务响应
	ServiceReplyMessageTopic = "/message/tsl/receive/service/reply"
	// 平台设置设备属性
	PropertySetMessageTopic = "/message/tsl/send/property/set"
	// 平台接收到设置设备属性响应
	PropertySetReplyMessageTopic = "/message/tsl/receive/property/reply"
	// 平台下发配置
	ConfigSendMessageTopic = "/message/tsl/send/config"
	// 平台下发配置响应
	ConfigSendMessageReplyTopic = "/message/tsl/receive/config/reply"
	// 平台获取配置
	ConfigGetMessageTopic = "/message/tsl/receive/config/get"
)
