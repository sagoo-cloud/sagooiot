package model

import (
	"context"
	"sagooiot/network/core/logic/model/up/event"
	"sagooiot/network/core/logic/model/up/property/batch"
	"sagooiot/network/core/logic/model/up/property/reporter"
	"sagooiot/network/core/logic/model/up/property/set"
	"sagooiot/network/core/logic/model/up/service"
)

func InitCoreLogic(ctx context.Context) error {
	for _, v := range []func() error{
		event.Init,
		batch.Init,
		reporter.Init,
		set.Init,
		service.Init,
	} {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}
