package notice

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"sagooiot/api/v1/notice"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var NoticeTemplate = cNoticeNoticeTemplate{}

type cNoticeNoticeTemplate struct{}

// GetNoticeTemplateList 获取列表
func (u *cNoticeNoticeTemplate) GetNoticeTemplateList(ctx context.Context, req *notice.GetNoticeTemplateListReq) (res *notice.GetNoticeTemplateListRes, err error) {
	var reqData = new(model.GetNoticeTemplateListInput)
	if err = gconv.Scan(req, &reqData); err != nil {
		return
	}
	total, currentPage, dataList, err := service.NoticeTemplate().GetNoticeTemplateList(ctx, reqData)
	res = new(notice.GetNoticeTemplateListRes)
	err = gconv.Scan(dataList, &res.Data)
	res.PaginationRes.Total = total
	res.PaginationRes.CurrentPage = currentPage
	return
}

// GetNoticeTemplateById 获取指定ID数据
func (u *cNoticeNoticeTemplate) GetNoticeTemplateById(ctx context.Context, req *notice.GetNoticeTemplateByIdReq) (res *notice.GetNoticeTemplateByIdRes, err error) {
	data, err := service.NoticeTemplate().GetNoticeTemplateById(ctx, req.Id)
	res = new(notice.GetNoticeTemplateByIdRes)
	err = gconv.Scan(data, &res)
	return
}

// GetNoticeTemplateByConfigId 获取指定ConfigID数据
func (u *cNoticeNoticeTemplate) GetNoticeTemplateByConfigId(ctx context.Context, req *notice.GetNoticeTemplateByConfigIdReq) (res *notice.GetNoticeTemplateByConfigIdRes, err error) {
	data, err := service.NoticeTemplate().GetNoticeTemplateByConfigId(ctx, req.ConfigId)
	res = new(notice.GetNoticeTemplateByConfigIdRes)
	if data == nil {
		return
	}
	err = gconv.Scan(data, &res)
	return
}

// AddNoticeTemplate 添加数据
func (u *cNoticeNoticeTemplate) AddNoticeTemplate(ctx context.Context, req *notice.AddNoticeTemplateReq) (res *notice.AddNoticeTemplateRes, err error) {
	var data = model.NoticeTemplateAddInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	data.Id = guid.S()
	err = service.NoticeTemplate().AddNoticeTemplate(ctx, data)
	return
}

// EditNoticeTemplate 修改数据
func (u *cNoticeNoticeTemplate) EditNoticeTemplate(ctx context.Context, req *notice.EditNoticeTemplateReq) (res *notice.EditNoticeTemplateRes, err error) {
	var data = model.NoticeTemplateEditInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	//data.UpdateBy = userInfo.Id //如果需要保存信息，把这个打开
	err = service.NoticeTemplate().EditNoticeTemplate(ctx, data)
	return
}

// SaveNoticeTemplate 直接更新数据
func (u *cNoticeNoticeTemplate) SaveNoticeTemplate(ctx context.Context, req *notice.SaveNoticeTemplateReq) (res *notice.SaveNoticeTemplateRes, err error) {
	var data = model.NoticeTemplateAddInput{}
	if err = gconv.Scan(req, &data); err != nil {
		return
	}
	if data.Id == "" {
		data.Id = guid.S()
	}
	err = service.NoticeTemplate().SaveNoticeTemplate(ctx, data)
	return
}

// DeleteNoticeTemplate 删除数据
func (u *cNoticeNoticeTemplate) DeleteNoticeTemplate(ctx context.Context, req *notice.DeleteNoticeTemplateReq) (res *notice.DeleteNoticeTemplateRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.NoticeTemplate().DeleteNoticeTemplate(ctx, req.Ids)
	return
}
