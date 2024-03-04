package dcache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/service"
	"sagooiot/pkg/iotModel"
)

// InertDeviceLog 插入设备日志
func InertDeviceLog(ctx context.Context, logType, deviceKey string, obj interface{}) {
	str, strIsOk := obj.(string)
	content := str
	if !strIsOk {
		objStr, _ := json.Marshal(obj)
		content = string(objStr)
	}
	// 向设备缓存数据库插入数据
	if err := DB().InsertData(context.Background(), deviceKey, iotModel.DeviceLog{
		Ts:      gtime.Now(),
		Device:  deviceKey,
		Type:    logType,
		Content: content,
	}); err != nil {
		g.Log().Debugf(ctx, "Failed to insert data: %v\n", err)
	}
}

// GetDeviceDetailData 获取设备解析后的详细数据
func GetDeviceDetailData(ctx context.Context, deviceKey string, dataType ...string) (res []map[string]iotModel.ReportPropertyNode) {
	// 从设备缓存数据库获取数据
	dataList, err := DB().GetData(ctx, deviceKey)
	if err != nil {
		g.Log().Debugf(ctx, "Failed to get data: %v", err)
		return
	}
	for _, data := range dataList {
		if data == "" {
			continue
		}
		var value = iotModel.DeviceLog{}
		if err := json.Unmarshal([]byte(data), &value); err != nil {
			g.Log().Debugf(ctx, "Failed to unmarshal data: %v", err)
		}
		// 基于物模型解析数据
		dataContent, err := service.DevTSLParse().ParseData(ctx, deviceKey, []byte(value.Content))
		if err != nil {
			continue
		}
		if len(dataContent) == 0 {
			continue
		}
		if len(dataType) > 0 {
			for _, vt := range dataType {
				if vt == value.Type {
					res = append(res, dataContent)
				}
			}
		} else {
			res = append(res, dataContent)
		}
	}
	return
}

// GetDeviceDetailDataByLatest 获取设备解析后的最新一条数据
func GetDeviceDetailDataByLatest(ctx context.Context, deviceKey string) (res iotModel.ReportPropertyData) {
	// 从设备缓存数据库获取数据
	data, err := DB().GetDataByLatest(context.Background(), deviceKey)
	if err != nil || data == "" {
		g.Log().Debugf(ctx, "Failed to get data: %v", err)
		return
	}

	var value = iotModel.DeviceLog{}
	if err := json.Unmarshal([]byte(data), &value); err != nil {
		g.Log().Debugf(ctx, "Failed to unmarshal data: %v", err)
		return
	}

	// 基于物模型解析数据
	res, err = service.DevTSLParse().ParseData(ctx, deviceKey, []byte(value.Content))
	if err != nil {
		g.Log().Debugf(ctx, "Failed to parse data: %v", err)
	}
	return
}

// GetDeviceDetailDataByPage 按分页获取设备详细数据，分页参数： pageNum 为页码， pageSize 为每页数量
func GetDeviceDetailDataByPage(ctx context.Context, deviceKey string, pageNum, pageSize int, dataType ...string) (res []map[string]iotModel.ReportPropertyNode, total, currentPage int) {
	// 获取 list 的名称
	listName := DeviceDataCachePrefix + deviceKey

	// 获取 list 的长度
	num, err := DB().client.LLen(context.Background(), listName).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	total = int(num)

	if pageNum <= 0 {
		pageNum = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}
	currentPage = pageNum
	// 计算分页的起始位置和结束位置
	start := (pageNum - 1) * pageSize
	end := start + pageSize

	if end > total {
		end = total
	}

	// 获取分页数据
	dataList, err := DB().client.LRange(context.Background(), listName, int64(start), int64(end)-1).Result()
	if err != nil {
		g.Log().Debugf(ctx, "Failed to get data: %v", err)
	}
	for _, data := range dataList {
		if data == "" {
			continue
		}
		var value = iotModel.DeviceLog{}
		if err := json.Unmarshal([]byte(data), &value); err != nil {
			g.Log().Debugf(ctx, "Failed to unmarshal data: %v", err)
		}

		// 基于物模型解析数据
		dataContent, err := service.DevTSLParse().ParseData(ctx, deviceKey, []byte(value.Content))
		if err != nil {
			return
		}
		if len(dataContent) == 0 {
			continue
		}
		if len(dataType) > 0 {
			for _, vt := range dataType {
				if vt == value.Type {
					res = append(res, dataContent)
				}
			}
		} else {
			res = append(res, dataContent)
		}
	}
	return
}
