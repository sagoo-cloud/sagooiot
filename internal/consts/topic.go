package consts

import "fmt"

const (
	MsgTypeOnline             = "设备上线"
	MsgTypeOffline            = "设备下线"
	MsgTypeEvent              = "事件上报"
	MsgTypePropertyRead       = "读取属性"
	MsgTypePropertyReadReply  = "读取属性回复"
	MsgTypePropertyWrite      = "修改属性"
	MsgTypePropertyWriteReply = "修改属性回复"
	MsgTypeFunctionSend       = "方法调用"
	MsgTypeFunctionReply      = "方法调用回复"
	MsgTypePropertyReport     = "属性上报"
	MsgTypePropertySet        = "属性设置"
	MsgTypePropertySetReply   = "属性设置回复"
	MsgTypeGatewayBatch       = "网关批量上报"
	MsgTypeGatewayBatchReply  = "网关批量上报回复"
	MsgTypeRegister           = "设备注册"
	MsgTypeUnRegister         = "设备解除注册"

	MsgTypeDeviceInForm         = "设备上报版本信息"
	MsgTypeDeviceUpgradeProcess = "设备更新进度"

	MsgTypeConfigPush      = "设置远程配置下发"
	MsgTypeConfigPushReply = "设置远程配置下发回复"

	MsgTypeConfigGet = "设置远程请求"
)

// todo delete code
type (
	Topic  string
	Action string
)

const (
	ActionError   Action = "error"
	ActionOnline  Action = "online"
	ActionOffline Action = "offline"
	ActionOpen    Action = "open"
	ActionClose   Action = "close"
	ActionTunnel  Action = "tunnel"
)

func GetWrapperTopic(topic Topic, action Action, id string) string {
	return fmt.Sprintf(string(topic), id, action)
}

const (
	DataBusServer       Topic = "/system/server/%d/%s"
	DataBusTunnel       Topic = "/system/tunnel/%d/%s"
	DataBusServerTunnel Topic = "/system/server/tunnel/%s/%s"

	DataBusUpgradeInfo      Topic = "/upgrade/get"
	DataBusUpgradeInfoReply Topic = "/upgrade/get/reply"
	IssueUpgradeCmd         Topic = "/system/server/upgrade/%s/%s" // 下发ota升级信息
	PostUpgradeResult       Topic = "ota/device/upgrade/+/+"       // 升级结果上报 {version:"1.0.1"}
)

func GetTopicTypes() []string {
	return []string{
		MsgTypeOnline,
		MsgTypeOffline,
		MsgTypeEvent,
		MsgTypePropertyRead,
		MsgTypePropertyReadReply,
		MsgTypePropertyWrite,
		MsgTypePropertyWriteReply,
		MsgTypeFunctionSend,
		MsgTypeFunctionReply,
		MsgTypePropertyReport,
		MsgTypeGatewayBatch,
		MsgTypeGatewayBatchReply,
		MsgTypeRegister,
		MsgTypeUnRegister,
		MsgTypeDeviceInForm,
		MsgTypeDeviceUpgradeProcess,

		MsgTypeConfigPush,
		MsgTypeConfigPushReply,
		MsgTypeConfigGet,
	}
}
