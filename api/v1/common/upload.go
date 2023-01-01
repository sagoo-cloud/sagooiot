package common

import "github.com/gogf/gf/v2/frame/g"

type UploadSingleImgReq struct {
	g.Meta `path:"/singleImg" tags:"文件上传下载" method:"post" summary:"上传图片"`
}

type UploadSingleFileReq struct {
	g.Meta `path:"/singleFile" tags:"文件上传下载" method:"post" summary:"上传文件"`
}

type UploadSingleRes struct {
	g.Meta `mime:"application/json"`
	UploadResponse
}

type UploadMultipleImgReq struct {
	g.Meta `path:"/multipleImg" tags:"文件上传下载" method:"post" summary:"上传多图片"`
}

type UploadMultipleFileReq struct {
	g.Meta `path:"/multipleFile" tags:"文件上传下载" method:"post" summary:"上传多文件"`
}

type UploadMultipleRes []*UploadResponse

type UploadResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
