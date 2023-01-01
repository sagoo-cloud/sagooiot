package model

type DefaultReportFormate struct {
	Attr       map[string]interface{} `json:"attr"`
	DeviceId   string                 `json:"device_id"`
	ReturnTime string                 `json:"return_time"`
}
