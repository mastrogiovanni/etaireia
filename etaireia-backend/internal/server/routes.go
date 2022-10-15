package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mastrogiovanni/etaireia/etaireia-backend/internal/server/controllers"
)

// NewRouter creates all the routes/endpoints, using Fizz.
func NewRouter() (*gin.Engine, error) {

	engine := gin.New()

	engine.Use(cors.Default())

	engine.GET("/health", controllers.Healthcheck)

	engine.POST("/api/v1/sign", controllers.SignController)
	engine.GET("/api/v1/sign/:publicKey", controllers.ListSigned)

	engine.POST("/api/v1/subscription", controllers.CreateSubscription)
	engine.GET("/api/v1/subscription", controllers.ListSubscription)
	engine.POST("/api/v1/subscription/approve", controllers.SubscriptionApprove)

	return engine, nil

}
