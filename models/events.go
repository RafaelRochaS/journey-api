package models

import "time"

type EventObject struct {
	SensorType SensorType `json:"sensor_type" binding:"required"`
	Measure1   int        `json:"measure_1" binding:"required"`
	Measure2   int        `json:"measure_2" binding:"required"`
	Measure3   int        `json:"measure_3" binding:"required"`
	Measure4   int        `json:"measure_4" binding:"required"`
}

type EventObjectDto struct {
	Event     EventObject
	Timestamp time.Time
	TrxId     string
}
