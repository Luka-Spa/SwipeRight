package cassandra

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/google/uuid"
)

type UserRepository struct {

}

func NewUserRepository() *UserRepository {
	return &UserRepository{
	}
}

func (repo *UserRepository) All() ([]model.UserProfile, error) {
	var query = "SELECT value FROM profile.user_profile_table WHERE pk=0;"
	var user model.UserProfile
	user = CassandraReadSingle(query, user)
	var users []model.UserProfile
	return append(users, user), nil
}

func (repo *UserRepository) Create(user model.UserProfile) {
	var query = "INSERT INTO profile.user_profile_table (pk, value)"
	CassandraWrite(query,uuid.New() , user)
}
