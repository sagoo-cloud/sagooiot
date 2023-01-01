package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/sagoo-cloud/sagooiot/utility/version"

	_ "github.com/sagoo-cloud/sagooiot/internal/packed"

	_ "github.com/sagoo-cloud/sagooiot/internal/logic"
	_ "github.com/sagoo-cloud/sagooiot/internal/task"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sagoo-cloud/sagooiot/internal/cmd"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	cmd.Main.Run(gctx.New())
}
