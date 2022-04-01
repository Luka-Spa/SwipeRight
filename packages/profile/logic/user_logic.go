package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile/repository"
)

var userRepository repository.IUserRepository

type IUserLogic interface {
	 GetAll() []model.UserProfile
	 Create(user model.UserProfile)
}
type logic struct {}

func NewUserlogic(repository repository.IUserRepository) *logic {
	userRepository = repository
	return &logic{}
}

func (*logic) GetAll() []model.UserProfile {
	users,_ := userRepository.All()
	return users
}

func (*logic) Create(user model.UserProfile) {
	userRepository.Create(user)
}
