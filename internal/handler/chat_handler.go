package handler

import (
	"kumulus/internal/auth"
	"kumulus/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	ChatService *service.ChatService
}

type ChatRequest struct {
	Question string `json:"question"`
}

type ChatResponse struct {
	Answer string `json:"answer"`
}

func (h *ChatHandler) Chat(c *gin.Context) {
	claims := c.MustGet("auth").(*auth.Claims)

	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	answer, err := h.ChatService.Perguntar(claims.OrganizationID, claims.UserID, claims.Role, req.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao consultar GPT"})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{Answer: answer})
}
