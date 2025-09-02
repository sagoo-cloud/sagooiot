package model

// ProductCount 产品数量统计
type ProductCount struct {
	Total   int `json:"total"`   // 产品总数
	Disable int `json:"disable"` // 禁用产品数
	Enable  int `json:"enable"`  // 启用产品数
	Added   int `json:"added"`   // 新增产品数
}

// DeviceOnlineOfflineCount 设备在线离线状态统计
type DeviceOnlineOfflineCount struct {
	Total   int `json:"total"`   // 设备总数
	Online  int `json:"online"`  // 在线设备数
	Offline int `json:"offline"` // 离线设备数
	Disable int `json:"disable"` // 禁用设备数
}

// CountData 统计数据
type CountData struct {
	Title string
	Value int64
}
