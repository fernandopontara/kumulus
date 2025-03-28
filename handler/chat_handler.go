package handler

import (
	"encoding/json"
	"kumulus/internal/auth"
	"kumulus/internal/service"
	"net/http"
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

func (h *ChatHandler) Chat(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("auth").(*auth.Claims)

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	answer, err := h.ChatService.Perguntar(claims.OrganizationID, claims.UserID, claims.Role, req.Question)
	if err != nil {
		http.Error(w, "Erro ao processar pergunta", http.StatusInternalServerError)
		return
	}

	resp := ChatResponse{Answer: answer}
	json.NewEncoder(w).Encode(resp)
}
