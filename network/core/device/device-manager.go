package device

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/service"
	"sagooiot/network/core/mapper"
	"sagooiot/network/model"
	"sync"
)

var allDevices sync.Map

func GetDevice(id uint64) *Device {
	d, ok := allDevices.Load(id)
	if ok {
		return d.(*Device)
	}
	return nil
}

func RemoveDevice(id uint64) error {
	d, ok := allDevices.LoadAndDelete(id)
	if ok {
		dev := d.(*Device)
		return dev.Stop()
	}
	return nil //error
}

func AddDevice(deviceKey string, device *Device) {
	allDevices.Store(deviceKey, device)
}

func LoadDevices(ctx context.Context) error {
	listDevice, err := service.DevDevice().List(ctx, "", "")
	if err != nil {
		return err
	}
	if len(listDevice) == 0 {
		return nil
	}
	devices := make([]*model.Device, len(listDevice))
	for index, node := range listDevice {
		d := mapper.Device(*node)
		devices[index] = &d
	}
	for index := range devices {
		if devices[index].Disabled {
			continue
		}
		dev, err := NewDevice(ctx, devices[index])
		if err != nil {
			g.Log().Error(ctx, err)
			return nil
		}
		AddDevice(devices[index].ProductKey, dev)
	}
	return nil
}

func LoadDevice(ctx context.Context, key string) (*Device, error) {
	deviceInfo, err := service.DevDevice().Detail(ctx, key)
	if err != nil {
		return nil, err
	}
	mDeviceInfo := mapper.Device(*deviceInfo)
	dev, err := NewDevice(ctx, &mDeviceInfo)
	if err != nil {
		return dev, err
	}
	AddDevice(key, dev)
	err = dev.Start(ctx)
	return dev, nil
}
