package authentication

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewAuthRouter(handler *Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/token", handler.GenerateToken)
	return r
}
