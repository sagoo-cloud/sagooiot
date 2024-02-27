package common

import "github.com/gogf/gf/v2/frame/g"

type UploadSingleImgReq struct {
	g.Meta `path:"/singleImg" tags:"文件上传下载" method:"post" summary:"上传图片"`
	Source int `json:"source" dc:"本地-0、腾讯云-1、阿里云-2、七牛云-3、MinIO-4"`
}

type UploadSingleFileReq struct {
	g.Meta `path:"/singleFile" tags:"文件上传下载" method:"post" summary:"上传文件"`
	Source int `json:"source" dc:"本地-0、腾讯云-1、阿里云-2、七牛云-3、MinIO-4"`
}

type UploadSingleRes struct {
	g.Meta `mime:"application/json"`
	UploadResponse
}

type UploadMultipleImgReq struct {
	g.Meta `path:"/multipleImg" tags:"文件上传下载" method:"post" summary:"上传多图片"`
	Source int `json:"source" dc:"本地-0、腾讯云-1、阿里云-2、七牛云-3、MinIO-4"`
}

type UploadMultipleFileReq struct {
	g.Meta `path:"/multipleFile" tags:"文件上传下载" method:"post" summary:"上传多文件"`
	Source int `json:"source" dc:"本地-0、腾讯云-1、阿里云-2、七牛云-3、MinIO-4"`
}

type UploadMultipleRes []*UploadResponse

type UploadResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
