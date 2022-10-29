package producers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/RafaelRochaS/journey-api/models"
	"github.com/RafaelRochaS/journey-api/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	kafkaInstance *kafka.Producer
	event         string
}

func NewKafkaProducer() (EventProducer, error) {
	kp := &KafkaProducer{}
	err := kp.createProducer()

	return kp, err
}

func (kp *KafkaProducer) createProducer() error {
	k, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": utils.KAFKA_HOSTS,
		"client.id":         utils.KAFKA_CLIENT_ID,
		"acks":              "all"})

	kp.kafkaInstance = k

	return err
}

func (kp *KafkaProducer) HandleEventProduction(event models.EventObjectDto) error {
	stringEvent, err := json.Marshal(event)
	kp.event = string(stringEvent)

	if err != nil {
		return err
	}

	kp.event = string(stringEvent)

	delivery_chan := make(chan kafka.Event, 10000)
	err = kp.kafkaInstance.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &utils.KAFKA_TOPIC, Partition: kafka.PartitionAny},
		Value:          []byte(kp.event),
		Timestamp:      time.Now(),
	},
		delivery_chan,
	)

	go func() {
		for e := range kp.kafkaInstance.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	return err
}
