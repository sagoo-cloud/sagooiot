package internal

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

func GetConn(ctx context.Context, tdname string) *sql.DB {
	driver := g.Cfg("tdengine").MustGet(ctx, tdname+".type")
	dsn := g.Cfg("tdengine").MustGet(ctx, tdname+".dsn")
	dbName := g.Cfg("tdengine").MustGet(ctx, tdname+".dbName")

	taos, err := sql.Open(driver.String(), dsn.String()+dbName.String())
	if err != nil {
		panic(tdname + "连接失败" + err.Error())
	}
	return taos
}
