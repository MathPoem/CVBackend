package handler

import (
	"CVBackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Estimate(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	var input models.Estimate
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Private.Estimate(input, userId)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetEstimate(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	estimates, err := h.service.Private.GetEstimate(userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, estimates)
}

func (h *Handler) CreatePerson(c *gin.Context) {
	userId, err := GetUserId(c)
	if err != nil {
		return
	}

	var input models.Person
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Private.CreatePerson(input, userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
