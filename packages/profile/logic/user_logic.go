package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile/repository"
	"github.com/Luka-Spa/SwipeRight/packages/profile/util/producer"
)

var userRepository repository.IUserRepository

type IUserLogic interface {
	GetAll() []model.UserProfile
	Create(user model.UserProfile) error
}
type logic struct{}

func NewUserlogic(repository repository.IUserRepository) *logic {
	userRepository = repository
	return &logic{}
}

func (*logic) GetAll() []model.UserProfile {
	users, _ := userRepository.All()
	return users
}

func (*logic) Create(user model.UserProfile) error {
	user, err := userRepository.Create(user)
	if err == nil {
		producer.CreateUserProfile(user)
	}
	return err
}
