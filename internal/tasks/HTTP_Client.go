package tasks

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

// GetAccessURL 执行访问URL
func (t TaskJob) GetAccessURL(accessURL string) {
	ctx := context.Background()
	g.Log().Debug(ctx, "访问URL：", accessURL)
	res, err := g.Client().Get(ctx, accessURL)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	defer func(res *gclient.Response) {
		if err := res.Close(); err != nil {
			g.Log().Error(ctx, err)
		}
	}(res)
}
