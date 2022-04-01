package model

import "time"

type UserProfile struct {
	Id int `cql:"id" json:"id"`
	Email string `cql:"email" json:"email" binding:"required,email"`
	ProfileImageUrl string `cql:"profile_image_url" json:"profile_image_url" binding:"required,url"`
	Firstname string `cql:"first_name" json:"first_name" binding:"required"`
	Lastname string `cql:"last_name" json:"last_name" binding:"required"`
	Gender string `cql:"gender" json:"gender" binding:"required,oneof=f m"`
	Bio string `cql:"bio" json:"bio" binding:"required"`
	Location string `cql:"location" json:"location" binding:"required"`
	Anthem string `cql:"anthem" json:"anthem" binding:"required"`
	School string `cql:"school" json:"school" binding:"required"`
	CreatedAt time.Time `cql:"created_at" json:"created_at"`
	UpdatedAt time.Time `cql:"updated_at" json:"updated_at"`
}
