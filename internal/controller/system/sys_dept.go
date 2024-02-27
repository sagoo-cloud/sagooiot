package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	systemV1 "sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysDept = cDept{}

type cDept struct{}

func (a *cDept) DeptTree(ctx context.Context, req *systemV1.DeptDoReq) (res *systemV1.DeptDoRes, err error) {
	//获取所有的部门
	out, err := service.SysDept().GetTree(ctx, req.DeptName, req.Status)
	if err != nil {
		return nil, err
	}
	var treeData []*model.DeptRes
	if out != nil {
		if err = gconv.Scan(out, &treeData); err != nil {
			return
		}
	}
	res = &systemV1.DeptDoRes{
		Data: treeData,
	}
	return
}

// AddDept 添加部门
func (a *cDept) AddDept(ctx context.Context, req *systemV1.AddDeptReq) (res *systemV1.AddDeptRes, err error) {
	var input *model.AddDeptInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysDept().Add(ctx, input)
	return
}

// EditDept 编辑部门
func (a *cDept) EditDept(ctx context.Context, req *systemV1.EditDeptReq) (res *systemV1.EditDeptRes, err error) {
	var input *model.EditDeptInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysDept().Edit(ctx, input)
	return
}

// DetailDept 获取部门详情
func (a *cDept) DetailDept(ctx context.Context, req *systemV1.DetailDeptReq) (res *systemV1.DetailDeptRes, err error) {
	data, err := service.SysDept().Detail(ctx, req.DeptId)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailDeptRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &systemV1.DetailDeptRes{
			Data: detailRes,
		}
	}
	return
}

// DelDept 根据ID删除部门
func (a *cDept) DelDept(ctx context.Context, req *systemV1.DelDeptReq) (res *systemV1.DelDeptRes, err error) {
	err = service.SysDept().Del(ctx, req.DeptId)
	return
}
