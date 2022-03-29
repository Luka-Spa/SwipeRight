package repository

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository/cassandra"

	"github.com/gocql/gocql"
)

var DB *gocql.Session
var UserProfile IUserRepository

func Init() {
	config := config.GetConfig()
	cassandra.Connect(config)
	UserProfile = cassandra.NewUserRepository()
}


