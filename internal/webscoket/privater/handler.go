package privater

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gorilla/websocket"
	"sagooiot/internal/webscoket"
	"time"
)

type Connection struct {
	Socket    *ghttp.WebSocket
	closed    bool
	Unique    string
	writeCh   chan []byte
	Ctx       context.Context
	CtxClose  func()
	Heartbeat *gtimer.Entry
}

var (
	WebsocketManager *gmap.Map
	workPool         *grpool.Pool
	ctx              = gctx.New()
)

func init() {
	//连接合集-线程安全
	WebsocketManager = gmap.New(true)
	workPool = grpool.New(100)
}

// 关闭连接
func (c *Connection) Close() {
	if gmlock.TryLock(c.Unique) {
		if false == c.closed {
			_ = c.Socket.Close()
			c.closed = true
			//必须优先关闭读写协程
			c.CtxClose()
			close(c.writeCh)
			c.Heartbeat.Close()
			WebsocketManager.Remove(c.Unique)
			g.Log().Debug(ctx, "remove:"+c.Unique)
		}
		gmlock.Unlock(c.Unique)
	}
}

// 连接读-同步响应
func (c *Connection) read() {
	for {
		select {
		case <-c.Ctx.Done():
			goto ERR
		default:
			if c.closed {
				goto ERR
			}
			_, msg, err := c.Socket.ReadMessage()
			if err != nil {
				goto ERR
			}
			data := gjson.New(msg)
			if data.IsNil() {
				goto ERR
			}
			cmd := data.Get("cmd", "").String()
			if len(cmd) == 0 {
				goto ERR
			}
			//client有消息上来重置心跳
			c.Heartbeat.Reset()
			err = workPool.AddWithRecover(ctx, func(ctx context.Context) {
				callBack := webscoket.Call(c.Unique, data)
				if len(callBack) > 0 {
					c.WriteMsg(c.Pack(callBack))
				}
			}, runRecover)
			if err != nil {
				goto ERR
			}
		}
	}
ERR:
	g.Log().Debug(ctx, c.Unique+":ReadSync quit")
	c.Close()
	return
}

func (c *Connection) write() {
	for {
		select {
		case <-c.Ctx.Done():
			goto ERR
		case msg := <-c.writeCh:
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				goto ERR
			}
		}
	}
ERR:
	c.Close()
	g.Log().Debug(ctx, c.Unique+":Write quit")
	return
}
func (c *Connection) Pack(data g.Map) []byte {
	data["pack_timestamp"] = time.Now().Unix()
	encode, err := gjson.Encode(data)
	if err != nil {
		fmt.Println(err)
		return []byte("server error")
	}
	return encode
}
func (c *Connection) WriteMsg(msg []byte) {
	//未关闭连接则发送消息，已关闭连接则后台服务关闭此连接
	if !c.closed {
		c.writeCh <- msg
	} else {
		c.Close()
	}
}

func runRecover(ctx context.Context, err error) {
	g.Log().Cat("websocket").Error(ctx, err)
}
