package repository

import (
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/repository/cassandra"
	"github.com/gocql/gocql"
)

var DB *gocql.Session
var RecommendationRepository IRecommendationRepository

func Init() {
	config := config.GetConfig()
	cassandra.Connect(config)
	RecommendationRepository = cassandra.NewRecommendationRepository()
}
