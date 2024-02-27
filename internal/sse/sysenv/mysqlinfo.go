package sysenv

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"log"
)

// GetMysqlStatusInfo 获取MySql的运行状态信息
func GetMysqlStatusInfo() (data []byte) {
	db := g.DB()
	var queryStatusSQL = "show global status"
	list, err := db.GetAll(context.TODO(), queryStatusSQL)
	if err != nil {
		log.Println(err.Error())
	}
	var tmpData = make(map[string]interface{})
	for _, v := range list {
		tmpData[v["Variable_name"].String()] = v["Value"]
	}

	data, _ = json.Marshal(tmpData)
	return
}

// GetMysqlVersionInfo 获取版本信息
func GetMysqlVersionInfo() (data []byte) {
	db := g.DB()
	var queryVersionSQL = "select version() as version limit 1"
	list, err := db.GetAll(context.TODO(), queryVersionSQL)
	if err != nil {
		log.Println(err.Error())
	}
	var tmpData = make(map[string]interface{})
	for _, v := range list {
		//log.Println(v)
		tmpData = v.Map()
		//tmpData[v["Variable_name"].String()] = v["Value"]
	}

	data, _ = json.Marshal(tmpData)
	return
}

// GetMysqlVariablesInfo 获取variables的运行信息
func GetMysqlVariablesInfo() (data []byte) {
	db := g.DB()
	var queryVariablesSQL = "show global variables"
	list, err := db.GetAll(context.TODO(), queryVariablesSQL)
	if err != nil {
		log.Println(err.Error())
	}
	var tmpData = make(map[string]interface{})
	for _, v := range list {
		tmpData[v["Variable_name"].String()] = v["Value"]
	}

	data, _ = json.Marshal(tmpData)
	return
}
