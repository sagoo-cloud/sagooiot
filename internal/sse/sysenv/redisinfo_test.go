package sysenv

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestGetRedisInfo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetRedisInfo()
		g.Dump(a)
	})
}
