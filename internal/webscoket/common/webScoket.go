package common

import (
	"github.com/gogf/gf/v2/encoding/gjson"
)

// 公共struct
type WorkContext struct {
	Data Medium
}
type Medium struct {
	Cmd      string
	Unique   interface{}
	Original *gjson.Json
}
