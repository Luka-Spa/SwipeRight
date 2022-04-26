package model

import "github.com/gocql/gocql"

type RecommendationProfile struct {
	Id              gocql.UUID      `cql:"id" json:"id"`
	ProfileImageUrl string          `cql:"profile_image_url" mapstructure:"profile_image_url" json:"profile_image_url" binding:"required,url"`
	Firstname       string          `cql:"first_name" json:"first_name" mapstructure:"first_name" binding:"required"`
	Gender          string          `cql:"gender" json:"gender" binding:"required,oneof=f m" encrypt:""`
	Bio             string          `cql:"bio" json:"bio" binding:"required"`
	Location        Location        `cql:"location" json:"location" binding:"required"`
	Anthem          string          `cql:"anthem" json:"anthem" binding:"required"`
	School          string          `cql:"school" json:"school" binding:"required" encrypt:""`
	Preferences     UserPreferences `cql:"preferences" mapstructure:"preferences" json:"preferences" encrypt:""`
}
