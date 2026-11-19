package tsd

import (
	"fmt"
	"sagooiot/pkg/tsd/comm"
	"sagooiot/pkg/tsd/internal/influxdb"
	"sagooiot/pkg/tsd/internal/iotdb"
	"sagooiot/pkg/tsd/internal/tdengine"
)

// AsIoTDB exposes Tree-model operations to callers with dialect-specific queries.
func AsIoTDB(db Database) (*iotdb.Database, bool) {
	value, ok := db.(*iotdb.Database)
	return value, ok
}

// DatabaseFactory 工厂函数，根据数据库类型返回相应的实例
func DatabaseFactory(dbType string, option comm.Option) Database {
	switch dbType {
	case comm.DBTdEngine:
		return &tdengine.TdEngine{Option: option}
	case comm.DBInfluxdb:
		return &influxdb.Influxdb{Option: option}
	case comm.DBIoTDB:
		return &iotdb.Database{Option: option}
	default:
		fmt.Println("Unsupported database type.", dbType)
		return nil
	}
}
