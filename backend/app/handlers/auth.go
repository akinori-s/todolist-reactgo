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

func (h *AuthHandler) Signup(c *gin.Context) {
	var user models.Auth
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtToken, err := h.AuthUsecase.Signup(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token", jwtToken, (60*60*24 * 7), "/", "localhost", false, true) 
	c.JSON(http.StatusOK, "SUCCESS")
}
