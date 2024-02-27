package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/module/hello/service"

	"sagooiot/module/hello/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Get("name")

	sysContext := g.RequestFromCtx(ctx)
	g.Log().Debug(ctx, sysContext)

	data, err := service.Hello().Speak(ctx, "hello===========")
	if err != nil {
		g.RequestFromCtx(ctx).Response.Writeln(data)

	}

	return
}
