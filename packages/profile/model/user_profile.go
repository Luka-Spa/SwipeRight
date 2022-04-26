package model

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	Id              uuid.UUID       `cql:"id" json:"id"`
	Email           string          `cql:"email" json:"email" binding:"required,email" encrypt:""`
	ProfileImageUrl string          `cql:"profile_image_url" mapstructure:"profile_image_url" json:"profile_image_url" binding:"required,url"`
	Firstname       string          `cql:"first_name" json:"first_name" mapstructure:"first_name" binding:"required"`
	Lastname        string          `cql:"last_name" json:"last_name" mapstructure:"last_name" binding:"required" encrypt:""`
	Gender          string          `cql:"gender" json:"gender" binding:"required,oneof=f m" encrypt:""`
	Bio             string          `cql:"bio" json:"bio" binding:"required"`
	Location        Location        `cql:"location" json:"location" binding:"required"`
	Anthem          string          `cql:"anthem" json:"anthem" binding:"required"`
	School          string          `cql:"school" json:"school" binding:"required" encrypt:""`
	CreatedAt       time.Time       `cql:"created_at" mapstructure:"created_at" json:"created_at"`
	UpdatedAt       time.Time       `cql:"updated_at" mapstructure:"updated_at" json:"updated_at"`
	Preferences     UserPreferences `cql:"preferences" mapstructure:"preferences" json:"preferences" encrypt:""`
}
