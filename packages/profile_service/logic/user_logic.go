package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository"
)

func GetAll() []model.UserProfile {
	//return cassandra.GetUserFirstnameById();
	//return cassandra.GetUser()
	return repository.UserProfile.All()
}
