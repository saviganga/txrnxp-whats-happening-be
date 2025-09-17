package service

import (
	external "txrnxp-whats-happening/external/events"
)

// EventService provides methods to interact with events.
type EventService struct {
	Scraper external.Scraper
}

// NewEventService creates a new EventService.
func NewEventService(s external.Scraper) *EventService {
	return &EventService{Scraper: s}
}

// GetEvents retrieves events from the specified URL.
func (s *EventService) GetEvents(url string) ([]external.Event, error) {
	return s.Scraper.ScrapeEvents(url)
}

// GetEvents retrieves events from the specified URL.
// func (s *EventService) GetEvents(url string) ([]external.Event, error) {
// 	return s.Scraper.ScrapeEvents(url)
// }
