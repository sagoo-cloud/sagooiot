package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysCertificate = cSysCertificate{}

type cSysCertificate struct{}

// GetList 获取列表
func (u *cSysCertificate) GetList(ctx context.Context, req *system.GetCertificateListReq) (res *system.GetCertificateListRes, err error) {
	var input *model.SysCertificateListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, currentPage, out, err := service.SysCertificate().GetList(ctx, input)
	res = new(system.GetCertificateListRes)
	res.PaginationRes.Total = total
	res.PaginationRes.CurrentPage = currentPage
	if out != nil && len(out) > 0 {
		if err = gconv.Scan(out, &res.Info); err != nil {
			return
		}
	}
	return
}

// GetCertificateById 获取指定ID数据
func (u *cSysCertificate) GetCertificateById(ctx context.Context, req *system.GetCertificateByIdReq) (res *system.GetCertificateByIdRes, err error) {
	out, err := service.SysCertificate().GetInfoById(ctx, req.Id)
	var info *model.SysCertificateOut
	if out != nil {
		if err = gconv.Scan(out, &info); err != nil {
			return
		}
	}
	res = &system.GetCertificateByIdRes{
		Info: info,
	}
	return
}

// AddCertificate 添加数据
func (u *cSysCertificate) AddCertificate(ctx context.Context, req *system.AddCertificateReq) (res *system.AddCertificateRes, err error) {
	var input *model.AddSysCertificateListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysCertificate().Add(ctx, input)
	return
}

// EditCertificate 修改数据
func (u *cSysCertificate) EditCertificate(ctx context.Context, req *system.EditCertificateReq) (res *system.EditCertificateRes, err error) {
	var input *model.EditSysCertificateListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysCertificate().Edit(ctx, input)
	return
}

// DeleteCertificate 删除数据
func (u *cSysCertificate) DeleteCertificate(ctx context.Context, req *system.DeleteCertificateReq) (res *system.DeleteCertificateRes, err error) {
	err = service.SysCertificate().Delete(ctx, req.Id)
	return
}

// EditCertificateStatus 修改数据
func (u *cSysCertificate) EditCertificateStatus(ctx context.Context, req *system.EditCertificateStatusReq) (res *system.EditCertificateStatusRes, err error) {
	err = service.SysCertificate().EditStatus(ctx, req.Id, req.Status)
	return
}

// GetCertificateAll 获取所有证书
func (u *cSysCertificate) GetCertificateAll(ctx context.Context, req *system.GetCertificateAllReq) (res *system.GetCertificateAllRes, err error) {
	out, err := service.SysCertificate().GetAll(ctx)
	var info []*model.SysCertificateOut
	if out != nil {
		if err = gconv.Scan(out, &info); err != nil {
			return
		}
	}
	res = &system.GetCertificateAllRes{
		Info: info,
	}
	return
}
