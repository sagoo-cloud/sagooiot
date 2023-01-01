package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type dingdingChannel struct {
	opts *options
}

var ins *dingdingChannel

var once sync.Once

//GetDingdingChannel 构造方法
func GetDingdingChannel(opts ...Option) *dingdingChannel {
	clusterOpts := options{}
	for _, opt := range opts {
		// 函数指针的赋值调用
		opt(&clusterOpts)
	}
	once.Do(func() {
		ins = &dingdingChannel{}
	})
	ins.opts = &clusterOpts

	return ins
}

// GetAccessToken 获取 access_token
func (d *dingdingChannel) GetAccessToken() (accessToken string, err error) {

	cacheKey := "Dingding" + d.opts.agentID
	//存缓存里获取accessToken
	accessTokenData, _ := gcache.Get(context.TODO(), cacheKey)
	if accessTokenData != nil {
		accessToken = accessTokenData.String()
		return
	}

	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", d.opts.appKey, d.opts.appSecret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Token   string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	if data.Errcode != 0 {
		return "", fmt.Errorf("%d: %s", data.Errcode, data.Errmsg)
	}

	// 钉钉的AccessToken的有效期是2小时，这里设置为缓存1小时
	_, err = gcache.SetIfNotExist(context.TODO(), cacheKey, data.Token, time.Hour)

	return
}

//Send 发送
func (d *dingdingChannel) Send(accessToken string, msg model.NoticeInfoData) (err error) {
	var sendObjectList []model.NoticeSendObject
	err = gjson.DecodeTo(msg.Totag, &sendObjectList)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}

	var touser string
	for _, object := range sendObjectList {
		if object.Name == "dingding" {
			touser = object.Value + ","
		}
	}
	err = d.SendTextMessage(accessToken, touser, "", "", msg.MsgBody, "")

	return
}

// SendTextMessage 发送文本消息
func (d *dingdingChannel) SendTextMessage(accessToken, touser, totag, toparty, message, atMobiles string) error {
	url := fmt.Sprintf("https://oapi.dingtalk.com/message/send?access_token=%s", accessToken)

	data := map[string]interface{}{
		"touser":  touser,
		"toparty": toparty,
		"totag":   totag,
		"msgtype": "text",
		"agentid": d.opts.agentID,
		"text": map[string]string{
			"content": message,
		},
		"at": map[string]interface{}{
			"atMobiles": []string{atMobiles},
			"isAtAll":   false,
		},
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewReader(payload))
	return err
}
