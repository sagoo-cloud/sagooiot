package sagooProtocol

// 平台下发topic
const (
	//平台服务调用请求topic /sys/${productKey}/${devicekey}/thing/service/${tsl.service.identifier}
	ServiceCallRegisterSubRequestTopic = "/sys/+/+/thing/service/+"
	//平台服务调用响应求topic /sys/${productKey}/${devicekey}/thing/service/${tsl.service.identifier}_reply
	//这里的请求和响应是不相同的，响应是最后的后缀reply结尾，需要注意
	ServiceCallRegisterSubResponseTopic = "/sys/+/+/thing/service/+"

	//平台设置属性请求topic /sys/${productKey}/${deviceKey}/thing/service/property/set
	PropertySetRegisterSubRequestTopic = "/sys/+/+/thing/service/property/set"
	//平台设置属性响应topic(设备端响应) /sys/${productKey}/thing/service/property/set_reply
	PropertySetRegisterSubResponseTopic = "/sys/+/+/thing/service/property/set_reply"

	//平台下发配置请求topic /sys/${productKey}/${deviceKey}/thing/config/push
	ConfigPushRegisterSubRequestTopic = "/sys/+/+/thing/config/push"
	//平台下发配置响应topic /sys/${productKey}/${deviceKey}/thing/config/push/reply
	ConfigPushRegisterSubResponseTopic = "/sys/+/+/thing/config/push/reply"
)

const (
	//设备上报属性请求topic /sys/${productKey}/${deviceKey}/thing/event/property/post
	PropertyRegisterSubRequestTopic = "/sys/+/+/thing/event/property/post"
	//设备上报属性响应topic(平台响应) /sys/${productKey}/${deviceKey}/thing/event/property/post_reply
	PropertyRegisterPubResponseTopic = "/sys/+/+/thing/event/property/post_reply"

	//设备上报事件请求topic /sys/${productKey}/${deviceKey}/thing/event/${tsl.event.identifier}/post
	EventRegisterSubRequestTopic = "/sys/+/+/thing/event/+/post"
	//设备上报事件响应topic(平台响应) /sys/${productKey}/${deviceKey}/thing/event/${tsl.event.identifier}_reply
	EventRegisterPubResponseTopic = "/sys/+/+/thing/event/+/post_reply"

	// 设备上报批量属性请求topic /sys/${productKey}/${deviceKey}/thing/event/property/pack/post
	BatchRegisterSubRequestTopic = "/sys/+/+/thing/event/property/pack/post"
	// 设备上报批量属性响应topic(平台响应) /sys/${productKey}/${deviceKey}/thing/event/property/pack/post_reply
	BatchRegisterPubResponseTopic = "/sys/+/+/thing/event/property/pack/post_reply"

	//设备主动请求配置信息(设备端发起) /sys/${productKey}/${deviceKey}/thing/config/get
	ConfigGetRequestTopic = "/sys/+/+/thing/config/get"
	//设备主动请求配置信息(平台响应) /sys/${productKey}/${deviceKey}/thing/config/get_reply
	ConfigGetResponseTopic = "/sys/+/+/thing/config/get_reply"
)

//ota相关

// 平台下发
const (
	//推送ota升级包(平台侧发起)/ota/device/upgrade/${productKey}/${deviceKey}
	OtaUpgradeCommandTopic = "/ota/device/upgrade/+/+"
)

// 设备上报
const (
	//设备上报版本信息topic(设备端发起) /ota/device/inform/${productKey}/${deviceKey}
	OtaDeviceInformVersionInfoTopic = "/ota/device/inform/+/+"
	//上报ota升级进度(设备端发起) /ota/device/progress/${productKey}/${deviceKey}
	OtaUpgradeProgressTopic = "/ota/device/progress/+/+"

	//设备请求OTA升级包信息(设备端发起) /sys/${productKey}/${deviceKey}/thing/ota/firmware/get
	OtaDeviceRequestUpgradePackageRequestTopic  = "/sys/+/+/thing/ota/firmware/get"
	OtaDeviceRequestUpgradePackageResponseTopic = "/sys/+/+/thing/ota/firmware/get_reply"

	//设备请求下载文件分片(设备端发起) /sys/${productKey}/${deviceKey}/thing/file/download
	OtaDeviceRequestDownloadFileRequestTopic  = "/sys/+/+/thing/file/download"
	OtaDeviceRequestDownloadFileResponseTopic = "/sys/+/+/thing/file/download_reply"
)
