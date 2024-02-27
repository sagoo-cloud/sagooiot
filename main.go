package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
	_ "sagooiot/internal/logic"
	_ "sagooiot/internal/packed"
	_ "sagooiot/network/core/logic/model"

	"github.com/gogf/gf/v2/os/gctx"
	"sagooiot/internal/cmd"
	"sagooiot/pkg/utility/version"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	ctx := gctx.GetInitCtx()
	cmd.AllSystemInit(ctx)
	cmd.Main.Run(ctx)
}
