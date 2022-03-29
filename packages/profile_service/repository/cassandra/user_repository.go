package cassandra

import (
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository"
)

func GetUser() model.UserProfile {
	var query = "SELECT value FROM profile.user_profile_table WHERE pk=0;"
	var user model.UserProfile
	user = repository.CassandraReadSingle(query, user)
	return user
}
