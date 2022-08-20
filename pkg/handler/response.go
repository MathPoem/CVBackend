package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}