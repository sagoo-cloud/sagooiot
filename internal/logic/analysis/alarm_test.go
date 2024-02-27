package analysis

import (
	"context"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

// 按年统计月度的告警数量
func TestGetDeviceAlertCountStatsByYearMonth(t *testing.T) {
	cache.SetAdapter(context.Background())
	data, err := service.AnalysisAlarm().GetDeviceAlertCountByYearMonth(context.Background(), "2023")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}

// 告警按等级统计
func TestAlarmLevelStats(t *testing.T) {
	cache.SetAdapter(context.Background())

	//指定年份的告警
	data, err := service.AnalysisAlarm().GetAlarmLevelCount(context.Background(), "year", "2023")
	if err != nil {
		t.Error(err)
	}
	t.Log(data)

	//本年度指定月份的
	data1, err := service.AnalysisAlarm().GetAlarmLevelCount(context.Background(), "month", "1")
	if err != nil {
		t.Error(err)
	}
	t.Log(data1)

	//本月指定天的
	data2, err := service.AnalysisAlarm().GetAlarmLevelCount(context.Background(), "day", "11")
	if err != nil {
		t.Error(err)
	}
	t.Log(data2)

}
