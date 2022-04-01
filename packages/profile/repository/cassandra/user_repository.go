package cassandra

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/google/uuid"
)

type UserRepository struct {

}

func NewUserRepository() *UserRepository {
	return &UserRepository{
	}
}

func (repo *UserRepository) All() ([]model.UserProfile, error) {
	var query = "SELECT value FROM profile.user_profile_table;"
	var user model.UserProfile
	users := CassandraRead(query, user)
	return users, nil
}

func (repo *UserRepository) Create(user model.UserProfile) error {
	var query = "INSERT INTO profile.user_profile_table (pk, value) values (?,?);"
	return CassandraWrite(query,uuid.New() , user)
}
