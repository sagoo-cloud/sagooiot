package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/extend"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	networkModel "github.com/sagoo-cloud/sagooiot/network/model"
	"io"
	"sync"
	"time"
)

type tunnelBase struct {
	tunnelId int

	lock sync.Mutex

	link io.ReadWriteCloser

	running bool
	online  bool
	first   bool

	retry      int
	retryTimer *time.Timer

	pipe io.ReadWriteCloser
	data chan []byte
}

func (l *tunnelBase) Running() bool {
	return l.running
}

func (l *tunnelBase) Online() bool {
	return l.online
}

// Close 关闭
func (l *tunnelBase) Close() error {
	if l.retryTimer != nil {
		l.retryTimer.Stop()
	}
	if !l.running {
		return errors.New("tunnel closed")
	}
	TunnelCloseAction(l.tunnelId)
	l.onClose()
	return l.link.Close()
}

func (l *tunnelBase) onClose() {
	l.running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
	TunnelCloseAction(l.tunnelId)
}

// Write 写
func (l *tunnelBase) Write(data []byte) error {
	if !l.running {
		return errors.New("tunnel closed")
	}
	if l.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	_, err := l.link.Write(data)
	return err
}

func (l *tunnelBase) wait(duration time.Duration) ([]byte, error) {
	select {
	case <-time.After(duration):
		return nil, errors.New("超时")
	case buf := <-l.data:
		return buf, nil
	}
}

func (l *tunnelBase) ReadData(ctx context.Context, deviceKey string, data []byte) {
	res := string(data)
	deviceDetail, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		g.Log().Errorf(ctx, "get deviceInfo error: %w,  deviceKey:%s, message ignored", err, deviceDetail.Key)
		return
	}
	productDetail, productDetailErr := service.DevProduct().Get(ctx, deviceDetail.Product.Key)
	if productDetailErr != nil || productDetail == nil {
		g.Log().Errorf(ctx, "find product info error: %w,  productKey:%s, message ignored", productDetailErr, deviceDetail.Product.Key)
		return
	}
	if productDetail.MessageProtocol != "" {
		if extend.GetProtocolPlugin() == nil {
			return
		}
		res, err = extend.GetProtocolPlugin().GetProtocolUnpackData(productDetail.MessageProtocol, data)
		if err != nil {
			g.Log().Errorf(ctx, "get plugin error: %w,  message:%s, message ignored", err, res)
			return
		}
	}
	var reportData networkModel.DefaultMessageType
	if reportDataErr := json.Unmarshal([]byte(res), &reportData); reportDataErr != nil {
		g.Log().Errorf(ctx, "parse data error: %w, topic:%s, message:%s, message ignored", reportDataErr, res)
		return
	}
	messageRouter{
		ctx:          ctx,
		data:         reportData.Data,
		msgType:      reportData.DataType,
		deviceDetail: deviceDetail,
	}.router()

}

func (l *tunnelBase) Ask(cmd []byte, timeout time.Duration) ([]byte, error) {
	if !l.running {
		return nil, errors.New("tunnel closed")
	}

	//堵塞
	l.lock.Lock()
	defer l.lock.Unlock() //自动解锁

	_, err := l.link.Write(cmd)
	if err != nil {
		return nil, err
	}
	return l.wait(timeout)
}

func (l *tunnelBase) Pipe(pipe io.ReadWriteCloser) {
	//关闭之前的透传
	if l.pipe != nil {
		_ = l.pipe.Close()
	}

	l.pipe = pipe
	//传入空，则关闭
	if pipe == nil {
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := pipe.Read(buf)
		if err != nil {
			//if err == io.EOF {
			//	continue
			//}
			//pipe关闭，则不再透传
			break
		}
		//将收到的数据转发出去
		n, err = l.link.Write(buf[:n])
		if err != nil {
			//发送失败，说明连接失效
			_ = pipe.Close()
			break
		}
	}
	l.pipe = nil
}
