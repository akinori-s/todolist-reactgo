package handlers

import (
	"todolist/app/models"
	"todolist/app/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthUsecase usecases.AuthUsecase
}

func NewAuthHandler(authUsecase usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		AuthUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var user models.Auth
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.AuthUsecase.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
