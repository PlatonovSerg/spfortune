package routes

import (
	"net/http"
	"time"

	"SparFortuneDDD/internal/modules/authentication"
	"SparFortuneDDD/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(jwtService *pkg.JWTService, authHandler *authentication.Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Compress(5, "gzip"))
	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/auth", authentication.NewAuthRouter(authHandler))
	})
	return r
}
