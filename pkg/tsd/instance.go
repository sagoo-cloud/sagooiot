package tsd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/pkg/tsd/comm"
	"sync"
)

var (
	instances Database
	once      sync.Once
)

// DB 返回一个数据库类型连接的单例
func DB() Database {
	once.Do(func() {
		instances = GetDB()
	})
	return instances
}

// GetDB 返回数据库实例
func GetDB() (db Database) {
	var option = comm.Option{}
	databaseType := g.Cfg().MustGet(context.Background(), "tsd.database", "TdEngine").String()
	switch databaseType {
	case comm.DBTdEngine:
		link := g.Cfg().MustGet(context.Background(), "tsd.tdengine.dsn", "root:taosdata@ws(127.0.0.1:6041)/")
		driverName := g.Cfg().MustGet(context.Background(), "tsd.tdengine.type", "taosWS")
		dbName := g.Cfg().MustGet(context.Background(), "tsd.tdengine.dbName", "sagoo_iot")
		option = comm.Option{
			Database:   dbName.String(),
			Link:       link.String(),
			DriverName: driverName.String(),
		}
	case comm.DBInfluxdb:
		link := g.Cfg().MustGet(context.Background(), "tsd.influxdb.addr", "http://localhost:8086")
		org := g.Cfg().MustGet(context.Background(), "tsd.influxdb.org", "sagoo")
		dbName := g.Cfg().MustGet(context.Background(), "tsd.influxdb.dbName", "sagooiot")
		token := g.Cfg().MustGet(context.Background(), "tsd.influxdb.token", "")
		option = comm.Option{
			Database: dbName.String(),
			Link:     link.String(),
			Org:      org.String(),
			Token:    token.String(),
		}
	case comm.DBIoTDB:
		host := g.Cfg().MustGet(context.Background(), "tsd.iotdb.host", "127.0.0.1").String()
		port := g.Cfg().MustGet(context.Background(), "tsd.iotdb.port", 6667).Int()
		database := g.Cfg().MustGet(context.Background(), "tsd.iotdb.database", "root.sagoo_iot").String()
		username := g.Cfg().MustGet(context.Background(), "tsd.iotdb.username", "root").String()
		password := g.Cfg().MustGet(context.Background(), "tsd.iotdb.password", "root").String()
		maxIdle := g.Cfg().MustGet(context.Background(), "tsd.iotdb.maxIdleConns", 8).Int()
		maxOpen := g.Cfg().MustGet(context.Background(), "tsd.iotdb.maxOpenConns", 32).Int()
		option = comm.Option{
			Host:         host,
			Port:         port,
			Database:     database,
			Username:     username,
			Password:     password,
			MaxIdleConns: maxIdle,
			MaxOpenConns: maxOpen,
		}
	}
	return DatabaseFactory(databaseType, option)
}
