package repository

import "github.com/Luka-Spa/SwipeRight/packages/profile_service/model"


type IUserRepository interface {
	All() ([]model.UserProfile, error)
}
