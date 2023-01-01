package internal

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"gopkg.in/gomail.v2"
	"strings"
	"sync"
)

type mailChannel struct {
	opts *options
}

var ins *mailChannel

var once sync.Once

//GetMailChannel 构造方法
func GetMailChannel(opts ...Option) *mailChannel {
	clusterOpts := options{}
	for _, opt := range opts {
		// 函数指针的赋值调用
		opt(&clusterOpts)
	}
	once.Do(func() {
		ins = &mailChannel{}
	})
	ins.opts = &clusterOpts

	return ins
}

//Send 发送
func (m *mailChannel) Send(msg model.NoticeInfoData) (err error) {

	var sendObjectList []model.NoticeSendObject
	err = gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}

	for _, object := range sendObjectList {
		if object.Name == "mail" {
			var data = make(map[string]string)
			data["mailTo"] = object.Value
			data["subject"] = msg.MsgTitle
			data["body"] = msg.MsgBody
			err = m.sendMail(data)
		}
	}
	return err
}

func (m *mailChannel) sendMail(data map[string]string) (err error) {

	mail := gomail.NewMessage()
	//设置发件人
	mail.SetHeader("From", m.opts.mailUser)
	//设置发送给多个用户
	mailArrTo := strings.Split(data["mailTo"], ",")
	mail.SetHeader("To", mailArrTo...)
	//设置邮件主题
	mail.SetHeader("Subject", data["subject"])

	//设置邮件正文
	mail.SetBody("text/html", data["body"])
	d := gomail.NewDialer(m.opts.mailHost, m.opts.mailPort, m.opts.mailUser, m.opts.mailPass)

	err = d.DialAndSend(mail)
	return err
}
