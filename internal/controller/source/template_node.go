package source

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/source"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var DataTemplateNode = cDataTemplateNode{}

type cDataTemplateNode struct{}

// 添加数据模型节点
func (c *cDataTemplateNode) Add(ctx context.Context, req *source.DataTemplateNodeAddReq) (res *source.DataTemplateNodeAddRes, err error) {
	err = service.DataTemplateNode().Add(ctx, req.DataTemplateNodeAddInput)
	return
}

// 编辑数据模型节点
func (c *cDataTemplateNode) Edit(ctx context.Context, req *source.DataTemplateNodeEditReq) (res *source.DataTemplateNodeEditRes, err error) {
	err = service.DataTemplateNode().Edit(ctx, req.DataTemplateNodeEditInput)
	return
}

// 删除数据模型节点
func (c *cDataTemplateNode) Del(ctx context.Context, req *source.DataTemplateNodeDelReq) (res *source.DataTemplateNodeDelRes, err error) {
	err = service.DataTemplateNode().Del(ctx, req.Id)
	return
}

// 数据模型节点列表
func (c *cDataTemplateNode) List(ctx context.Context, req *source.DataTemplateNodeListReq) (res *source.DataTemplateNodeListRes, err error) {
	list, err := service.DataTemplateNode().List(ctx, req.Tid)
	res = &source.DataTemplateNodeListRes{
		List: list,
	}
	return
}
