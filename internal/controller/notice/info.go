package notice

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/notice"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var NoticeInfo = cNoticeNoticeInfo{}

type cNoticeNoticeInfo struct{}

// GetNoticeInfoList 获取列表
func (u *cNoticeNoticeInfo) GetNoticeInfoList(ctx context.Context, req *notice.GetNoticeInfoListReq) (res *notice.GetNoticeInfoListRes, err error) {
	var reqData = new(model.GetNoticeInfoListInput)
	if err = gconv.Scan(req, &reqData); err != nil {
		return
	}
	total, currentPage, dataList, err := service.NoticeInfo().GetNoticeInfoList(ctx, reqData)
	res = new(notice.GetNoticeInfoListRes)
	if err = gconv.Scan(dataList, &res.Data); err != nil {
		return
	}
	res.PaginationRes.Total = total
	res.PaginationRes.CurrentPage = currentPage
	return
}

// GetNoticeInfoById 获取指定ID数据
func (u *cNoticeNoticeInfo) GetNoticeInfoById(ctx context.Context, req *notice.GetNoticeInfoByIdReq) (res *notice.GetNoticeInfoByIdRes, err error) {
	data, err := service.NoticeInfo().GetNoticeInfoById(ctx, req.Id)
	res = new(notice.GetNoticeInfoByIdRes)
	err = gconv.Scan(data, &res)
	return
}

// AddNoticeInfo 添加数据
func (u *cNoticeNoticeInfo) AddNoticeInfo(ctx context.Context, req *notice.AddNoticeInfoReq) (res *notice.AddNoticeInfoRes, err error) {
	var data = model.NoticeInfoAddInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	err = service.NoticeInfo().AddNoticeInfo(ctx, data)
	return
}

// EditNoticeInfo 修改数据
func (u *cNoticeNoticeInfo) EditNoticeInfo(ctx context.Context, req *notice.EditNoticeInfoReq) (res *notice.EditNoticeInfoRes, err error) {
	var data = model.NoticeInfoEditInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	err = service.NoticeInfo().EditNoticeInfo(ctx, data)
	return
}

// DeleteNoticeInfo 删除数据
func (u *cNoticeNoticeInfo) DeleteNoticeInfo(ctx context.Context, req *notice.DeleteNoticeInfoReq) (res *notice.DeleteNoticeInfoRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.NoticeInfo().DeleteNoticeInfo(ctx, req.Ids)
	return
}
