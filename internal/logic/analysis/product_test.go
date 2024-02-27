package analysis

import (
	"context"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"testing"
)

// 测试获取产品下设备数
func TestGetDeviceCountForProduct(t *testing.T) {
	// 设置缓存适配器
	cache.SetAdapter(context.Background())
	data, err := service.AnalysisProduct().GetDeviceCountForProduct(context.Background(), "monipower20221103")
	if err != nil {
		t.Log(err)
	}
	t.Log(data)
}
