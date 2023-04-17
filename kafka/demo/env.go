package demo

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	key    = "ZJK3NO7AKWMXL6XN"
	secret = "FiPtTlSB2rUa1WVW9lMvGTgG2KnBJX0QlzvgLEPA4uOM+CiYNsnrKhi/ZKvhNGvr"
)

var topic = "topic_0"

func readConfig() kafka.ConfigMap {
	m := make(map[string]kafka.ConfigValue)
	m["bootstrap.servers"] = "pkc-ldvr1.asia-southeast1.gcp.confluent.cloud"
	m["security.protocol"] = "SASL_SSL"
	m["sasl.mechanisms"] = "PLAIN"
	m["sasl.username"] = key
	m["sasl.password"] = secret

	return m
}

func init() {
	log.SetFlags(0)
}
