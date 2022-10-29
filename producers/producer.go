package producers

import (
	"github.com/RafaelRochaS/journey-api/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	kafkaInstance *kafka.Producer
	event         models.EventObjectDto
}

func NewKafkaProducer() (EventProducer, error) {
	kp := &KafkaProducer{}
	err := kp.createProducer()

	return kp, err
}

func (kp *KafkaProducer) createProducer() error {
	k, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "host1:9092,host2:9092",
		"client.id":         "journey-api-1",
		"acks":              "all"})

	kp.kafkaInstance = k

	return err
}

func (kp *KafkaProducer) HandleEventProduction(event models.EventObjectDto) error {
	kp.event = event
	return nil
}
