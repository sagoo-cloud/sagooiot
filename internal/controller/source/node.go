package source

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/source"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var DataNode = cDataNode{}

type cDataNode struct{}

// 添加数据节点
func (c *cDataNode) Add(ctx context.Context, req *source.DataNodeAddReq) (res *source.DataNodeAddRes, err error) {
	err = service.DataNode().Add(ctx, req.DataNodeAddInput)
	return
}

// 编辑数据节点
func (c *cDataNode) Edit(ctx context.Context, req *source.DataNodeEditReq) (res *source.DataNodeEditRes, err error) {
	err = service.DataNode().Edit(ctx, req.DataNodeEditInput)
	return
}

// 删除数据节点
func (c *cDataNode) Del(ctx context.Context, req *source.DataNodeDelReq) (res *source.DataNodeDelRes, err error) {
	err = service.DataNode().Del(ctx, req.NodeId)
	return
}

// 数据节点列表
func (c *cDataNode) List(ctx context.Context, req *source.DataNodeListReq) (res *source.DataNodeListRes, err error) {
	list, err := service.DataNode().List(ctx, req.SourceId)
	res = &source.DataNodeListRes{
		List: list,
	}
	return
}
