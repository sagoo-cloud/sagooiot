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
	"sagooiot/pkg/tsd"
	"sagooiot/pkg/worker"
	"time"
)

var DeviceDataSaveWorker = new(worker.Scheduled)
var deviceDataSaveAggregator *channelx.Aggregator //批量处理器

// TaskDeviceDataTsdSaveRun 设备数据保存到时序数据
func TaskDeviceDataTsdSaveRun() {
	DeviceDataSaveWorker = worker.RegisterProcess(DeviceDataSave)

	batchSize := g.Cfg().MustGet(context.Background(), "tsd.aggregator.batchSize", 1000).Int()                  //批处理大小
	workers := g.Cfg().MustGet(context.Background(), "tsd.aggregator.workers", 100).Int()                       //工作线程数
	channelBufferSize := g.Cfg().MustGet(context.Background(), "tsd.aggregator.channelBufferSize", 50000).Int() //通道缓冲区大小
	lingerTime := g.Cfg().MustGet(context.Background(), "tsd.aggregator.lingerTime", 100).Int()                 //防止数据积压时间

	// 创建聚合器实例
	deviceDataSaveAggregator = channelx.NewAggregator(
		deviceDataSaveBatchProcessFunc,
		channelx.WithBatchSize(batchSize),
		channelx.WithWorkers(workers),
		channelx.WithChannelBufferSize(channelBufferSize),
		channelx.WithLingerTime(time.Duration(lingerTime)*time.Millisecond),
		channelx.WithLogger(nil),
	)
	// 开始聚合器
	deviceDataSaveAggregator.Start()
}

var DeviceDataSave = &qDeviceDataSave{}

type qDeviceDataSave struct{}

// GetTopic 主题
func (q *qDeviceDataSave) GetTopic() string {
	return consts.QueueDeviceDataSaveTopic
}

// Handle 处理消息
func (q *qDeviceDataSave) Handle(ctx context.Context, p worker.Payload) (err error) {
	if p.Payload == nil || q.GetTopic() != p.Group {
		return nil
	}
	var deviceLog = iotModel.DeviceLog{}
	if err := json.Unmarshal(p.Payload, &deviceLog); err != nil {
		g.Log().Debugf(ctx, "DeviceDataSaveWorker Failed to unmarshal data: %v", err)
	}

	//数据进入到批量操作，等待批量处理
	if !deviceDataSaveAggregator.EnqueueWithRetry(deviceLog, 2, 100*time.Millisecond) {
		g.Log().Debug(ctx, "Failed to enqueue item: ", deviceLog)
	}

	return
}

// deviceDataSaveBatchProcessFunc 批处理函数
func deviceDataSaveBatchProcessFunc(items []interface{}) error {
	// 创建数据库连接
	db := tsd.GetDB()
	defer db.Close()
	deviceDataList := make(map[string][]iotModel.ReportPropertyData)
	for _, item := range items {
		var devLog = iotModel.DeviceLog{}
		err := gconv.Scan(item, &devLog)
		if err != nil {
			return err
		}

		// 基于物模型解析数据
		if devLog.Type == consts.MsgTypePropertyReport {
			deviceData, err := service.DevTSLParse().ParseData(context.Background(), devLog.Device, []byte(devLog.Content))
			if err != nil {
				g.Log().Debug(context.Background(), "解析设备日志数据失败:", err, devLog.Content)
				continue
			}
			deviceDataList[devLog.Device] = append(deviceDataList[devLog.Device], deviceData)
		}
	}
	_, err := db.BatchInsertMultiDeviceData(deviceDataList)
	if err != nil {
		g.Log().Debug(context.Background(), "批量插入设备日志数据失败:", err)
	}
	return nil
}
