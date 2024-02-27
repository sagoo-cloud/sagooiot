package sagooProtocol

// 相关topci信息
const ()

// 事件上报结构体
type (
	// 事件上报请求报文
	ReportEventReq struct {
		Id      string            `json:"id"`
		Version string            `json:"version"`
		Sys     SysInfo           `json:"sys"`
		Params  ReportEventParams `json:"params"`
	}
	ReportEventParams struct {
		Value    map[string]string `json:"value"`
		CreateAt int64             `json:"time"`
	}
	// 事件上报响应报文
	ReportEventReply struct {
		Code int `json:"code"`
		Data struct {
		} `json:"data"`
		Id      string `json:"id"`
		Message string `json:"message"`
		Method  string `json:"method"`
		Version string `json:"version"`
	}
)
