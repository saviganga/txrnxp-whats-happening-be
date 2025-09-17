package service

import (
	"errors"
	"fmt"
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
		fmt.Println(err)
		return events, errors.New("unable to fetch what's happening events")
	}
	return events, nil

}

func (s *WhatsHappeningService) CreateEvents(input dto.EventRequest) (tables.WhatsHappening, error) {

	event, err := s.repo.CreateEvent(input)
	if err != nil {
		fmt.Println(err)
		return event, errors.New("unable to create what's happening events")
	}
	return event, nil

}
