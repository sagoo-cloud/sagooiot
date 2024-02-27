package sysenv

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"runtime"
	"time"
)

// GetHostInfo 获取主机信息
func GetHostInfo() (data []byte) {
	hostInfo, _ := host.Info()
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	tmpData := gconv.Map(hostInfo)
	tmpData["bootTime"] = t

	tmpData["intranet_ip"] = LocalIP
	tmpData["public_ip"] = PublicIP
	data = gconv.Bytes(tmpData)
	return
}

// GetSysLoad 获取系统负载信息
func GetSysLoad() (data []byte) {
	loadInfo, _ := load.Avg()
	data = gconv.Bytes(g.Map{
		"load1":  loadInfo.Load1,
		"load5":  loadInfo.Load5,
		"load15": loadInfo.Load15,
	})
	return
}

type CpuInfo struct {
	Number      int       //cup个数
	Cores       int32     //核数
	UsedPercent []float64 //cpu使用率
	ModelName   string
}

// GetCpuInfo 获取CPU信息
func GetCpuInfo() (data []byte) {
	var CpuInfoData CpuInfo
	cpus, _ := cpu.Info()
	for _, c := range cpus {
		CpuInfoData.Cores = CpuInfoData.Cores + c.Cores
	}
	CpuInfoData.Number = len(cpus)
	percent, _ := cpu.Percent(time.Second, false) //获取CPU使用率
	CpuInfoData.UsedPercent = percent
	CpuInfoData.ModelName = cpus[0].ModelName //CPU型号
	data, _ = json.Marshal(CpuInfoData)
	return
}

// GetMemInfo 获取内存信息
func GetMemInfo() (data []byte) {
	hostInfo, _ := mem.VirtualMemory()
	tmpData := gconv.Map(hostInfo)
	var gomem runtime.MemStats
	runtime.ReadMemStats(&gomem)
	if tmpData == nil {
		tmpData = make(map[string]interface{})
	}
	tmpData["goUsed"] = gomem.Sys
	data, _ = json.Marshal(tmpData)
	return
}

// GetDiskInfo 获取磁盘信息
func GetDiskInfo() (data []byte) {
	diskUsed, _ := disk.Usage("/")
	data = gconv.Bytes(diskUsed)
	return
}

// NetWorkInfo 网速信息
type NetWorkInfo struct {
	Name         string
	Receive      uint64
	Sent         uint64
	ReceiveSpeed uint64
	SentSpeed    uint64
}

// GetNetStatusInfo 获取网络信息
func GetNetStatusInfo() (data []byte) {
	IOCountersStat, _ := net.IOCounters(true)
	netWorkInfo := make([]NetWorkInfo, len(IOCountersStat))
	for i, n := range IOCountersStat {
		netWorkInfo[i].Name = n.Name
		netWorkInfo[i].Receive = n.BytesRecv
		netWorkInfo[i].Sent = n.BytesSent
		if netWorkInfo != nil && len(netWorkInfo) > i {
			netWorkInfo[i].ReceiveSpeed = n.BytesRecv - netWorkInfo[i].Receive
			netWorkInfo[i].SentSpeed = n.BytesSent - netWorkInfo[i].Sent
		}
	}
	data, _ = json.Marshal(netWorkInfo)
	return
}
