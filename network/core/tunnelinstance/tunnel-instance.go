package tunnelinstance

import (
	"context"
	"io"
	"time"
)

// Tunnel 通道
type TunnelInstance interface {
	Write(data []byte) error

	Open(context.Context) error

	Close() error

	Running() bool

	Online() bool

	//Pipe 透传
	Pipe(pipe io.ReadWriteCloser)

	//Ask 发送指令，接收数据
	Ask(cmd []byte, timeout time.Duration) ([]byte, error)
}
