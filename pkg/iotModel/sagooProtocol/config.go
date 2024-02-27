package sagooProtocol

type (
	// 配置下发报文
	ConfigPushRequest struct {
		Id      string                `json:"id"`
		Version string                `json:"version"`
		Params  ConfigPushRequestData `json:"params"`
		Method  string                `json:"method"`
	}
	ConfigPushRequestData struct {
		ConfigId      string `json:"configId"`
		ConfigSize    int    `json:"configSize"`
		ConfigContent string `json:"configContent"`
		Sign          string `json:"sign"`
		SignMethod    string `json:"signMethod"`
		Url           string `json:"url"`
		GetType       string `json:"getType"`
	}
	// 配置下发报文响应
	ConfigPushResponse struct {
		Code int                    `json:"code"`
		Data map[string]interface{} `json:"data"`
		Id   string                 `json:"id"`
	}
)

type (
	// 配置获取
	ConfigGetRequest struct {
		Id      string `json:"id"`
		Version string `json:"version"`
		Sys     struct {
			Ack int `json:"ack"`
		} `json:"sys"`
		Params struct {
			ConfigScope string `json:"configScope"`
			GetType     string `json:"getType"`
		} `json:"params"`
		Method string `json:"method"`
	}
	ConfigGetResponse struct {
		Id      string                `json:"id"`
		Version string                `json:"version"`
		Code    int                   `json:"code"`
		Data    ConfigGetResponseData `json:"data"`
	}
	ConfigGetResponseData struct {
		ConfigId      string `json:"configId"`
		ConfigSize    int    `json:"configSize"`
		Sign          string `json:"sign"`
		SignMethod    string `json:"signMethod"`
		Url           string `json:"url"`
		GetType       string `json:"getType"`
		ConfigContent string `json:"configContent"`
	}
)
