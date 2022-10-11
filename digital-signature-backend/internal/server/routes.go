package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mastrogiovanni/digital-signature-backend/internal/server/controllers"
)

// NewRouter creates all the routes/endpoints, using Fizz.
func NewRouter() (*gin.Engine, error) {

	engine := gin.New()

	engine.Use(cors.Default())

	engine.GET("/health", controllers.Healthcheck)

	engine.POST("/api/v1/sign", controllers.SignController)

	engine.POST("/api/v1/upload", controllers.UploadController)

	return engine, nil

}
