package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func HandleEventProduction(event string) error {
	producer, err := createProducer()

	if err != nil {
		return err
	}

	produce(producer, event)

	return nil
}

func createProducer() (*kafka.Producer, error) {
	k, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "host1:9092,host2:9092",
		"client.id":         "journey-api-1",
		"acks":              "all"})

	return k, err
}

func produce(producer *kafka.Producer, event string) error {
	return nil
}
