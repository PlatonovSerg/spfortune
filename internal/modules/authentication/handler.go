package authentication

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, `{"error": "user_id is required"}`, http.StatusBadRequest)
		return
	}

	token, err := h.service.GenerateToken(userID)
	if err != nil {
		http.Error(w, `{"error": "failed to generate token"}`, http.StatusBadRequest)
		return
	}

	response := TokenResponse{
		Token: token,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
		return
	}
}
