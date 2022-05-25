package httpRouter

import (
	"fmt"
	"net/http"

	"github.com/Luka-Spa/SwipeRight/packages/recommendation/config"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/logic"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/repository/cassandra"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/router/http/handler"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/service/consumer"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() {
	config := config.GetConfig()
	if config.GetString("environment") == "production" {
		gin.SetMode("production")
	}
	engine := gin.Default()

	recommendationRepository := cassandra.NewRecommendationRepository()
	recomendationLogic := logic.NewUserlogic(recommendationRepository)
	recommendationHandler := handler.NewUserHandler(recomendationLogic)
	userProfileConsumer := consumer.NewKafkaConsumer(recomendationLogic)
	if config.GetBool("kafka.enabled") {
		userProfileConsumer.ConsumeUserProfile()
	}

	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	api := engine.Group("/api/")
	api.GET("/recommend/:user_id", recommendationHandler.Recommend)
	//Routes are defined here

	fmt.Println(engine.Run(fmt.Sprintf("%s:%s", config.GetString("host"), config.GetString("http.port"))))
}
