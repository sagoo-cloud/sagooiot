package model

import (
	"time"
)

// Point 数据点
type Point struct {
	Name         string   `json:"name"`
	Label        string   `json:"label"`
	Unit         string   `json:"unit"`
	Type         DataType `json:"type"`
	LittleEndian bool     `json:"le"`
	Precision    int      `json:"precision"`
	Code         string   `json:"code"`
	Address      string   `json:"address"`
	Store        bool     `json:"store"`
	//Default   float64           `json:"default"`
}

type DataPoint struct {
	Value interface{} `json:"value"`
	Time  time.Time   `json:"time"`
}
