package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"regexp"
	"time"
)

type Server struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"` //tcp udp
	Addr      string            `json:"addr"`
	Register  RegisterPacket    `json:"register"`
	Heartbeat HeartBeatPacket   `json:"heartbeat"`
	Options   map[string]string `json:"options"`
	Protocol  Protocol          `json:"protocol"`
	Devices   []DefaultDevice   `json:"devices"` //默认设备
	Disabled  bool              `json:"disabled"`
	Created   time.Time         `json:"created" xorm:"created"`
}

func (s *Server) Open() {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionOpen, s.Id), nil)
}

// RegisterPacket 注册包
type RegisterPacket struct {
	Regex  string `json:"regex,omitempty"`
	Length int    `json:"length,omitempty"`

	regex *regexp.Regexp
}
