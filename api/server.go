package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	v1 "txrnxp-whats-happening/api/v1"
	service "txrnxp-whats-happening/internal/services"
)

// NewServer creates a new server instance with chi.
func NewServer(bundle service.Bundle) http.Handler {
	mux := chi.NewRouter()

	// --- Middleware ---
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger) // simple request logging
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // tie to cfg.AllowedOrigins if needed
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
		},
		AllowedHeaders: allowedHeaders,
	}))

	// --- Swagger docs ---
	basePath := "/api/v1"

	// --- Versioned routes ---
	mux.Route(basePath, func(r chi.Router) {
		v1.Routes(r, bundle)
	})

	// --- Root ping ---
	mux.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Server is live"))
	})

	// --- Not found handler ---
	mux.NotFound(routeNotFoundHandler)

	return mux
}

// Not found
func routeNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("Content-Type") != "" {
		return
	}
	http.Error(w, "no route found for this path", http.StatusNotFound)
}

var allowedHeaders = []string{
	"Authorization",
	"Content-Type",
	"Accept",
	"Origin",
	"Referer",
	"User-Agent",
	"X-Request-ID",
}
