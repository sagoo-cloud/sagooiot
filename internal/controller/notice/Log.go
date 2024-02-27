package notice

import (
	"context"
	"sagooiot/api/v1/notice"
	"sagooiot/internal/service"
)

var NoticeLog = cNoticeNoticeLog{}

type cNoticeNoticeLog struct{}

// 删除日志
func (u *cNoticeNoticeLog) Del(ctx context.Context, req *notice.LogDelReq) (res *notice.LogDelRes, err error) {
	err = service.NoticeLog().Del(ctx, req.Ids)
	return
}

// 通知日志搜索
func (u *cNoticeNoticeLog) Search(ctx context.Context, req *notice.LogSearchReq) (res *notice.LogSearchRes, err error) {
	out, err := service.NoticeLog().Search(ctx, req.NoticeLogSearchInput)
	res = &notice.LogSearchRes{
		NoticeLogSearchOutput: out,
	}
	return
}
