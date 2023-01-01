package consts

import "fmt"

type (
	Topic  string
	Action string
)

const (
	TopicDeviceData Topic = "device/+/#"
)

const (
	ActionError   Action = "error"
	ActionOnline  Action = "online"
	ActionOffline Action = "offline"
	ActionOpen    Action = "open"
	ActionClose   Action = "close"
	ActionTunnel  Action = "tunnel"
)

func GetWrapperTopic(topic Topic, action Action, id int) string {
	return fmt.Sprintf(string(topic), id, action)
}

func GetDataBusWrapperTopic(productKey, deviceKey string, topic Topic) string {
	return fmt.Sprintf("/device/%s/%s%s", productKey, deviceKey, topic)
}

const CommonDataBusPrefix = "/device/+/+"

const (
	DataBusOnline  Topic = "/online"
	DataBusOffline Topic = "/offline"

	DataBusEvent Topic = "/message/event/{eventId}"

	DataBusPropertyReport Topic = "/message/property/report"

	DataBusPropertyRead      Topic = "/message/send/property/read"
	DataBusPropertyReadReply Topic = "/message/property/read/reply"

	DataBusPropertyWrite      Topic = "/message/send/property/write"
	DataBusPropertyWriteReply Topic = "/message/property/write/reply"

	DataBusFunctionSend  Topic = "/message/send/function"
	DataBusFunctionReply Topic = "/message/function/reply"

	DataBusRegister   Topic = "/register"
	DataBusUnRegister Topic = "/unregister"

	DataBusChildDeviceMessage      Topic = "/message/children/{childrenDeviceId}/{topic}"
	DataBusChildDeviceMessageReply Topic = "/message/children/reply/{childrenDeviceId}/{topic}"

	DataBusDirect Topic = "/message/direct"
	DataBusUpdate Topic = "/message/tags/update"

	DataBusFirmwarePull      Topic = "/firmware/pull"
	DataBusFirmwarePullReply Topic = "/firmware/pull/reply"

	DataBusFirmwarePush      Topic = "/firmware/push"
	DataBusFirmwarePushReply Topic = "/firmware/push/reply"

	DataBusFirmwareReport Topic = "/firmware/report"

	DataBusFirmwareProgress Topic = "/firmware/progress"

	DataBusLog Topic = "/message/log"

	DataBusMetadataDerived Topic = "/metadata/derived"
)

const (
	DataBusServer Topic = "/system/server/%d/%s"
	DataBusTunnel Topic = "/system/tunnel/%d/%s"
)

var topicToDescMap = map[Topic]string{
	DataBusOnline:                  "设备上线",
	DataBusOffline:                 "设备下线",
	DataBusEvent:                   "事件上报",
	DataBusPropertyRead:            "读取属性",
	DataBusPropertyReadReply:       "读取属性回复",
	DataBusPropertyWrite:           "修改属性",
	DataBusPropertyWriteReply:      "修改属性回复",
	DataBusFunctionSend:            "方法调用",
	DataBusFunctionReply:           "方法调用回复",
	DataBusPropertyReport:          "属性上报",
	DataBusChildDeviceMessage:      "子设备消息",
	DataBusChildDeviceMessageReply: "子设备消息回复",
	DataBusRegister:                "设备注册",
	DataBusUnRegister:              "设备解除注册",
}

const (
	Online                  = "online"
	Offline                 = "offline"
	Event                   = "event"
	PropertyRead            = "property_read"
	PropertyReadReply       = "property_read_reply"
	PropertyWrite           = "property_write"
	PropertyWriteReply      = "property_write_reply"
	FunctionSend            = "function_send"
	FunctionReply           = "function_reply"
	PropertyReport          = "property_report"
	ChildDeviceMessage      = "child_device_message"
	ChildDeviceMessageReply = "child_device_message_reply"
	Register                = "register"
	UnRegister              = "un_register"
)

func GetTopicType(topic Topic) string {
	return topicToDescMap[topic]
}

func GetTopicTypes() []string {
	var topicTypes = make([]string, 0)
	for _, t := range topicToDescMap {
		topicTypes = append(topicTypes, t)
	}
	return topicTypes
}
