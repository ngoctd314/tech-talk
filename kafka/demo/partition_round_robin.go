package demo

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// PartitionRoundRobin ...
func PartitionRoundRobin() {
	conf := readConfig()

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Fatal("create producer error", err)
	}
	deliveryChan := make(chan kafka.Event, 1)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("admicro's log"),
	}, deliveryChan)

	e := <-deliveryChan
	switch m := e.(type) {
	case *kafka.Message:
		if m.TopicPartition.Error != nil {
			log.Printf("Failed to deliver message: %v\n", m.TopicPartition)
		} else {
			log.Printf("topic:%s, partition:[%d], offset:%v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}

	producer.Close()
}
