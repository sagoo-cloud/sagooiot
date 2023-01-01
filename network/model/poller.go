package model

// Poller 采集器
type Poller struct {
	Disabled bool   `json:"disabled,omitempty"`
	Interval int    `json:"interval"`
	Code     string `json:"code"`
	Address  string `json:"address"`
	Length   int    `json:"length"`
}
