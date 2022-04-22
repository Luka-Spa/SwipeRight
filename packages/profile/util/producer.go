package util

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile/util/kafka"
	log "github.com/sirupsen/logrus"
)

type IProducer interface {
	CreateUserProfile(user model.UserProfile)
}

var _producers []IProducer

func InitProducers() {
	config := config.GetConfig()
	if config.GetBool("kafka.enabled") {
		log.Info("Initializing with a kafka producer")
		_producers = append(_producers, kafka.NewProducer(config))
	}
}

func CreateReservation(user model.UserProfile) {
	for _, producer := range _producers {
		go producer.CreateUserProfile(user)
	}
}
