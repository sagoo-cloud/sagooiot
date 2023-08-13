package model

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// DataReq 数据请求参数
type DataReq struct {
	Data      []byte
	DataIdent string
}

type SagooMqttData struct {
	Attr       map[string]any `json:"attr"`
	DeviceID   string         `json:"device_id"`
	ReturnTime int64          `json:"return_time"`
}
