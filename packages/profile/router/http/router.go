package httpRouter

import (
	"fmt"
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/profile/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile/logic"
	"github.com/Luka-Spa/SwipeRight/packages/profile/repository"
	"github.com/Luka-Spa/SwipeRight/packages/profile/router/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	config := config.GetConfig()
	if config.GetString("environment") == "production" {
		gin.SetMode("production")
	}
	engine := gin.Default()

	userRepository := repository.UserRepository
	userLogic := logic.NewUserlogic(userRepository)
	userHandler := handler.NewUserHandler(userLogic)

	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := engine.Group("/api/")

	//Routes are defined here
	api.GET("/user", func(c *gin.Context) { handler.Auth(c, []string{"read:users"}) }, userHandler.GetAllUsers)
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	api.POST("/user", func(c *gin.Context) { handler.Auth(c, []string{"create:users"}) }, userHandler.CreateUser)

	fmt.Println(engine.Run(fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("http.port"))))
}
