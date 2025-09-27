package v1

import (
	"fmt"
	"net/http"
	"os"

	"txrnxp-whats-happening/api/v1/handlers"
	service "txrnxp-whats-happening/internal/services"

	"github.com/go-chi/chi/v5"
)

// Routes sets up the API routes for events using chi.Router.
func Routes(router chi.Router, bundle service.Bundle) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "v1" // fallback default
	}
	fmt.Printf("\n\nversion: %s\n\n", version)

	pathPrefix := fmt.Sprintf("/events")

	// Mount all event routes under /api/{version}/events
	router.Route(pathPrefix, func(r chi.Router) {
		// GET /api/{version}/events/whats-happening
		r.Get("/whats-happening", func(w http.ResponseWriter, req *http.Request) {
			handlers.GetWhatsHappeningEvents(w, req, *bundle.WhatsHappeningService)
		})

		// GET /api/{version}/events/whats-happening/{event-id}
		r.Get("/whats-happening/{event-id}", func(w http.ResponseWriter, req *http.Request) {
			handlers.GetWhatsHappeningEvent(w, req, *bundle.WhatsHappeningService)
		})

		// POST /api/{version}/events
		r.Post("/", func(w http.ResponseWriter, req *http.Request) {
			handlers.CreateEvents(w, req, *bundle.WhatsHappeningService)
		})

		r.Post("/{event-id}/upload-image", func(w http.ResponseWriter, req *http.Request) {
			handlers.UploadEventImage(w, req, *bundle.WhatsHappeningService)
		})

		// GET /api/{version}/events/health
		r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Events API is healthy"))
		})
	})
}
