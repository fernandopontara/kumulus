package handler

import (
	"encoding/json"
	"kumulus/internal/auth"
	"kumulus/internal/repository"
	"net/http"
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.FindByEmail(req.Email)
	if err != nil || !auth.VerifyPassword(user.PasswordHash, req.Password) {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.OrganizationID, user.Role)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
