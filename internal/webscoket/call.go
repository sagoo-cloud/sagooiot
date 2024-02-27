package webscoket

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"sagooiot/internal/webscoket/common"
	"sagooiot/internal/webscoket/sys"
)

var (
	ctx = gctx.New()
)

var logic = map[string]map[string]func(ctx *common.WorkContext) (g.Map, error){
	"sys": {
		"ping":  sys.Ping,
		"go":    sys.Go,
		"sync":  sys.Sync,
		"point": sys.Point,
	},
}

func Call(unique interface{}, data *gjson.Json) g.Map {
	cmd := data.Get("cmd", "").String()
	context := &common.WorkContext{
		Data: common.Medium{
			Cmd:      cmd,
			Unique:   unique,
			Original: data,
		},
	}
	var callBack g.Map
	if gstr.Contains(cmd, ".") {
		command := gstr.Explode(".", cmd)
		back := g.Map{}
		for _, cmd := range command[1:] {
			call, err := callFunc(command[0], cmd, context)
			if err != nil {
				g.Log().Debug(ctx, command[0]+"."+cmd+" error:"+err.Error())
				back[command[0]+"."+cmd] = g.Map{
					"error": err.Error(),
				}
			} else {
				back[command[0]+"."+cmd] = call
			}
		}
		callBack = back
	} else {
		callBack = callAllFunc(cmd, context)
	}
	return callBack
}

func callFunc(alias, method string, context *common.WorkContext) (data g.Map, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf(`%v`, e))
			runRecover(err)
		}
		return
	}()
	data, err = logic[alias][method](context)
	return
}

func callAllFunc(alias string, context *common.WorkContext) g.Map {
	back := g.Map{}
	for method := range logic[alias] {
		call, err := callFunc(alias, method, context)
		if err != nil {
			g.Log().Debug(ctx, alias+"."+method+" error:"+err.Error())
			back[alias+"."+method] = g.Map{
				"error": err.Error(),
			}
		} else {
			back[alias+"."+method] = call
		}
	}
	return back
}

func runRecover(err error) {
	g.Log().Cat("logic").Error(ctx, err)
}
