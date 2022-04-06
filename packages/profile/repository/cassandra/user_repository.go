package cassandra

import (
	"time"

	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/Luka-Spa/SwipeRight/packages/profile/util"
	"github.com/gocql/gocql"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) All() ([]model.UserProfile, error) {
	var query = "SELECT * FROM profile.user_profile_table;"
	var user model.UserProfile
	users := CassandraRead(query, user)
	return users, nil
}

func (repo *UserRepository) Create(user model.UserProfile) error {
	util.EncryptProps(&user, "xMmz2AWsH4Vmzcidgt2a043394qgnUrI")
	var query = `INSERT INTO profile.user_profile_table
							(id,email,profile_image_url,first_name,
							last_name,gender,bio,location,anthem,school,created_at,updated_at)
							values (?,?,?,?,?,?,?,?,?,?,?,?);`
	return CassandraWrite(query, gocql.TimeUUID(),
		user.Email, user.ProfileImageUrl,
		user.Firstname, user.Lastname,
		user.Gender, user.Bio,
		user.Location, user.Anthem,
		user.School, time.Now().UnixMilli(), time.Now().UnixMilli())
}
