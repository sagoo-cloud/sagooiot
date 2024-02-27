package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/network/core/server"
	"sagooiot/network/core/server/base"
	"sagooiot/pkg/iotModel/topicModel"
	"sagooiot/pkg/plugins"
	"sagooiot/pkg/plugins/model"
)

// 往通道写数据
func WriteTunnel(ctx context.Context, funcKey string, request topicModel.TopicDownHandlerData, paramsBytes []byte) error {
	t, err := base.GetServerTunnel(ctx, request.DeviceDetail.Key)
	if err != nil {
		return err
	}
	s := server.GetServer(t.ServerId)
	if s == nil {
		return fmt.Errorf("server not found,serverId:%d,deviceKey:%s", t.ServerId, request.DeviceDetail.Key)
	}
	//	获取通道然后写数据，只管写入，不管响应，响应单独走路由
	if request.DeviceDetail.Product.MessageProtocol != consts.DefaultProtocol && request.DeviceDetail.Product.MessageProtocol != "" {
		if plugins.GetProtocolPlugin() == nil {
			return fmt.Errorf("protocol plugin not found")
		}
		//
		var reqData = model.DataReq{}
		reqData.Data = paramsBytes
		pluginData, err := plugins.GetProtocolPlugin().GetProtocolEncodeData(request.DeviceDetail.Product.MessageProtocol, reqData)
		if err != nil {
			g.Log().Errorf(ctx, "get plugin error: %v, deviceKey:%s, data:%s, message ignored", err, request.DeviceDetail.Key, string(paramsBytes))
			return err
		}
		if pluginData.Code != 0 {
			g.Log().Errorf(ctx, "plugin parse error: code:%d message:%s, deviceKey:%s, data:%s, message ignored", pluginData.Code, pluginData.Message, request.DeviceDetail.Key, string(paramsBytes))
			return fmt.Errorf("plugin parse error: code:%d message:%s, deviceKey:%s, data:%s", pluginData.Code, pluginData.Message, request.DeviceDetail.Key, string(paramsBytes))
		}
		paramsBytes, _ = json.Marshal(pluginData.Data)
	}
	return s.Instance.GetTunnel(base.GetTunnelIdByDeviceKey(ctx, request.DeviceDetail.Key)).Write(paramsBytes)
}
