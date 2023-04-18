package demo

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ConsumerOffset ...
// Kafka brokers use an internal topic named __consumer_offsets that keeps track of what messages a given
// consumer group last successfully processed
//
// Delivery semantics for consumers
// By default, Java consumers automatically commit offsets (5s by default)
// A consumer may opt to commit offsets by itself (enable.auto.commit=false). Depending on when it chooses
// to commit offsets, there are delivery semantics available to the consumer. The three delivery semantics
// are explained below.
//
// - At most once
// Offset are committed as soon as the message is received
// If the processing goes wrong, the message will be lost (it won't be read again)
// - At least once (usually preferred)
// Offset are committed after the message is processed
// If the processing goes wrong, the message will be read again
// - Exactly once
// Achieved for Kafka topic using the transactions API.
//
func ConsumerOffset() {
	conf := readConfig()
	topic := "topic_0"
	conf["debug"] = "msg"
	conf["group.id"] = "test-consumer"
	conf["enable.auto.commit"] = false
	conf["auto.offset.reset"] = "earliest"
	// conf["auto.offset.reset"] = "latest"

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		redLogBold.Println("create consumer error", err)
	}
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		redLogBold.Println("subscribe consumer error", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case <-sigChan:
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				cyanLog.Printf("%% Message on %s:%s\n", e.TopicPartition, string(e.Value))
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}
	log.Println("close consumer")
	consumer.Close()
}
