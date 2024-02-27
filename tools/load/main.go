package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 结构体定义
type SysInfo struct {
	Ack int `json:"ack"`
}

type PropertyNode struct {
	Value      interface{} `json:"value"`
	CreateTime int64       `json:"time"`
}

// 属性上报结构体
type ReportPropertyReq struct {
	Id      string                 `json:"id"`
	Version string                 `json:"version"`
	Sys     SysInfo                `json:"sys"`
	Params  map[string]interface{} `json:"params"`
	Method  string                 `json:"method"`
}

type ReportPropertyReply struct {
	Code    int      `json:"code"`
	Data    struct{} `json:"data"`
	Id      string   `json:"id"`
	Message string   `json:"message"`
	Method  string   `json:"method"`
	Version string   `json:"version"`
}

const (
	MqttClientIdPrefix = "SagooIOT-Tools-Load"
)

// 全局变量
var (
	h bool
	d = flag.Int("d", 5, "模拟的设备数")
	m = flag.Int("m", 2, "设备多少秒发送一次")
)

var address, mqttUserName, mqttPassword string

func main() {
	flag.Usage = usage
	flag.Parse()
	if h {
		flag.Usage()
	}

	NumGoroutines := *d
	OneMessageDuration := *m
	if OneMessageDuration < 1 {
		fmt.Println("设备发送间隔不能小于1秒")
		return
	}

	address = g.Cfg().MustGet(context.Background(), "mqtt.addr", "127.0.0.1:1883").String()
	mqttUserName = g.Cfg().MustGet(context.Background(), "mqtt.auth.userName", "").String()
	mqttPassword = g.Cfg().MustGet(context.Background(), "mqtt.auth.userPassWorld", "").String()

	initDeviceDelays(NumGoroutines, OneMessageDuration)
}

func initDeviceDelays(numDevices, messageInterval int) {
	var wg sync.WaitGroup
	var messageCount int32
	for i := 1; i <= numDevices; i++ {
		wg.Add(1)
		// 为每个设备计算一个随机的起始延迟，以确保它们的启动时间分散
		startDelay := rand.Intn(messageInterval * 1000) // 随机延迟，最大不超过消息间隔
		go func(deviceID, delay int) {
			defer wg.Done()
			time.Sleep(time.Duration(delay) * time.Millisecond) // 等待起始延迟
			fmt.Println("设备", deviceID, "已初始化", time.Now())

			sendMessages(deviceID, messageInterval, &messageCount)

		}(i, startDelay)
	}
	go monitorMessageCount(&messageCount)

	wg.Wait() // 等待所有设备完成初始化
}

func sendMessages(id, messageInterval int, messageCount *int32) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(address)
	opts.SetClientID(fmt.Sprintf("%s_%d", MqttClientIdPrefix+guid.S(), id))
	opts.SetUsername(mqttUserName)
	opts.SetPassword(mqttPassword)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(1 * time.Second)
	opts.SetKeepAlive(30 * time.Second)
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer mqttClient.Disconnect(250)

	ticker := time.NewTicker(time.Duration(messageInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		data, _ := json.Marshal(&ReportPropertyReq{
			Id:      guid.S(),
			Version: "1.0",
			Sys: SysInfo{
				Ack: 0,
			},
			Params: map[string]interface{}{
				"va": PropertyNode{
					Value:      randFloatNum(180.1000, 230.5000, 4),
					CreateTime: time.Now().Unix(),
				},
				"vb": PropertyNode{
					Value:      randFloatNum(190.1000, 230.5000, 4),
					CreateTime: time.Now().Unix(),
				},
				"vc": PropertyNode{
					Value:      randFloatNum(185.1000, 230.5000, 4),
					CreateTime: time.Now().Unix(),
				},
			},
			Method: "thing.event.property.post",
		})

		deviceNum := leftPad(id, 5)
		topic := fmt.Sprintf("/sys/monipower20221103/t2022%s/thing/event/property/post", deviceNum)

		token := mqttClient.Publish(topic, 2, false, data)
		token.Wait()
		if token.Error() != nil {
			fmt.Printf("携程 %d 发送消息失败: %s\n", id, token.Error())
		} else {
			atomic.AddInt32(messageCount, 1)
		}
	}
}

func monitorMessageCount(messageCount *int32) {
	for {
		time.Sleep(time.Second)
		mc := atomic.LoadInt32(messageCount)
		fmt.Printf("每秒发送消息数量: %d time:%s\n", mc, time.Now().Format("2006-01-02 15:04:05"))
		atomic.StoreInt32(messageCount, 0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `SagooIOT Device Load Version: 0.0.1
Usage: load [-h] [-g device number] [-m send interval]

Options:
`)
	flag.PrintDefaults()
}

func randFloatNum(min, max float64, precision int) float64 {
	randomNum := min + rand.Float64()*(max-min+1e-10)
	result := math.Round(randomNum*math.Pow10(precision)) / math.Pow10(precision)
	resultStr := strconv.FormatFloat(result, 'f', precision, 64)
	f, err := strconv.ParseFloat(resultStr, 64)
	if err != nil {
		return 0
	}
	return f
}

func leftPad(num, digit int) string {
	str := fmt.Sprintf("%d", num)
	padLen := digit - len(str)
	padStr := strings.Repeat("0", padLen)
	return padStr + str
}
