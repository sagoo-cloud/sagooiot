package notice

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/notice"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var NoticeConfig = cNoticeNoticeConfig{}

type cNoticeNoticeConfig struct{}

// GetNoticeConfigList 获取列表
func (u *cNoticeNoticeConfig) GetNoticeConfigList(ctx context.Context, req *notice.GetNoticeConfigListReq) (res *notice.GetNoticeConfigListRes, err error) {
	var reqData = new(model.GetNoticeConfigListInput)
	err = gconv.Scan(req, &reqData)
	total, currentPage, dataList, err := service.NoticeConfig().GetNoticeConfigList(ctx, reqData)
	res = new(notice.GetNoticeConfigListRes)
	err = gconv.Scan(dataList, &res.Data)
	res.PaginationRes.Total = total
	res.PaginationRes.CurrentPage = currentPage
	return
}

// GetNoticeConfigById 获取指定ID数据
func (u *cNoticeNoticeConfig) GetNoticeConfigById(ctx context.Context, req *notice.GetNoticeConfigByIdReq) (res *notice.GetNoticeConfigByIdRes, err error) {
	data, err := service.NoticeConfig().GetNoticeConfigById(ctx, req.Id)
	res = new(notice.GetNoticeConfigByIdRes)
	err = gconv.Scan(data, &res)
	return
}

// AddNoticeConfig 添加数据
func (u *cNoticeNoticeConfig) AddNoticeConfig(ctx context.Context, req *notice.AddNoticeConfigReq) (res *notice.AddNoticeConfigRes, err error) {
	var data = model.NoticeConfigAddInput{}
	err = gconv.Scan(req, &data)
	err = service.NoticeConfig().AddNoticeConfig(ctx, data)
	return
}

// EditNoticeConfig 修改数据
func (u *cNoticeNoticeConfig) EditNoticeConfig(ctx context.Context, req *notice.EditNoticeConfigReq) (res *notice.EditNoticeConfigRes, err error) {
	var data = model.NoticeConfigEditInput{}
	err = gconv.Scan(req, &data)
	if err != nil {
		return
	}
	err = service.NoticeConfig().EditNoticeConfig(ctx, data)
	return
}

// DeleteNoticeConfig 删除数据
func (u *cNoticeNoticeConfig) DeleteNoticeConfig(ctx context.Context, req *notice.DeleteNoticeConfigReq) (res *notice.DeleteNoticeConfigRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.NoticeConfig().DeleteNoticeConfig(ctx, req.Ids)
	return
}
