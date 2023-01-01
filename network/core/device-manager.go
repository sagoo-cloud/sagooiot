package core

import (
	"context"
	logicModel "github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
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

func LoadDevices(ctx context.Context) error {
	listDevice, err := service.DevDevice().List(ctx, &logicModel.ListDeviceInput{})
	if err != nil {
		return err
	}
	if len(listDevice) == 0 {
		return nil
	}
	devices := make([]*model.Device, len(listDevice))
	for index, node := range listDevice {
		d := MapperDevice(*node)
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
		allDevices.Store(devices[index].Id, dev)
	}
	return nil
}

func LoadDevice(ctx context.Context, id uint) (*Device, error) {
	deviceInfo, err := service.DevDevice().Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	mDeviceInfo := MapperDevice(*deviceInfo)
	dev, err := NewDevice(ctx, &mDeviceInfo)
	if err != nil {
		return dev, err
	}
	allDevices.Store(id, dev)
	err = dev.Start(ctx)
	return dev, nil
}
