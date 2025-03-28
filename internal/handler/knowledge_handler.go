package handler

import (
	"kumulus/internal/auth"
	"kumulus/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KnowledgeHandler struct {
	Service *service.KnowledgeService
}

type KnowledgeRequest struct {
	Content   string  `json:"content"`
	RoleScope *string `json:"role_scope"` // opcional: "vendedor", "financeiro", etc
}

func (h *KnowledgeHandler) AddEntry(c *gin.Context) {
	claims := c.MustGet("auth").(*auth.Claims)

	var req KnowledgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	err := h.Service.AddEntry(claims.OrganizationID, claims.UserID, claims.Role, req.Content, req.RoleScope)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao salvar conteúdo"})
		return
	}

	c.Status(http.StatusCreated)
}
