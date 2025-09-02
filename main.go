package main

import (
	"sagooiot/command"
	_ "sagooiot/internal/logic"
	_ "sagooiot/internal/packed"
	_ "sagooiot/network/core/logic/model"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"

	"sagooiot/pkg/utility/version"

	"github.com/gogf/gf/v2/os/gctx"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	ctx := gctx.GetInitCtx()
	command.AllSystemInit(ctx)
	app := command.GetServer(ctx)
	cmd := command.NewMainCommand(app.Server)
	cmd.Run(ctx)
}
