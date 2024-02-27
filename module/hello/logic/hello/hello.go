package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/module/hello/service"
)

type sHello struct{}

func init() {
	service.RegisterHello(helloNew())
}

func helloNew() *sHello {
	return &sHello{}
}

func (s *sHello) Speak(ctx context.Context, doc string) (out string, err error) {
	g.Log().Debug(ctx, "hello", doc)
	return
}
