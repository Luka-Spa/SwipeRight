package repository

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository/cassandra"

	"github.com/gocql/gocql"
)

var DB *gocql.Session
var UserRepository IUserRepository

func Init() {
	config := config.GetConfig()
	cassandra.Connect(config)
	UserRepository = cassandra.NewUserRepository()
}


