package models

type SensorType int

const (
	SENSOR_1 = iota
	SENSOR_2
	SENSOR_3
	SENSOR_4
)

func (st SensorType) String() string {
	switch st {
	case SENSOR_1:
		return "Sensor 1"
	case SENSOR_2:
		return "Sensor 2"
	case SENSOR_3:
		return "Sensor 3"
	case SENSOR_4:
		return "Sensor 4"
	}

	return "Unknown"
}
