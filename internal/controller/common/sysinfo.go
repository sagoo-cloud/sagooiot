package common

import (
	"context"
	"sagooiot/api/v1/common"
	"sagooiot/internal/service"
)

type cSysInfo struct{}

var SysInfo = cSysInfo{}

// GetSysInfo 系统初始化显示的相关信息
func (c *cSysInfo) GetSysInfo(ctx context.Context, req *common.GetSysInfoReq) (res *common.GetSysInfoRes, err error) {

	out, err := service.SysInfo().GetSysInfo(ctx)
	if err != nil {
		return
	}

	var resData = common.GetSysInfoRes(out)

	res = &resData

	return
}
