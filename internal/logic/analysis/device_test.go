package analysis

import (
	"context"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

// 测试获取设备在线离线统计数据
func TestGetDeviceOnlineOfflineStats(t *testing.T) {
	// 设置缓存适配器
	cache.SetAdapter(context.Background())
	data, err := service.AnalysisDevice().GetDeviceOnlineOfflineCount(context.Background())
	if err != nil {
		t.Log(err)
	}
	t.Log(data)
}

// 统计 本年度，年，月，日总数
func TestGetDeviceDataNumberTotalStats(t *testing.T) {
	cache.SetAdapter(context.Background())
	number, err := service.AnalysisDevice().GetDeviceDataTotalCount(context.Background(), "year")
	if err != nil {
		t.Log(err)
	}
	t.Log("今年总数：", number)

	number, err = service.AnalysisDevice().GetDeviceDataTotalCount(context.Background(), "month")
	if err != nil {
		t.Log(err)
	}
	t.Log("当月总数：", number)

	number, err = service.AnalysisDevice().GetDeviceDataTotalCount(context.Background(), "day")
	if err != nil {
		t.Log(err)
	}
	t.Log("今日总数：", number)

}

// 测试获取设备数据统计数据
func TestGetCountDeviceDataNumberList(t *testing.T) {
	// 设置缓存适配器
	cache.SetAdapter(context.Background())

	// 按年统计，月度数据
	data, err := service.AnalysisDevice().GetDeviceDataCountList(context.Background(), "year")
	if err != nil {
		t.Log(err)
	}
	t.Log(data)

	//按月统计，日度数据，为当月数据
	data, err = service.AnalysisDevice().GetDeviceDataCountList(context.Background(), "month")
	if err != nil {
		t.Log(err)
	}
	t.Log(data)

}
