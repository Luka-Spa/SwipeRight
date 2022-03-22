package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	//var users []model.UserProfile
	//l := repository.DB.Query("SELECT * FROM USER_PROFILE")
	//fmt.Println(l)
	c.JSON(http.StatusOK, gin.H{"data": "cool"})
}
