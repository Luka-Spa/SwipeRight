package model

type Location struct {
	City string  `cql:"city" json:"city" binding:"required" encrypt:""`
	Lat  float32 `cql:"lat" json:"lat" binding:"required,min=-90,max=90"`
	Long float32 `cql:"long" json:"long" binding:"required,min=-180,max=180"`
}
