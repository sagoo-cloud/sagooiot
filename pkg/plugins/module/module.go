package module

import (
	"encoding/json"
	"github.com/hashicorp/go-plugin"
	"sagooiot/pkg/plugins/model"
)

// HandshakeConfig 握手配置，插件进程和宿主机进程，都需要保持一致
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SAGOO_PLUGIN",
	MagicCookieValue: "sagoo_plugin",
}

// OutJsonRes 输出json字符串结果
func OutJsonRes(code int, message string, data interface{}) string {
	var res = model.JsonRes{}
	res.Code = code
	res.Message = message
	res.Data = data
	jsonByte, _ := json.Marshal(res)
	return string(jsonByte)
}
