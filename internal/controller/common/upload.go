package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
)

type cUpload struct{}

var Upload = cUpload{}

// SingleImg 上传单图
func (c *cUpload) SingleImg(ctx context.Context, req *common.UploadSingleImgReq) (res *common.UploadSingleRes, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("上传文件必须")
		return
	}
	response, err := service.Upload().UploadFile(ctx, file, consts.CheckFileTypeImg, req.Source)
	if err != nil {
		return
	}
	res = &common.UploadSingleRes{
		UploadResponse: response,
	}
	// 上传第三方
	return
}

// MultipleImg 上传多图
func (c *cUpload) MultipleImg(ctx context.Context, req *common.UploadMultipleImgReq) (res *common.UploadMultipleRes, err error) {
	r := g.RequestFromCtx(ctx)
	files := r.GetUploadFiles("file")
	if len(files) == 0 {
		err = gerror.New("上传文件必须")
		return
	}
	mf, err := service.Upload().UploadFiles(ctx, files, consts.CheckFileTypeImg, req.Source)
	if err != nil {
		return
	}
	res = &mf
	return
}

// SingleFile 上传单文件
func (c *cUpload) SingleFile(ctx context.Context, req *common.UploadSingleFileReq) (res *common.UploadSingleRes, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")
	if file == nil {
		err = gerror.New("上传文件必须")
		return
	}
	response, err := service.Upload().UploadFile(ctx, file, consts.CheckFileTypeFile, req.Source)
	if err != nil {
		return
	}
	res = &common.UploadSingleRes{
		UploadResponse: response,
	}
	return
}

// MultipleFile 上传多文件
func (c *cUpload) MultipleFile(ctx context.Context, req *common.UploadMultipleFileReq) (res *common.UploadMultipleRes, err error) {
	r := g.RequestFromCtx(ctx)
	files := r.GetUploadFiles("file")
	if len(files) == 0 {
		err = gerror.New("上传文件必须")
		return
	}
	mf, err := service.Upload().UploadFiles(ctx, files, consts.CheckFileTypeFile, req.Source)
	if err != nil {
		return
	}
	res = &mf
	return
}
