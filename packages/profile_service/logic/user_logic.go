package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository/cassandra"
)

func GetUserFirstnameById() model.UserProfile {
	//return cassandra.GetUserFirstnameById();
	return cassandra.GetUser()
}
