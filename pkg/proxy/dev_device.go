package proxy

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	proxyModel "sagooiot/pkg/proxy/model"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDeviceDetail 获取设备详情
func GetDeviceDetail(ctx context.Context, key string) (out *proxyModel.DeviceOutput, err error) {
	detail, err := service.DevDevice().Detail(ctx, key)
	if err != nil {
		return
	}
	if detail != nil {
		if err = gconv.Scan(detail, &out); err != nil {
			return
		}
	}
	return
}

// RunStatus 运行状态
func RunStatus(ctx context.Context, deviceKey string) (out *proxyModel.DeviceRunStatusOutput, err error) {
	deviceRunStatusOutput, err := service.DevDevice().RunStatus(ctx, deviceKey)
	if err != nil {
		return
	}
	if deviceRunStatusOutput != nil {
		if err = gconv.Scan(deviceRunStatusOutput, &out); err != nil {
			return
		}
	}
	return
}

func AllProperty(ctx context.Context, key string) (list []proxyModel.TSLProperty, err error) {
	propertyOut, err := service.DevTSLProperty().AllProperty(ctx, key)
	if err != nil {
		return
	}
	if propertyOut != nil {
		if err = gconv.Scan(propertyOut, &list); err != nil {
			return
		}
	}
	return
}

// Get 设备详情
func Get(ctx context.Context, deviceKey string) (out *proxyModel.DeviceOutput, err error) {
	detail, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		return
	}
	if detail != nil {
		if err = gconv.Scan(detail, &out); err != nil {
			return
		}
	}
	return
}

func ListForPage(ctx context.Context, input *proxyModel.ListDeviceForPageInput) (total, page int, out []*proxyModel.DeviceOutput, err error) {
	var in *model.ListDeviceForPageInput
	if err = gconv.Scan(input, &in); err != nil {
		return
	}
	list, err := service.DevDevice().ListForPage(ctx, in)
	if err != nil {
		return
	}
	if list != nil {
		total, page = list.Total, list.CurrentPage
		if err = gconv.Scan(list.Device, &out); err != nil {
			return
		}
	}
	return
}

// GetPropertyList 设备属性详情列表
func GetPropertyList(ctx context.Context, input *proxyModel.DeviceGetPropertyListInput) (out *proxyModel.DeviceGetPropertyListOutput, err error) {
	var in *model.DeviceGetPropertyListInput
	if err = gconv.Scan(input, &in); err != nil {
		return
	}
	data, err := service.DevDevice().GetPropertyList(ctx, in)
	if err != nil {
		return
	}
	if data != nil {
		err = gconv.Scan(data, &out)
	}
	return
}
