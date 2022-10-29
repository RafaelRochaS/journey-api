package producers

import "github.com/RafaelRochaS/journey-api/models"

type EventProducer interface {
	HandleEventProduction(models.EventObjectDto) error
}
