package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var brokers []string

func NewKafkaConsumer() IConsumer {
	var conf = config.GetConfig()
	brokers = conf.GetStringSlice("kafka.brokers")
	return &Consumer{}
}

func (*Consumer) ConsumeUserProfile() {
	var topic = "userapi.user.create"
	fmt.Printf("Consuming topic: %s \n", topic)
	go read(topic, func(m kafka.Message) {
		var user model.RecommendationProfile
		err := json.Unmarshal(m.Value, &user)
		if err != nil {
			log.Error(err)
		}
		log.Info(user)
	})
}

func read(topic string, callback func(kafka.Message)) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:     topic,
		Brokers:   brokers,
		Partition: 0,
	})
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		callback(m)
		log.Infof("Message from topic %s with key %s", topic, m.Key)
	}
	defer reader.Close()
}
