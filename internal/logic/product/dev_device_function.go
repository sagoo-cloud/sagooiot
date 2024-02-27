package product

import (
	"context"
	"encoding/json"
	"errors"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	dservice "sagooiot/network/core/logic/model/down/service"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/iotModel/topicModel"
)

type sDevDeviceFunction struct{}

func init() {
	service.RegisterDevDeviceFunction(devDeviceFunction())
}

func devDeviceFunction() *sDevDeviceFunction {
	return &sDevDeviceFunction{}
}

// Do 执行设备功能
func (s *sDevDeviceFunction) Do(ctx context.Context, in *model.DeviceFunctionInput) (out *model.DeviceFunctionOutput, err error) {
	device, err := dcache.GetDeviceDetailInfo(in.DeviceKey)
	if dcache.GetDeviceStatus(ctx, in.DeviceKey) != model.DeviceStatusOn {
		err = errors.New("设备不在线")
		return
	}

	var params []byte
	if len(in.Params) > 0 {
		if params, err = json.Marshal(in.Params); err != nil {
			return
		}
	}
	request := topicModel.TopicDownHandlerData{
		DeviceDetail: device,
		PayLoad:      params,
	}

	out = &model.DeviceFunctionOutput{}
	out.Data, err = dservice.ServiceCall(ctx, in.FuncKey, request)

	return
}
