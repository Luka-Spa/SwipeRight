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

func (repo *UserRepository) Create(user model.UserProfile) (model.UserProfile, error) {
	user.Id = gocql.TimeUUID()
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
	uEnc := user
	util.EncryptProps(&uEnc, "xMmz2AWsH4Vmzcidgt2a043394qgnUrI")
	var query = `INSERT INTO profile.user_profile_table
							(id,email,profile_image_url,first_name,
							last_name,gender,bio,location,anthem,school,preferences,created_at,updated_at)
							values (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	err := CassandraWrite(query, user.Id,
		uEnc.Email, uEnc.ProfileImageUrl,
		uEnc.Firstname, uEnc.Lastname,
		uEnc.Gender, uEnc.Bio,
		uEnc.Location, uEnc.Anthem,
		uEnc.School, uEnc.Preferences, uEnc.CreatedAt, uEnc.UpdatedAt)
	return user, err
}
