package handler

import (
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/profile/logic"
	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type userHandler struct{}

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
	if err := validate(c, &user); err != nil {
		return
	}
	if err := userLogic.Create(user); err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		c.Done()
		return
	}
	c.JSON(http.StatusCreated, user)
}
