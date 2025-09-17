package database

import (
	"txrnxp-whats-happening/api/v1/dto"
	"txrnxp-whats-happening/internal/database/tables"
)

type PaginatedResponse struct {
	Page       int                     `json:"page"`
	PageSize   int                     `json:"page_size"`
	TotalCount int64                   `json:"total_count"`
	TotalPages int                     `json:"total_pages"`
	Data       []tables.WhatsHappening `json:"data"`
}

type Repository interface {
	CreateEvent(event dto.EventRequest) (tables.WhatsHappening, error)
	GetWhatsHappeningEvents(page int) (PaginatedResponse, error)
}
