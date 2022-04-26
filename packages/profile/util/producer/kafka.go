package producer

import (
	"context"
	"encoding/json"

	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var writer *kafka.Writer
var topicPrefix string

type KafkaProducer struct{}

func NewProducer(config *viper.Viper) *KafkaProducer {
	//mechanism, err := scram.Mechanism(scram.SHA256, config.GetString("kafka.username"), config.GetString("kafka.password"))
	//if err != nil {
	//	log.Error(err)
	//}
	// writer = kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers:     config.GetStringSlice("kafka.brokers"),
	// 	Logger:      kafka.LoggerFunc(log.Infof),
	// 	ErrorLogger: kafka.LoggerFunc(log.Errorf),
	// 	Dialer: &kafka.Dialer{
	// 		Timeout:   10 * time.Second,
	// 		DualStack: true,
	// 		TLS:       &tls.Config{},
	// 		//SASLMechanism: mechanism,
	// 	},
	// })
	writer = &kafka.Writer{
		Addr:     kafka.TCP(config.GetStringSlice("kafka.brokers")...),
		Balancer: &kafka.LeastBytes{},
	}
	topicPrefix = config.GetString("kafka.topic_prefix")
	return &KafkaProducer{}
}

func (*KafkaProducer) CreateUserProfile(user model.UserProfile) {
	reservationJson, _ := json.Marshal(user)
	topic := topicPrefix + "userapi.user.create"
	log.Info("Producing user")
	err := writer.WriteMessages(context.TODO(), kafka.Message{
		Topic: topic,
		Key:   []byte(user.Id.String()),
		Value: reservationJson,
	})
	if err != nil {
		log.Error(err)
	}
}
