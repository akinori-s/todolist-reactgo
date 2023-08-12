package handlers

import (
	"todolist/app/models"
	"todolist/app/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthRepository repositories.AuthRepository
}

func NewAuthHandler(authRepository repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{
		AuthRepository: authRepository,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.AuthRepository.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}
