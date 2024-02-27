package sagooProtocol

// 属性上报结构体
type (
	// 属性上报请求报文
	ReportPropertyReq struct {
		Id      string                 `json:"id"`
		Version string                 `json:"version"`
		Sys     SysInfo                `json:"sys"`
		Params  map[string]interface{} `json:"params"`
		Method  string                 `json:"method"`
	}
	// 属性上报响应报文
	ReportPropertyReply struct {
		Code int `json:"code"`
		Data struct {
		} `json:"data"`
		Id      string `json:"id"`
		Message string `json:"message"`
		Method  string `json:"method"`
		Version string `json:"version"`
	}
)

type (
	// 属性下发平台请求报文
	PropertySetRequest struct {
		Id      string                 `json:"id"`
		Version string                 `json:"version"`
		Params  map[string]interface{} `json:"params"`
		Method  string                 `json:"method"`
	}
	// 属性下发设备响应报文
	PropertySetRes struct {
		Code    int                    `json:"code"`
		Data    map[string]interface{} `json:"data"`
		Id      string                 `json:"id"`
		Message string                 `json:"message"`
		Version string                 `json:"version"`
	}
)
