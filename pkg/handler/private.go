package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Estimate(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) GetEstimate(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}
