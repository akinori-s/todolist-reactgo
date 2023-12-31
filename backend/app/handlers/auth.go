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
	jwtToken, dbUser, err := h.AuthUsecase.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token", jwtToken, (60*60*24 * 7), "/", "localhost", false, true) 
	c.JSON(http.StatusOK, dbUser)
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var user models.Auth
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwtToken, newUser, err := h.AuthUsecase.Signup(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("token", jwtToken, (60*60*24 * 7), "/", "localhost", false, true) 
	c.JSON(http.StatusOK, newUser)
}

func (h *AuthHandler) CheckLogin(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user, err := h.AuthUsecase.CheckLogin(cookie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Signout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true) 
	c.JSON(http.StatusOK, gin.H{"message": "Signed out successfully."})
}
