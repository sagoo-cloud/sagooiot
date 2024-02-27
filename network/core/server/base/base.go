package base

import (
	"context"
	"sagooiot/network/core/tunnel/base"
)

// Server 通道
type ServerInstance interface {
	Open(ctx context.Context) error
	Close() error
	GetTunnel(id string) base.TunnelInstance
	RemoveTunnel(id string)
	Running() bool
}
