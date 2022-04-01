package handler

import (
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/logic"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/model"
	"github.com/gin-gonic/gin"
)

type userHandler struct {}

var userLogic logic.IUserLogic

func NewUserHandler(logic logic.IUserLogic) *userHandler {
	userLogic = logic
	return &userHandler{}
}

func (handler *userHandler) GetAllUsers(c *gin.Context) {
	var users = userLogic.GetAll()
	c.JSON(http.StatusOK, users)
}

func (handler *userHandler) CreateUser(c *gin.Context) {
	var user model.UserProfile
	if err :=validate(c, &user); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}
