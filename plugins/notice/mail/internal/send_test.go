package internal

import (
	"fmt"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"testing"
)

func TestSend(t *testing.T) {
	var msg = model.NoticeInfoData{}
	// 设定相关参数
	opts := []Option{
		MailHost("smtp.qq.com"),
		MailPort(465),
		MailUser("xinjy@qq.com"),
		MailPass("zdqkqrdzplnabiig"),
	}
	m := GetMailChannel(opts...)

	msg.Totag = "[{\"name\":\"mail\",\"value\":\"940290@qq.com\"},{\"name\":\"webhook\",\"value\":\"cccc\"}]"
	msg.MsgTitle = "test sagoo iot msg"
	msg.MsgBody = "this is doc"
	if err := m.Send(msg); err != nil {
		fmt.Println(err)
	}

}
