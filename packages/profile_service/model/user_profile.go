package model

import "time"

type UserProfile struct {
	Id int32 `cql:"id"`
	Email string `cql:"email"`
	ProfileImageUrl string `cql:"profile_image_url"`
	Firstname string `cql:"first_name"`
	Lastname string `cql:"last_name"`
	Gender string `cql:"gender"`
	Bio string `cql:"bio"`
	Location string `cql:"location"`
	Anthem string `cql:"anthem"`
	School string `cql:"school"`
	CreatedAt time.Time `cql:"created_at"`
	UpdatedAt time.Time `cql:"updated_at"`
}
