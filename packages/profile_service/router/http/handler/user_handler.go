package handler

import (
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/logic"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var firstname = logic.GetAll()
	c.JSON(http.StatusOK, gin.H{"first_name":firstname})
}
