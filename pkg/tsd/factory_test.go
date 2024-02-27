package tsd

import (
	"errors"
	"sagooiot/pkg/tsd/comm"
	"testing"
)

func TestDatabaseFactory(t *testing.T) {

	// 定义数据库连接选项
	option := comm.Option{
		Database:   "sagoo_iot",
		Link:       "root:taosdata@ws(127.0.0.1:6041)/",
		DriverName: "taosWS",
	}

	// 使用工厂函数创建 Tdengine 数据库实例
	td := DatabaseFactory(comm.DBTdEngine, option)

	if td == nil {
		t.Error(errors.New("factory err"))
	}
	td.Close()
	t.Log("创建 Tdengine 数据库实例成功")

	// 定义数据库连接选项
	option2 := comm.Option{
		Database: "sagoo_iot",
		Link:     "http://localhost:8086",
		Org:      "sagoo",
		Token:    "ez4BQ5QQCUpcAp1FDhhdY9jfcvxq2Z9OLkQSuQG_IPOzE9GvGRHfRm_YYwfuHtCaS7TVefxhEnzCOHi_nGtsCw==",
	}
	// 使用工厂函数创建 Influxdb 数据库实例
	idb := DatabaseFactory(comm.DBInfluxdb, option2)
	if td == nil {
		t.Error(errors.New("factory err"))
	}
	idb.Close()
	t.Log("创建 Influxdb 数据库实例成功")
}
