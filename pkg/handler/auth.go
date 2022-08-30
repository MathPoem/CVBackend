package handler

import (
	"CVBackend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.Authorization.GetUserByName(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.SetCookie("refresh", tokens.RefreshToken, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":      tokens.AccessToken,
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"university": user.Univ,
		"program":    user.Program,
		"confirmed":  user.Confirmed,
	},
	)
}

func (h *Handler) Refresh(c *gin.Context) {

	cookie, err := c.Cookie("refresh")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Authorization.Refresh(cookie)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "error refresh")
	}

	c.SetCookie("refresh", tokens.RefreshToken, 24*10000, "/", "localhost", false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"AccessToken": tokens.AccessToken,
	})
}

type LogoutInput struct {
	Id    int    `json:"id" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (h *Handler) Logout(c *gin.Context) {
	var input LogoutInput
	err := c.BindJSON(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Authorization.Logout(input.Id, input.Email); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.SetCookie("refresh", "", 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "logout ok",
	})
}

type ActivateInput struct {
	Key string `json:"key" binding:"required"`
}

func (h *Handler) Activate(c *gin.Context) {
	var input ActivateInput

	err := c.BindJSON(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusOK, err.Error())
		return
	}

	err = h.service.Authorization.Activate(input.Key)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "activate success",
	})
}
