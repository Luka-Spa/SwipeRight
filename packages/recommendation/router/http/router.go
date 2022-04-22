package httpRouter

import (
	"fmt"
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/logic"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/repository"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/router/http/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	if config.GetString("environment") == "production" {
		gin.SetMode("production")
	}
	engine := gin.Default()

	recommendationRepository := repository.RecommendationRepository
	reccomendationLogic := logic.NewUserlogic(recommendationRepository)
	reccomendationHandler := handler.NewUserHandler(reccomendationLogic)

	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())

	api := engine.Group("/api/")
	api.GET("/reccomend/:user_id", reccomendationHandler.Reccomend)
	//Routes are defined here
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	fmt.Println(engine.Run(fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("http.port"))))
}
