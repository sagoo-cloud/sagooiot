package model

import (
	"time"
)

// Product 产品
type Product struct {
	Id           string   `json:"id" xorm:"pk"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"` //厂家
	Version      string   `json:"version"`      //SEMVER
	Protocol     Protocol `json:"protocol"`
	//Tunnel       string `json:"tunnel"` // serial tcp udp ???

	Tags     []string   `json:"tags,omitempty"`
	Pollers  []*Poller  `json:"pollers"`
	Commands []*Command `json:"commands"`

	Created time.Time `json:"created" xorm:"created"`
}

// Device 设备
type Device struct {
	DeviceKey  string `json:"device_key"`
	TunnelId   uint64 `json:"tunnel_id" boltholdIndex:"TunnelId"`
	ProductKey string `json:"product_key"`
	Name       string `json:"name"`
	Station    int    `json:"station"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}
