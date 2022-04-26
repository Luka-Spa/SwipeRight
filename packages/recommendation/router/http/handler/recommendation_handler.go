package handler

import (
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/logic"
	"github.com/gin-gonic/gin"
)

type reccomendationHandler struct{}

var reccomendationLogic logic.IRecomendationLogic

func NewUserHandler(logic logic.IRecomendationLogic) *reccomendationHandler {
	reccomendationLogic = logic
	return &reccomendationHandler{}
}

func (*reccomendationHandler) Reccomend(c *gin.Context) {
}
