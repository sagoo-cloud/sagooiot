package model

// 设备上报的报文解析为平台统一的消息，消息体结构如下
type DefaultMessageType struct {
	ReturnTime string         `json:"return_time"`
	DataType   string         `json:"data_type"`
	DeviceKey  string         `json:"device_key"`
	Data       map[string]any `json:"data"`
}

// 通用的结构体
type (
	Header struct {
		Async                    bool `json:"async" desc:"是否异步"`
		Timeout                  int  `json:"timeout" desc:"超时时间，单位为毫秒"`
		FragMsgId                int  `json:"fragMsgId" desc:"分片主消息ID，为平台下发消息时的消息ID（messageId"`
		FragNum                  int  `json:"fragNum" desc:"分片总数"`
		FragPart                 int  `json:"fragPart" desc:"当前分片索引"`
		FragLast                 int  `json:"fragLast" desc:"是否为最后一个分片。当无法确定分片数量的时候，可以将分片设置到足够大，最后一个分片设置frag_last=true来完成返回。"`
		KeepOnline               int  `json:"keepOnline" desc:"与DeviceOnlineMessage配合使用,在TCP短链接,保持设备一直在线状态,连接断开不会设置设备离线."`
		KeepOnlineTimeoutSeconds int  `json:"keepOnlineTimeoutSeconds" desc:"指定在线超时时间。在短链接时，如果超过此间隔没有收到消息则认为设备离线。"`
		IgnoreStorage            int  `json:"ignoreStorage" desc:"不存储此消息数据。如：读写属性回复默认也会记录到属性时序数据库中，设置为true后，将不记录。"`
		IgnoreLog                bool `json:"ignoreLog" desc:"不记录此消息的日志。如：设置为true，将不记录此消息的日志。"`
		MergeLatest              bool `json:"mergeLatest" desc:"是否合并最新属性数据。设置此消息头后，将会把最新的消息合并到消息体里（需要开启最新数据存储。"`
	}
	Common struct {
		H         Header
		DeviceKey string
		MessageId string
		Timestamp int64
	}
)

// 读取属性的消息结构体
type (
	ReadPropertyMessage struct {
		Common
		Properties []string
	}
	ReadPropertyMessageReply struct {
		Common
		Success    bool
		Properties map[string]any
	}
)

// 写属性消息结构体
type (
	WritePropertyMessage struct {
		Common
		Properties map[string]any
	}

	WritePropertyMessageReply struct {
		Common
		Success    bool
		Properties map[string]any
	}

	ReportPropertyMessage struct {
		Common
		Properties map[string]any
	}
)

// 功能消息结构体
type (
	FunctionParameter struct {
		Name  string
		Value any
	}
	FunctionInvokeMessage struct {
		Common
		FunctionId string
		Inputs     []FunctionParameter
	}
	FunctionInvokeMessageReply struct {
		Common
		Success bool
		Output  any
	}
)

//设备上报消息结构体

type (
	EventMessage struct {
		H         Header
		Event     string
		Data      any
		Timestamp int64
	}
)

// 设备上线下线结构体
type (
	Message             any
	DeviceOnlineMessage struct {
		DeviceKey string
		Timestamp int64
	}

	DeviceOfflineMessage struct {
		DeviceKey string
		Timestamp int64
	}
)

// 子设备消息结构体
type (
	ChildDeviceMessage struct {
		DeviceKey          string
		Timestamp          int64
		ChildDeviceKey     string
		ChildDeviceMessage Message
	}

	ChildDeviceMessageReply struct {
		DeviceKey          string
		MessageId          string
		ChildDeviceKey     string
		ChildDeviceMessage Message
	}
)

// 设备元信息相关结构体
type (
	UpdateTagMessage struct {
		DeviceKey string
		Timestamp int64
		tags      map[string]any
	}

	DerivedMetadataMessage struct {
		DeviceKey string
		Timestamp int64
		Metadata  string
		All       bool
	}
)

type (
	DeviceRegisterMessage struct {
		DeviceKey string
		Timestamp int64
	}
)

// 更新相关消息结构体
type (
	UpgradeFirmwareMessage struct {
		DeviceKey  string
		Timestamp  int64
		Url        string
		Version    string
		Parameters map[string]any
		Sign       string
		SignMethod string
		FirmwareId string
		Size       float64
	}
	UpgradeFirmwareMessageReply struct {
		DeviceKey string
		Timestamp int64
		Success   bool
	}

	UpgradeFirmwareProgressMessage struct {
		DeviceKey   string
		Progress    int
		Complete    bool
		Version     string
		Success     bool
		ErrorReason string
		FirmwareId  string
	}
)
