package handler

import (
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/logic"
	"github.com/gin-gonic/gin"
)

type recommendationHandler struct{}

// var recommendationLogic logic.IRecomendationLogic

func NewUserHandler(logic logic.IRecomendationLogic) *recommendationHandler {
	// recommendationLogic = logic
	return &recommendationHandler{}
}

func (*recommendationHandler) Recommend(c *gin.Context) {
}
