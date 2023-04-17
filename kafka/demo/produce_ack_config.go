package demo

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ProduceWithAck ...
func ProduceWithAck() {
	now := time.Now()
	defer func() {
		redLogBold.Printf("all producer are completed after %fs", time.Since(now).Seconds())
	}()
	for i := 0; i < 5; i++ {
		produceWithAck()
	}

}

// produceWithAck0 ...
func produceWithAck() {
	conf := readConfig()
	// conf["acks"] = 0
	// conf["acks"] = 1
	conf["acks"] = "all"

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		redLogBold.Println("create producer error", err)
	}

	var now = time.Now()
	defer func() {
		redLogBold.Printf("producer exist after %fs\n", time.Since(now).Seconds())
	}()

	deliveryCh := make(chan kafka.Event)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("ack 0"),
	}, deliveryCh)

	v := <-deliveryCh
	switch m := v.(type) {
	case *kafka.Message:
		if m.TopicPartition.Error != nil {
			redLogBold.Printf("Failed to deliver message: %v\n", m.TopicPartition)
		} else {
			cyanLog.Printf("topic:%s, partition:[%d], offset:%v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}

	producer.Close()
}
