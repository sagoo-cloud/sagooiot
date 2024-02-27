package tunnel

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/network/core/tunnel/action"
	tunelBase "sagooiot/network/core/tunnel/base"
	networkModel "sagooiot/network/model"
	"sagooiot/pkg/iotModel/topicModel"
	"sagooiot/pkg/jsinterpreter"
	"sagooiot/pkg/plugins"
	"sync"
	"time"
)

type TunnelBase struct {
	TunnelId string
	ServerId int
	Lock     sync.Mutex

	Link io.ReadWriteCloser

	running bool
	online  bool
	First   bool

	retry      int
	retryTimer *time.Timer

	pipe io.ReadWriteCloser
	data chan []byte
}

func (l *TunnelBase) Running() bool {
	return l.running
}

func (l *TunnelBase) SetRunning(isRunning bool) {
	l.running = isRunning
}

func (l *TunnelBase) Online() bool {
	return l.online
}

func (l *TunnelBase) SetOnline(isOnline bool) {
	l.online = isOnline
}

func (l *TunnelBase) GetPipe() io.ReadWriteCloser {
	return l.pipe
}

func (l *TunnelBase) SetPipe(pipe io.ReadWriteCloser) {
	l.pipe = pipe
}

// Close 关闭
func (l *TunnelBase) Close() error {
	if l.retryTimer != nil {
		l.retryTimer.Stop()
	}
	if !l.running {
		return errors.New("tunnel closed")
	}
	action.TunnelCloseAction(l.ServerId, l.TunnelId)
	l.OnClose()
	return l.Link.Close()
}

func (l *TunnelBase) OnClose() {
	l.running = false
	if l.pipe != nil {
		_ = l.pipe.Close()
	}
	action.TunnelCloseAction(l.ServerId, l.TunnelId)
}

// Write 写
func (l *TunnelBase) Write(data []byte) error {
	if !l.running {
		return errors.New("tunnel closed")
	}
	if l.pipe != nil {
		return nil //透传模式下，直接抛弃
	}
	//TODO 需要加入黏包的分割字符，不然客户端会阻塞
	//data = append(data, []byte("\n")...)
	_, err := l.Link.Write(data)
	return err
}

func (l *TunnelBase) Wait(duration time.Duration) ([]byte, error) {
	select {
	case <-time.After(duration):
		return nil, errors.New("超时")
	case buf := <-l.data:
		return buf, nil
	}
}

func (l *TunnelBase) ReadData(ctx context.Context, deviceKey string, data []byte) {
	deviceDetail, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		g.Log().Errorf(ctx, "get deviceInfo error: %v,  deviceKey:%s, message ignored", err, deviceKey)
		return
	}
	productDetail, productDetailErr := service.DevProduct().Detail(ctx, deviceDetail.Product.Key)
	if productDetailErr != nil || productDetail == nil {
		g.Log().Errorf(ctx, "find product info error: %v,  productKey:%s, message ignored", productDetailErr, deviceDetail.Product.Key)
		return
	}
	if deviceDetail != nil && deviceDetail.Status != consts.DeviceStatueOnline {
		if deviceOnlineErr := baseLogic.Online(ctx, networkModel.DeviceOnlineMessage{
			DeviceKey:  deviceKey,
			ProductKey: deviceDetail.Product.Key,
			Timestamp:  time.Now().Unix(),
		}); deviceOnlineErr != nil {
			g.Log().Errorf(ctx, "device online error: %v, deviceKey:%s, message:%s, message ignored", deviceOnlineErr, deviceKey, string(data))
		}
	}
	l.router(ctx, productDetail, deviceDetail, data)
}

func (l *TunnelBase) Ask(cmd []byte, timeout time.Duration) ([]byte, error) {
	if !l.running {
		return nil, errors.New("tunnel closed")
	}

	//堵塞
	l.Lock.Lock()
	defer l.Lock.Unlock() //自动解锁

	_, err := l.Link.Write(cmd)
	if err != nil {
		return nil, err
	}
	return l.Wait(timeout)
}

func (l *TunnelBase) Pipe(pipe io.ReadWriteCloser) {
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
		n, err = l.Link.Write(buf[:n])
		if err != nil {
			//发送失败，说明连接失效
			_ = pipe.Close()
			break
		}
	}
	l.pipe = nil
}

func (l *TunnelBase) router(ctx context.Context, productDetail *model.DetailProductOutput, deviceDetail *model.DeviceOutput, data []byte) {
	res := string(data)
	if productDetail.MessageProtocol != consts.DefaultProtocol && productDetail.MessageProtocol != "" {
		if plugins.GetProtocolPlugin() == nil {
			return
		}
		var err error
		// 通过消息协议插件解析数据
		pluginData, err := plugins.GetProtocolPlugin().GetProtocolDecodeData(productDetail.MessageProtocol, data)
		g.Log().Debug(context.TODO(), "GetProtocolDecodeData", pluginData)
		if err != nil {
			g.Log().Debugf(ctx, "get plugin error: %v, deviceKey:%s, data:%s, message ignored", err, deviceDetail.Key, string(data))
			return
		}
		if pluginData.Code != 0 || pluginData.Data == nil {
			g.Log().Debugf(ctx, "plugin parse error: code:%d message:%s, deviceKey:%s, data:%s, message ignored", pluginData.Code, pluginData.Message, deviceDetail.Key, string(data))
			return
		}
		pluginDataByte, _ := json.Marshal(pluginData.Data)
		res = string(pluginDataByte)
	}
	// 如果有js脚本，根据js脚本处理解析后的数据，处理后的数据数据格式为默认的消息协议格式
	if productDetail.ScriptInfo != "" {
		var runScriptErr error
		res, runScriptErr = jsinterpreter.RunScript(res, productDetail.ScriptInfo)
		if runScriptErr != nil {
			g.Log().Errorf(ctx, "runScriptErr error: %v, deviceKey:%s, data:%s, message ignored", runScriptErr, deviceDetail.Key, string(data))
			return
		}
	}

	var dataInfo = map[string]interface{}{}
	if err := json.Unmarshal([]byte(res), &dataInfo); err != nil {
		g.Log().Errorf(ctx, "json.Unmarshal error: %v, deviceKey:%s, data:%s, message ignored", err, deviceDetail.Key, string(data))
		return
	}
	modelFuncName, ok := dataInfo["model_func_name"].(string)
	if !ok {
		g.Log().Errorf(ctx, "model_func_name not found, deviceKey:%s, data:%s, message ignored", deviceDetail.Key, string(data))
		return
	}
	modelIdentifyName, ok := dataInfo["model_func_identify"].(string)
	if !ok {
		g.Log().Errorf(ctx, "model_func_identify not found, deviceKey:%s, data:%s, message ignored", deviceDetail.Key, string(data))
		return
	}
	handleF := tunelBase.GetModelHandle(modelFuncName)
	if modelFuncName == tunelBase.UpProperty {
		modelIdentifyName = "property"
	}
	if err := handleF.Handle(ctx, topicModel.TopicHandlerData{
		Topic:      handleF.GetTopicWithInfo(productDetail.Key, deviceDetail.Key, modelIdentifyName),
		ProductKey: productDetail.Key,
		DeviceKey:  deviceDetail.Key,
		PayLoad:    []byte(res),
		//ProductDetail: productDetail,
		DeviceDetail: deviceDetail,
	}); err != nil && err.Error() != "ignore" {
		g.Log().Infof(ctx, "handleF error: %v, topic:%s, message:%s, message ignored", err, handleF.GetTopicWithInfo(productDetail.Key, deviceDetail.Key, modelIdentifyName), string(data))
		return
	}
	// 记录原始日志,网关批量的放在网关内部处理
	if handleF.LogType != consts.MsgTypeGatewayBatch {
		baseLogic.InertTdLog(ctx, handleF.LogType, deviceDetail.Key, string(data))
	}
}
