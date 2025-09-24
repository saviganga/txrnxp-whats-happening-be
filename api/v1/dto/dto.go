package dto

import (
	"time"

	"github.com/google/uuid"
)

// EventRequest is used when creating/updating an event
type EventRequest struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	EventType   string    `json:"event_type"`
	Country     string    `json:"country"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Category    string    `json:"category"`
	Duration    string    `json:"duration"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

// EventResponse is what gets returned to API clients
type EventResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	EventType   string    `json:"event_type"`
	Country     string    `json:"country"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Category    string    `json:"category"`
	Duration    string    `json:"duration"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UploadImageRequest struct {
	Image string `json:"image"`
}
