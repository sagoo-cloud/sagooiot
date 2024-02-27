package dcache

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestCountDeviceOnlineNum(t *testing.T) {
	treeObj := CountDeviceOnlineNum()
	g.Dump(treeObj)
}

func TestSearchKey(t *testing.T) {
	data, err := SearchKey("DeviceDetailInfo:")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestGetDeviceDetailInfo(t *testing.T) {
	data, err := GetDeviceDetailInfo("t202201621")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)

}

func TestGetOnlineDeviceList(t *testing.T) {
	data, err := GetOnlineDeviceList()
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)

}
