package product

import (
	"context"
	"encoding/json"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	dset "sagooiot/network/core/logic/model/down/property/set"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/iotModel/topicModel"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDevDeviceProperty struct{}

func init() {
	service.RegisterDevDeviceProperty(devDeviceProperty())
}

func devDeviceProperty() *sDevDeviceProperty {
	return &sDevDeviceProperty{}
}

// Set 设备属性设置
func (s *sDevDeviceProperty) Set(ctx context.Context, in *model.DevicePropertyInput) (out *model.DevicePropertyOutput, err error) {
	device, err := dcache.GetDeviceDetailInfo(in.DeviceKey)
	if dcache.GetDeviceStatus(ctx, in.DeviceKey) != model.DeviceStatusOn {
		err = gerror.New("设备不在线")
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

	out = &model.DevicePropertyOutput{}
	if out.Data, err = dset.PropertySet(ctx, request); err != nil {
		return
	}

	// 写日志
	logData := &model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  in.DeviceKey,
		Type:    consts.MsgTypePropertyWrite,
		Content: string(params),
	}
	err = service.TdLogTable().Insert(ctx, logData)

	return
}
