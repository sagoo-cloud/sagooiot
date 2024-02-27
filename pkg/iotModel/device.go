package iotModel

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// 设备属性上报
type (
	PropertyReportMessage struct {
		Properties map[string]ReportPropertyNode `json:"properties"` //map类型，属性列表，key为属性标识，value为属性值
	}
)

// 上报属性数据
type ReportPropertyData map[string]ReportPropertyNode

// 属性值
type ReportPropertyNode struct {
	Value      any   // 属性值
	CreateTime int64 // 上报时间
}

// 上报事件数据
type ReportEventData struct {
	Key   string           // 事件标识
	Param ReportEventParam // 事件输出参数
}

// 事件输出参数
type ReportEventParam struct {
	Value      map[string]any // 事件输出参数
	CreateTime int64          // 上报时间
}

// 设备上下线状态
type ReportStatusData struct {
	Status     string // 状态：online、offline
	CreateTime int64  // 上下线时间
}

type DevicePropertiy struct {
	Key   string      `json:"key" dc:"属性标识"`
	Name  string      `json:"name" dc:"属性名称"`
	Type  string      `json:"type" dc:"属性值类型"`
	Unit  string      `json:"unit" dc:"属性值单位"`
	Value *gvar.Var   `json:"value" dc:"属性值"`
	List  []*gvar.Var `json:"list" dc:"当天属性值列表"`
}

// DeviceLog 设备日志
type DeviceLog struct {
	Ts      *gtime.Time `json:"ts" dc:"时间"`
	Device  string      `json:"device" dc:"设备标识"`
	Type    string      `json:"type" dc:"日志类型"`
	Content string      `json:"content" dc:"日志内容"`
}

// 设备上线
type DeviceOnlineMessage struct {
	Timestamp int64  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	Desc      string `json:"desc"`      //string类型，描述信息
}

// 设备下线
type DeviceOfflineMessage struct {
	Timestamp int64  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	Desc      string `json:"desc"`      //string类型，描述信息
}

// 添加设备
type DeviceAddMessage struct {
	Timestamp int64  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	Desc      string `json:"desc"`      //string类型，描述信息
}

// 删除设备
type DeviceDeleteMessage struct {
	Timestamp int64  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	Desc      string `json:"desc"`      //string类型，描述信息
}

// 设备事件上报
type (
	EventReportMessage struct {
		EventId   string                 `json:"eventId"`   //string类型，事件标识
		Events    map[string]interface{} `json:"events"`    //map类型，事件列表，key为事件标识，value为事件值
		Timestamp int64                  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	}
)

// 设备服务调用请求
type (
	ServiceCallMessage struct {
		ServiceId string                 `json:"serviceId"` //string类型，服务标识
		Params    map[string]interface{} `json:"params"`    //map类型，服务参数列表，key为参数标识，value为参数值
		Timestamp int64                  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	}
)

// 平台接收到设备服务响应
type (
	ServiceCallReplyMessage struct {
		ServiceId string                 `json:"serviceId"` //string类型，服务标识
		Code      int                    `json:"code"`      //int类型，响应码
		Data      map[string]interface{} `json:"data"`      //map类型，响应数据列表，key为数据标识，value为数据值
		Timestamp int64                  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	}
)

// 平台设置设备属性
type (
	PropertySetMessage struct {
		Properties map[string]interface{} `json:"properties"` //map类型，属性列表，key为属性标识，value为属性值
		Timestamp  int64                  `json:"timestamp"`  //int64类型，时间戳，单位为毫秒
	}
)

// 平台接收到设置设备属性响应
type (
	PropertySetReplyMessage struct {
		Code      int                    `json:"code"`      //int类型，响应码
		Data      map[string]interface{} `json:"data"`      //map类型，响应数据列表，key为数据标识，value为数据值
		Timestamp int64                  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	}
)

// 平台接收到配置设置响应
type (
	ConfigSetReplyMessage struct {
		Code      int                    `json:"code"`      //int类型，响应码
		Data      map[string]interface{} `json:"data"`      //map类型，响应数据列表，key为数据标识，value为数据值
		Timestamp int64                  `json:"timestamp"` //int64类型，时间戳，单位为毫秒
	}
)

// 平台接收到配置获取请求

type (
	ConfigGetMessage struct {
		ConfigScope string `json:"configScope"`
		GetType     string `json:"getType"`
	}
)

// DeviceStatusLog 设备状态动态变化日志
type DeviceStatusLog struct {
	DeviceKey string
	Status    int       //设备状态
	Timestamp time.Time //数据时间
}
