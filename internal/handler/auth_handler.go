package handler

import (
	"kumulus/internal/auth"
	"kumulus/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	user, err := h.UserRepo.FindByEmail(req.Email)
	if err != nil || !auth.VerifyPassword(user.PasswordHash, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.OrganizationID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
