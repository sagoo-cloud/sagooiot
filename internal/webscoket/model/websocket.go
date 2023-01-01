package model

type DataMsg struct {
	Type string      `json:"type" v:"required"`
	Data interface{} `json:"data" v:"required"`
	From string      `json:"name" v:""`
}
