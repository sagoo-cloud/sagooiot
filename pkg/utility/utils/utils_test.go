package utils

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestFileSize(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		size := FileSize(10240000)
		g.Dump(size)
	})
}

func TestGetLocalIP(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ip, _ := GetLocalIP()
		g.Dump(ip)
	})
}
func TestGetPublicIP(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ip, _ := GetPublicIP()
		g.Dump(ip)
	})
}

func TestRead(t *testing.T) {
	ct, err := ReverseRead("./t.log", 1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range ct {
		fmt.Println(v)
	}
}

func TestM(t *testing.T) {
	// 第一种调用方法
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)

	// 第二种调用方法
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x\n", h.Sum(nil))
}
