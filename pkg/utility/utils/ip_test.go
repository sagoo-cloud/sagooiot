package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestIsInRange(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//ipList := []string{
		//	"192.168.0.1/32",
		//	"192.168.0.10-192.168.1.200",
		//	"192.168.aaa.100",
		//}
		ipList := []string{
			"192.168.0.1/32",
			"192.168.0.10-192.168.1.200",
		}

		ip := "192.168.1.100"
		g.Dump(IpInBlackListRange(ip, ipList))
	})
}
