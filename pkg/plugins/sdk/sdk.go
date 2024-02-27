package sdk

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"sagooiot/pkg/plugins/model"
)

func DecodeNoticeData(data []byte) (res model.NoticeData, err error) {
	dataJson := gjson.New(string(data))
	configJson := dataJson.Get("Config")
	sendParamJson := dataJson.Get("SendParam")
	msgJson := dataJson.Get("Msg")

	//解析通知内容数据
	if err = gjson.Unmarshal(msgJson.Bytes(), &res.Msg); err != nil {
		return
	}

	//解析配置数据
	if err = gyaml.DecodeTo(configJson.Bytes(), &res.Config); err != nil {
		return
	}

	//解析参数数据
	res.SendParam = sendParamJson.MapStrAny()

	return
}
