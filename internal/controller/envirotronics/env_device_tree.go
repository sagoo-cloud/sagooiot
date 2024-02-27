package envirotronics

import (
	"context"
	"sagooiot/api/v1/envirotronics"
	"sagooiot/internal/service"
)

var EnvDeviceTree = cEnvDeviceTree{}

type cEnvDeviceTree struct{}

func (c *cEnvDeviceTree) Statistic(ctx context.Context, req *envirotronics.EnvDeviceTreeStatisticReq) (res *envirotronics.EnvDeviceTreeStatisticRes, err error) {
	out, err := service.EnvDeviceTree().Statistic(ctx, req.EnvDeviceTreeStatisticInput)
	if err != nil {
		return
	}
	res = &envirotronics.EnvDeviceTreeStatisticRes{
		EnvDeviceTreeStatisticOutput: out,
	}
	return
}
