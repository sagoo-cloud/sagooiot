package sys

import (
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/webscoket/common"
)

func Ping(ctx *common.WorkContext) (g.Map, error) {
	return g.Map{
		"ping": "pong",
	}, nil
}

func Go(ctx *common.WorkContext) (g.Map, error) {
	return g.Map{
		"go": "run",
	}, nil
}

func Sync(ctx *common.WorkContext) (g.Map, error) {
	return nil, errors.New("sync error")
}

func Point(ctx *common.WorkContext) (g.Map, error) {
	panic("point panic")
	//return nil, errors.New("point error")

}
