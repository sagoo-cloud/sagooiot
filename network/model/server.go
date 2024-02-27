package model

import (
	"regexp"
	"sagooiot/internal/consts"
	"sagooiot/internal/mqtt"
	"strconv"
	"time"
)

type Server struct {
	Id            int               `json:"id"`
	Name          string            `json:"name"`
	Type          string            `json:"type"` //tcp udp
	Addr          string            `json:"addr"`
	Register      RegisterPacket    `json:"register"`
	Heartbeat     HeartBeatPacket   `json:"heartbeat"`
	Options       map[string]string `json:"options"`
	Protocol      Protocol          `json:"protocol"`
	Devices       []DefaultDevice   `json:"devices"` //默认设备
	Disabled      bool              `json:"disabled"`
	Created       time.Time         `json:"created" xorm:"created"`
	IsTls         uint              `json:"isTls" dc:"开启TLS:1=是，0=否"`
	AuthType      int               `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string            `json:"authUser" dc:"认证用户"`
	AuthPasswd    string            `json:"authPasswd" dc:"认证密码"`
	AccessToken   string            `json:"accessToken" dc:"AccessToken"`
	CertificateId int               `json:"certificateId" dc:"证书ID"`
	Stick         string            `json:"stick" dc:"粘包处理方式"`
}

func (s *Server) Open() {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionOpen, strconv.Itoa(s.Id)), nil)
}

// RegisterPacket 注册包
type RegisterPacket struct {
	Regex  string `json:"regex,omitempty"`
	Length int    `json:"length,omitempty"`

	regex *regexp.Regexp
}
