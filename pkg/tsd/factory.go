package tsd

import (
	"fmt"
	"sagooiot/pkg/tsd/comm"
	"sagooiot/pkg/tsd/internal/influxdb"
	"sagooiot/pkg/tsd/internal/tdengine"
)

// DatabaseFactory 工厂函数，根据数据库类型返回相应的实例
func DatabaseFactory(dbType string, option comm.Option) Database {
	switch dbType {
	case comm.DBTdEngine:
		return &tdengine.TdEngine{Option: option}
	case comm.DBInfluxdb:
		return &influxdb.Influxdb{Option: option}
	default:
		fmt.Println("Unsupported database type.", dbType)
		return nil
	}
}
