package model

import (
	"bytes"
	"encoding/hex"
	"regexp"
	"strings"
	"time"
)

type Protocol struct {
	Name    string                 `json:"name"`
	Options map[string]interface{} `json:"options"`
}

/*
boltholdIndex 标签是嵌入式数据库 BoltDB 上面的query layer需要的一个特殊标记
*/
//Tunnel 通道模型
type Tunnel struct {
	Id        uint64          `json:"id"`
	ServerId  int             `json:"server_id" boltholdIndex:"ServerId"`
	Name      string          `json:"name"`
	SN        string          `json:"sn" boltholdIndex:"Addr"`
	Type      string          `json:"type"` //serial tcp-client tcp-server udp-client udp-server server-tcp server-udp
	Addr      string          `json:"addr"`
	Remote    string          `json:"remote"`
	Retry     Retry           `json:"retry"`     //重试
	Heartbeat HeartBeatPacket `json:"heartbeat"` // 主要是检查设备的心跳，过滤掉心跳数据
	Serial    SerialOptions   `json:"serial"`    // 串口通信配置参数
	Protocol  Protocol        `json:"protocol"`  //协议配置，包含协议名称和配置
	//Devices   []DefaultDevice  `json:"devices"` //默认设备
	Disabled bool      `json:"disabled"`
	Last     time.Time `json:"last"`
	Created  time.Time `json:"created" xorm:"created"`
}

type DefaultDevice struct {
	Station    int    `json:"station"`
	ProductKey string `json:"product_key"`
}

type TunnelEx struct {
	Tunnel  `xorm:"extends"`
	Running bool   `json:"running"`
	Online  bool   `json:"online"`
	Server  string `json:"server"`
}

func (tunnel *TunnelEx) TableName() string {
	return "tunnel"
}

type Retry struct {
	Enable  bool `json:"enable"`
	Timeout int  `json:"timeout"`
	Maximum int  `json:"maximum"`
}

// SerialOptions 串口参数
type SerialOptions struct {
	Port     string `json:"port"`      // /dev/tty.usb.. COM1
	BaudRate uint   `json:"baud_rate"` //9600 ... 115200 ...
	DataBits uint   `json:"data_bits"` //5 6 7 8
	StopBits uint   `json:"stop_bits"` //1 2
	Parity   uint   `json:"parity"`    // 0:NONE 1:ODD 2:EVEN
	//RS485    bool   `json:"rs485"`
}

func (p *RegisterPacket) Check(buf []byte) (deviceKey string, checkOk bool) {
	if p.Regex != "" {
		if p.regex == nil {
			p.regex = regexp.MustCompile(p.Regex)
		}
		data := string(buf)
		data = strings.ReplaceAll(data, "\n", "")
		data = strings.ReplaceAll(data, "\r", "")
		re := regexp.MustCompile(p.Regex)
		match := re.FindStringSubmatch(data)
		if match == nil {
			return "", false
		}
		//todo  这里check获取的数据需要再次确认
		return match[1], p.regex.Match(buf)
	}
	if p.Length > 0 {
		if len(buf) != p.Length {
			return "", false
		}
	}
	return string(buf), true
}

// HeartBeatPacket 心跳包
type HeartBeatPacket struct {
	Enable  bool   `json:"enable"`
	Timeout int64  `json:"timeout"`
	Regex   string `json:"regex,omitempty"`
	Hex     string `json:"hex,omitempty"`
	Text    string `json:"text,omitempty"`
	Length  int    `json:"length,omitempty"`

	hex   []byte
	regex *regexp.Regexp
	last  int64
}

// Check 检查
func (p *HeartBeatPacket) Check(buf []byte) bool {

	now := time.Now().Unix()
	if p.last == 0 {
		p.last = now
	}
	if p.last+p.Timeout > now {
		p.last = now
		return false
	}
	p.last = now

	if p.Regex != "" {
		if p.regex == nil {
			p.regex = regexp.MustCompile(p.Regex)
		}
		return p.regex.Match(buf)
	}

	if p.Length > 0 {
		if len(buf) != p.Length {
			return false
		}
	}

	if p.Hex != "" {
		if p.hex == nil {
			//var err error
			p.hex, _ = hex.DecodeString(p.Hex)
		}
		return bytes.Equal(p.hex, buf)
	}

	if p.Text != "" {
		return p.Text == string(buf)
	}

	return true
}
