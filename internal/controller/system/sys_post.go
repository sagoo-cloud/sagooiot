package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	systemV1 "sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysPost = cPost{}

type cPost struct{}

func (a *cPost) PostTree(ctx context.Context, req *systemV1.PostDoReq) (res *systemV1.PostDoRes, err error) {
	//获取所有的岗位
	out, err := service.SysPost().GetTree(ctx, req.PostName, req.PostCode, req.Status)
	if err != nil {
		return nil, err
	}
	var treeData []*model.PostRes
	if out != nil {
		if err = gconv.Scan(out, &treeData); err != nil {
			return
		}
	}
	res = &systemV1.PostDoRes{
		Data: treeData,
	}
	return
}

// AddPost 添加岗位
func (a *cPost) AddPost(ctx context.Context, req *systemV1.AddPostReq) (res *systemV1.AddPostRes, err error) {
	var input *model.AddPostInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysPost().Add(ctx, input)
	return
}

// EditPost 编辑岗位
func (a *cPost) EditPost(ctx context.Context, req *systemV1.EditPostReq) (res *systemV1.EditPostRes, err error) {
	var input *model.EditPostInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	err = service.SysPost().Edit(ctx, input)
	return
}

// DetailPost 获取岗位详情
func (a *cPost) DetailPost(ctx context.Context, req *systemV1.DetailPostReq) (res *systemV1.DetailPostRes, err error) {
	data, err := service.SysPost().Detail(ctx, req.PostId)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var detailRes *model.DetailPostRes
		if err = gconv.Scan(data, &detailRes); err != nil {
			return nil, err
		}
		res = &systemV1.DetailPostRes{
			Data: detailRes,
		}
	}
	return
}

// DelPost 根据ID删除岗位
func (a *cPost) DelPost(ctx context.Context, req *systemV1.DelPostReq) (res *systemV1.DelPostRes, err error) {
	err = service.SysPost().Del(ctx, req.PostId)
	return
}
