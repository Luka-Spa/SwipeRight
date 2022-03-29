package cassandra

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
)

type UserRepository struct {

}

func NewUserRepository() *UserRepository {
	return &UserRepository{
	}
}

func (repo *UserRepository) All() []model.UserProfile {
	var query = "SELECT value FROM profile.user_profile_table WHERE pk=0;"
	var user model.UserProfile
	user = CassandraReadSingle(query, user)
	var users []model.UserProfile
	return append(users, user)
}
