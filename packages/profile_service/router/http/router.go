package httpRouter

import (
	"SwipeRight_Profile_Service/config"
	"SwipeRight_Profile_Service/router/http/handler"
	"fmt"

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
