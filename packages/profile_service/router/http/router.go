package httpRouter

import (
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/router/http/handler"

	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	engine := gin.Default()

	engine.Use(gin.Recovery())

	apiv1 := engine.Group("/api/")
	apiv1.GET("/user", handler.GetAllUsers)

	fmt.Println(engine.Run(fmt.Sprintf(":%s", config.GetString("http.port"))))
}
