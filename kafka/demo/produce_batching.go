package demo

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ProduceWithBatching ...
func ProduceWithBatching() {
	now := time.Now()
	defer func() {
		redLog.Printf("since: %fs", time.Since(now).Seconds())
	}()
	conf := readConfig()
	conf["debug"] = "msg"
	conf["batch.num.messages"] = 10
	conf["linger.ms"] = 1000
	// conf["batch.size"] = 10000
	// conf["delivery.timeout.ms"] = 1001
	// conf["request.timeout.ms"] = 1000

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		redLogBold.Println("create producer error", err)
	}

	go func() {
		defer func() {
			redLogBold.Println("producer exist")
		}()

		i := 0
		for e := range producer.Events() {
			i++
			switch m := e.(type) {
			case *kafka.Message:
				if m.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", m.TopicPartition)
				} else {
					cyanLog.Println(fmt.Sprintf("no_%d topic:%s, partition:[%d], offset:%v",
						i, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))
				}
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: 1,
			},
			Value: []byte("admicro's log"),
		}, nil)
	}

	producer.Flush(1000 * 10)
	producer.Close()
}
