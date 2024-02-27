package sagooProtocol

type (
	// ack 标记是否需要回复，1需要回复，0不需要回复
	SysInfo struct {
		Ack int `json:"ack"`
	}
	// 属性节点，包含属性值和时间
	PropertyNode struct {
		Value      interface{} `json:"value"`
		CreateTime int64       `json:"time"`
	}
)

const (
	NeedAck = 1
)
