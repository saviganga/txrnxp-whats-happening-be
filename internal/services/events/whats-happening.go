package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"txrnxp-whats-happening/api/v1/dto"
	configs "txrnxp-whats-happening/config"
	"txrnxp-whats-happening/internal/database"
	"txrnxp-whats-happening/internal/database/tables"
	mediaService "txrnxp-whats-happening/internal/services/media"
)

type WhatsHappeningService struct {
	env          *configs.Config
	repo         database.Repository
	mediaService mediaService.MediaService
}

func NewWhatsHappeningService(
	env *configs.Config,
	repo database.Repository,
	mediaService mediaService.MediaService,
) *WhatsHappeningService {
	return &WhatsHappeningService{
		env:          env,
		repo:         repo,
		mediaService: mediaService,
	}
}

func (s *WhatsHappeningService) GetEvents(page int) (database.PaginatedResponse, error) {

	events, err := s.repo.GetWhatsHappeningEvents(page)
	if err != nil {
		return events, errors.New("unable to fetch what's happening events")
	}
	eventss := []tables.WhatsHappening{}
	for _, event := range events.Data {

		var eventImage string
		if event.Image != "" {
			eventImage = s.mediaService.GetMediaURL(event.Image)
		} else {
			eventImage = eventImage
		}

		eventt := tables.WhatsHappening{
			ID:          event.ID,
			Name:        event.Name,
			Image:       eventImage,
			EventType:   event.EventType,
			Country:     event.Country,
			Description: event.Description,
			Address:     event.Address,
			Category:    event.Category,
			Duration:    event.Duration,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
		}
		eventss = append(eventss, eventt)

	}
	events.Data = eventss
	return events, nil

}

func (s *WhatsHappeningService) CreateEvents(input dto.EventRequest) (tables.WhatsHappening, error) {

	event, err := s.repo.CreateEvent(input)
	if err != nil {
		return event, errors.New("unable to create what's happening events")
	}
	return event, nil

}

func (s *WhatsHappeningService) UploadEventImage(eventID string, input dto.UploadImageRequest) error {

	// validate the base64 encoding
	if !strings.Contains(input.Image, "data:image") {
		return errors.New("invalid image format")
	}

	// get image data from base 64 string
	parts := strings.Split(input.Image, ",")
	if len(parts) < 2 {
		return errors.New("invalid image data")
	}

	// decode base 64 image
	imageData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return errors.New("unable to decode image")
	}

	event, err := s.repo.GetWhatsHappeningEvent(eventID)
	if err != nil {
		return errors.New("unable to fetch event")
	}

	// name the file
	fileName := event.ID.String() + ".png"

	imageURL, err := s.mediaService.UploadMedia(fileName, imageData)
	if err != nil {
		return errors.New("unable to upload image")
	}

	// upload user image
	err = s.repo.UploadEventImage(event, imageURL)
	if err != nil {
		return errors.New("unable to save image")
	}

	return nil

}
