package topicModel

import "sagooiot/internal/model"

type TopicHandlerData struct {
	Topic        string
	ProductKey   string
	DeviceKey    string
	PayLoad      []byte
	DeviceDetail *model.DeviceOutput
}

type TopicDownHandlerData struct {
	PayLoad      []byte
	DeviceDetail *model.DeviceOutput
}
