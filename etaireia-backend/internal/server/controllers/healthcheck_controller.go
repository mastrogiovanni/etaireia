package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mastrogiovanni/etaireia/etaireia-backend/internal/server/models"
)

func Healthcheck(c *gin.Context) {
	models.ResponseSuccess(c)
}
