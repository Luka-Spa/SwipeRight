package model

type UserPreferences struct {
	Visible          bool   `cql:"visible" json:"visible" binding:"required"`
	AgeMin           int    `cql:"age_min" json:"age_min" binding:"required"`
	AgeMax           int    `cql:"age_max" json:"age_max" binding:"required"`
	GenderPreference string `cql:"gender_preference" json:"gender_preference" binding:"required" encrypt:""`
	MaxDistance      int    `cql:"max_distance" json:"max_distance" binding:"required"`
}
