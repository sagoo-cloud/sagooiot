package sagooProtocol

type InformRequest struct {
	Id     string `json:"id"`
	Params struct {
		Version string `json:"version"`
		Module  string `json:"module"`
	} `json:"params"`
}

type (
	UpgradeCommandRequest struct {
		Code    string                    `json:"code"`
		Data    UpgradeCommandRequestData `json:"data"`
		Id      int                       `json:"id"`
		Message string                    `json:"message"`
	}

	UpgradeCommandRequestData struct {
		Size       int               `json:"size"`
		Version    string            `json:"version"`
		SignMethod string            `json:"signMethod"`
		DProtocol  string            `json:"dProtocol"`
		Url        string            `json:"url"`
		Sign       string            `json:"sign"`
		Module     string            `json:"module"`
		ExtData    map[string]string `json:"extData"`
	}
)
