package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestGetWeekDay(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a, b := GetWeekDay()
		g.Dump(a, b)
	})
}

func TestGetQuarterDay(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a, b := GetQuarterDay()
		g.Dump(a, b)
	})
}

func TestGetBetweenDates(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := GetBetweenDates("2022-1-5", "2023-1-20")
		g.Dump(a)
	})
}
