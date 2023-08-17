package model

import "encoding/gob"

func init() {
	gob.Register(SagooMqttModel{})
}

// SagooMqttModel 主结构
type (
	SagooMqttModel struct {
		Id                string           `json:"id"`
		Version           string           `json:"version"`
		Sys               SysInfo          `json:"sys"`
		Params            map[string]Param `json:"params"`
		Method            string           `json:"method"`
		ModelFuncName     string           `json:"model_func_name"`
		ModelFuncIdentify string           `json:"model_func_identify"`
	}

	SysInfo struct {
		Ack int `json:"ack"`
	}

	// Param 属性
	Param struct {
		Value interface{} `json:"value"`
		Time  int64       `json:"time"`
	}
)
