package httpRouter

import (
<<<<<<< HEAD
	"SwipeRight_Profile_Service/config"
	"SwipeRight_Profile_Service/router/http/handler"
	"fmt"

=======
	"fmt"

	"github.com/Luka-Spa/SwipeRight/packages/profile_service/config"
	"github.com/Luka-Spa/SwipeRight/packages/profile_service/router/http/handler"

>>>>>>> 18c1201 (new commit)
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
