package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// service "txrnxp-whats-happening/internal/services/events"
	"txrnxp-whats-happening/api/v1/dto"
	services "txrnxp-whats-happening/internal/services/events"

	"github.com/go-chi/chi/v5"
)

// GetEvents handles GET requests to fetch events.
// func GetEvents(w http.ResponseWriter, r *http.Request, repo database.Repository) {
// 	log.Println("Received GET /events request")

// 	events, err := service.GetEvents(repo) // call service function with repo
// 	if err != nil {
// 		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(events)
// }

// GetWhatsHappeningEvents handles GET requests to fetch "whats happening" events.
func GetWhatsHappeningEvents(w http.ResponseWriter, r *http.Request, whatsHappening services.WhatsHappeningService) {
	log.Println("Received GET /events/whats-happening request")

	// Get ?page query param (default = 1)
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	// Fetch events with pagination
	events, err := whatsHappening.GetEvents(page)
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// CreateEvents handles POST requests to create a new event.
func CreateEvents(w http.ResponseWriter, r *http.Request, whatsHappening services.WhatsHappeningService) {
	log.Println("Received POST /events request")

	var input dto.EventRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	event, err := whatsHappening.CreateEvents(input)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// CreateEvents handles POST requests to create a new event.
func UploadEventImage(w http.ResponseWriter, r *http.Request, whatsHappening services.WhatsHappeningService) {

	id := chi.URLParam(r, "event-id")
	if id == "" {
		http.Error(w, "Missing event ID in URL", http.StatusBadRequest)
		return
	}

	var input dto.UploadImageRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := whatsHappening.UploadEventImage(id, input)
	if err != nil {
		http.Error(w, "Failed to upload event image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nil)
}
