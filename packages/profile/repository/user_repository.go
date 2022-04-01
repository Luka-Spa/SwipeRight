package repository

import "github.com/Luka-Spa/SwipeRight/packages/profile/model"


type IUserRepository interface {
	All() ([]model.UserProfile, error)
	Create(user model.UserProfile) error
}
