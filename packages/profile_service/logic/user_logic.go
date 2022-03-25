package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository/cassandra"
)

func GetUserFirstnameById() map[string]interface{} {
	//return cassandra.GetUserFirstnameById();
	return cassandra.GetUser()
}
