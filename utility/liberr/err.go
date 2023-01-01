/*
* @desc:错误处理
 */

package liberr

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func ErrIsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

func ValueIsNil(value interface{}, msg string) {
	if g.IsNil(value) {
		panic(msg)
	}
}
