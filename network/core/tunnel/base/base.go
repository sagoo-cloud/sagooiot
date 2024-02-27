package base

import (
	"context"
	"io"
	"sagooiot/pkg/iotModel/topicModel"
	"sync"
	"time"
)

// Tunnel 通道
type TunnelInstance interface {
	Write(data []byte) error

	Open(context.Context) error

	Close() error

	Running() bool

	Online() bool

	//Pipe 透传
	Pipe(pipe io.ReadWriteCloser)

	//Ask 发送指令，接收数据
	Ask(cmd []byte, timeout time.Duration) ([]byte, error)
}

type ModelType struct {
	LogType          string
	GetTopicWithInfo func(deviceKey, productKey, identify string) string
	Handle           func(context.Context, topicModel.TopicHandlerData) error
}

type ModelTypeMap struct {
	sync.RWMutex
	MapInfo map[string]ModelType
}

var ModelMapInfo = ModelTypeMap{MapInfo: map[string]ModelType{}}

const (
	UpProperty      = "upProperty"
	UpSetProperty   = "upSetProperty"
	UpEvent         = "upEvent"
	UpBatch         = "upBatch"
	UpServiceOutput = "upServiceOutput"

	UpSetConfig = "upSetConfig"

	UpConfigGet = "upConfigGet"

	UpInform = "upInform"

	UpProcess = "upProcess"

	DownProperty     = "downProperty"
	DownServiceInput = "downServiceInput"
)

func RegisterModelType(name string, modelType ModelType) {
	ModelMapInfo.Lock()
	defer ModelMapInfo.Unlock()
	ModelMapInfo.MapInfo[name] = modelType
}

func GetModelHandle(name string) ModelType {
	ModelMapInfo.Lock()
	defer ModelMapInfo.Unlock()
	return ModelMapInfo.MapInfo[name]
}
