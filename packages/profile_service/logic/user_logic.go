package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository/cassandra"
)

func GetUserFirstnameById() string {
	return cassandra.GetUserFirstnameById();
}
