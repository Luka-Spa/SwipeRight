package consumer

import (
	"context"
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/segmentio/kafka-go"
)

var brokers []string

func NewKafkaConsumer() IConsumer {
	var conf = config.GetConfig()
	brokers = conf.GetStringSlice("kafka.brokers")
	return &Consumer{}
}

func (*Consumer) ConsumeUserProfile() {
	var topic = "reservationapi.user.create"
	fmt.Printf("Consuming topic: %s \n", topic)
	go read(topic, func(m kafka.Message) {
		print(m.Value)
	})
}

func read(topic string, callback func(kafka.Message)) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:   topic,
		Brokers: brokers,
	})
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		callback(m)
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	defer reader.Close()
}
