package mapper

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
)

func StrToPointInterfaceWithError(str string, a interface{}) error {
	return json.Unmarshal([]byte(str), a)
}

func StrToPointInterfaceWithoutError(ctx context.Context, str string, a interface{}) {
	if err := json.Unmarshal([]byte(str), a); err != nil {
		g.Log().Debugf(ctx, "StrToPointInterfaceWithoutError %s", err.Error())
	}
}
