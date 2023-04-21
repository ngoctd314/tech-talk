package demo

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	conf["group.id"] = "test-consumers12"
	conf["enable.auto.commit"] = false
	conf["auto.offset.reset"] = "earliest"
	// conf["max.poll.records"] = 1000
	conf["fetch.max.bytes"] = 5000
	conf["message.max.bytes"] = 1000
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
	cnt := 0
	now := time.Now()
	for run {
		select {
		case <-sigChan:
			run = false
		default:
			now := time.Now()
			ev := consumer.Poll(5000)
			if cnt == 0 {
				redLog.Printf("Time %fs\n", time.Since(now).Seconds())
			} else {

				cyanLog.Printf("Time %fs\n", time.Since(now).Seconds())
			}
			time.Sleep(time.Second)
			if ev == nil {
				cyanLog.Println("ev = nil")
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				cnt++
				cyanLog.Printf("%% Message on %s:%s cnt: %d since:%fs\n", e.TopicPartition, string(e.Value), cnt, time.Since(now).Seconds())
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v, cnt: %d\n", e.Code(), e, cnt)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}
	log.Printf("close consumer %fs\n", time.Since(now).Seconds())
	consumer.Close()
}
