package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	MsgChanel chan **ckafka.Message
}

func (l *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"boostrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":         os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message")
	}
}
