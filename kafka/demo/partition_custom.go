package demo

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// PartitionMurmur2 ...
func PartitionMurmur2() {
	conf := readConfig()
	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Fatal("create producer error", err)
	}
	deliveryCh := make(chan kafka.Event, 1)
	key := "kenh14"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("kenh14's log"),
		Key:   []byte(key),
	}, deliveryCh)

	e := <-deliveryCh
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

func customPartitioner(key string) int32 {
	rand.Seed(time.Now().UnixNano())
	if key == "kenh14" {
		return 1
	}

	acc := 0
	for _, r := range []rune(key) {
		acc += int(math.Abs(float64(r - 'a')))
	}

	return int32(acc % 6)
}

// PartitionCustom ...
func PartitionCustom() {
	conf := readConfig()
	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Fatal("create producer error", err)
	}
	deliveryCh := make(chan kafka.Event, 1)
	key := "kenh14"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: customPartitioner(key),
		},
		Value: []byte("kenh14's log"),
		Key:   []byte(key),
	}, deliveryCh)

	e := <-deliveryCh
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
