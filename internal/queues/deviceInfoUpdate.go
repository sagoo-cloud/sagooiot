package queues

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
	"sagooiot/pkg/channelx"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/worker"
	"time"
)

var DeviceStatusInfoUpdateWorker = new(worker.Scheduled)
var deviceInfoUpdateAggregator *channelx.Aggregator //批量处理器

// DeviceInfoUpdateRun 更新设备状态信息，设备上线、离线、注册，更新数据库
func DeviceInfoUpdateRun() {
	DeviceStatusInfoUpdateWorker = worker.RegisterProcess(DeviceStatusInfoUpdate)

	batchSize := 200 // 批处理大小
	workers := 200
	channelBufferSize := 50000 // 通道缓冲区大小
	lingerTime := 100 * time.Millisecond

	// 创建聚合器实例
	deviceInfoUpdateAggregator = channelx.NewAggregator(
		deviceInfoUpdateBatchProcessFunc,
		channelx.WithBatchSize(batchSize),
		channelx.WithWorkers(workers),
		channelx.WithChannelBufferSize(channelBufferSize),
		channelx.WithLingerTime(lingerTime),
		channelx.WithLogger(nil),
	)
	// 开始聚合器
	deviceInfoUpdateAggregator.Start()
}

// DeviceStatusInfoUpdate 系统日志
var DeviceStatusInfoUpdate = &qDeviceStatusInfoUpdate{}

type qDeviceStatusInfoUpdate struct{}

// GetTopic 主题
func (q *qDeviceStatusInfoUpdate) GetTopic() string {
	return consts.QueueDeviceStatusInfoUpdate
}

// Handle 处理消息
func (q *qDeviceStatusInfoUpdate) Handle(ctx context.Context, p worker.Payload) (err error) {
	if p.Payload == nil || q.GetTopic() != p.Group {
		return nil
	}
	var data iotModel.DeviceStatusLog
	if err = json.Unmarshal(p.Payload, &data); err != nil {
		return err
	}
	//数据进入到批量操作，等待批量处理
	if !deviceInfoUpdateAggregator.EnqueueWithRetry(data, 3, 100*time.Millisecond) {
		g.Log().Debug(ctx, "Failed to enqueue item: ", data)
	}
	return
}

// deviceInfoUpdateBatchProcessFunc 批处理函数
func deviceInfoUpdateBatchProcessFunc(items []interface{}) (err error) {
	var data []iotModel.DeviceStatusLog
	err = gconv.Scan(items, &data)
	err = service.DevDevice().BatchUpdateDeviceStatusInfo(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}
