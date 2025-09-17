package service

import (
	configs "txrnxp-whats-happening/config"
	media "txrnxp-whats-happening/external/media/files"
	"txrnxp-whats-happening/internal/database"
	whatsHappening "txrnxp-whats-happening/internal/services/events"
	mediaService "txrnxp-whats-happening/internal/services/media"
)

type Bundle struct {
	WhatsHappeningService *whatsHappening.WhatsHappeningService
}

func NewBundle(
	repo database.Repository, config *configs.Config,
	media media.MediaStorageProvider, mediaService mediaService.MediaService,
) *Bundle {
	return &Bundle{
		WhatsHappeningService: whatsHappening.NewWhatsHappeningService(config, repo, mediaService),
	}
}
