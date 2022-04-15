package consumer

import (
	"context"
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/segmentio/kafka-go"
)

var reader *kafka.Reader

func NewKafkaConsumer() IConsumer {
	config := config.GetConfig()
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: config.GetStringSlice("kafka.brokers"),
		Topic:   config.GetString("kafka.topic"),
	})
	return &Consumer{}
}

func (*Consumer) Run() {
	fmt.Printf("Consuming topic: %s \n", reader.Config().Topic)
	go read()
}

func read() {
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	defer reader.Close()
}
