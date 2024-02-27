package sysenv

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"

	"testing"
)

func TestGetMysqlStatusInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetMysqlStatusInfo()
		g.Dump(a)
	})
}
func TestGetMysqlVersionInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetMysqlVersionInfo()
		g.Dump(a)
	})
}
func TestGetMysqlVariablesInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetMysqlVariablesInfo()
		g.Dump(a)
	})
}
