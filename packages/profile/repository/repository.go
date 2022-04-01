package repository

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile/repository/cassandra"

	"github.com/gocql/gocql"
)

var DB *gocql.Session
var UserRepository IUserRepository

func Init() {
	config := config.GetConfig()
	cassandra.Connect(config)
	UserRepository = cassandra.NewUserRepository()
}


