package httpRouter

import (
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/logic"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/repository"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/router/http/handler"

	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	engine := gin.Default()

	userRepository := repository.UserRepository
	userLogic := logic.NewUserlogic(userRepository)
	userHandler := handler.NewUserHandler(userLogic)

	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())

	api := engine.Group("/api/")

	//Routes are defined here
	api.GET("/user", userHandler.GetAllUsers)
	api.POST("/user", userHandler.CreateUser)

	fmt.Println(engine.Run(fmt.Sprintf(":%s", config.GetString("http.port"))))
}
