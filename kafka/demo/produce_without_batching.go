package demo

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ProduceWithoutBatching ...
func ProduceWithoutBatching() {
	conf := readConfig()
	// conf["debug"] = "msg"
	// conf["batch.num.messages"] = 5
	// conf["linger.ms"] = 5000

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Fatal("create producer error", err)
	}

	go func() {
		defer func() {
			log.Println("producer ack exist")
		}()

		i := 0
		for e := range producer.Events() {
			i++
			switch m := e.(type) {
			case *kafka.Message:
				if m.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", m.TopicPartition)
				} else {
					log.Printf("no_%d topic:%s, partition:[%d], offset:%v\n",
						i, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
			}
		}
	}()

	for i := 0; i < 10; i++ {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte("admicro's log"),
		}, nil)
	}

	producer.Flush(1000 * 10)
	producer.Close()
}
