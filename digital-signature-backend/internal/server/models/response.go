package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DSResponse struct {
	Success bool        `json:"success"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, DSResponse{Success: true})
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, DSResponse{Success: true, Data: data})
}

func ResponseFail(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, DSResponse{Success: false, Message: &message})
}

func ResponseError(c *gin.Context, err error) {
	msg := err.Error()
	c.JSON(http.StatusBadRequest, DSResponse{Success: false, Message: &msg})
}
