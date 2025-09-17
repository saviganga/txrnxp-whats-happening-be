package database

import (
	"fmt"
	"time"
	"txrnxp-whats-happening/api/v1/dto"
	"txrnxp-whats-happening/internal/database/tables"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) CreateEvent(event dto.EventRequest) (tables.WhatsHappening, error) {

	whatsHappening := tables.WhatsHappening{
		Name:        event.Name,
		Address:     event.Address,
		Description: event.Description,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Category:    event.Category,
		EventType:   event.EventType,
		Country:     event.Country,
	}

	err := r.db.Create(&whatsHappening).Error
	if err != nil {
		return whatsHappening, fmt.Errorf("unable to create event: %w", err)
	}

	return whatsHappening, nil

}

func (r *GormRepository) GetEvents() (events []tables.WhatsHappening, err error) {

	err = r.db.Model(&tables.WhatsHappening{}).Find(&events).Error
	if err != nil {
		return events, fmt.Errorf("unable to fetch events: %w", err)
	}

	return events, nil

}

func (r *GormRepository) GetWhatsHappeningEvents(page int) (PaginatedResponse, error) {
	const pageSize = 10
	var events []tables.WhatsHappening
	var totalCount int64

	// Ensure page starts at 1
	if page < 1 {
		page = 1
	}

	// Get today's midnight (local time)
	today := time.Now().Truncate(24 * time.Hour)

	// Count total rows starting from today
	if err := r.db.Model(&tables.WhatsHappening{}).
		Where("start_time >= ?", today).
		Count(&totalCount).Error; err != nil {
		return PaginatedResponse{}, fmt.Errorf("unable to count events: %w", err)
	}

	// Fetch paginated events starting from today
	if err := r.db.Model(&tables.WhatsHappening{}).
		Where("start_time >= ?", today).
		Order("start_time ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&events).Error; err != nil {
		return PaginatedResponse{}, fmt.Errorf("unable to fetch events: %w", err)
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize)) // ceil division

	return PaginatedResponse{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
		Data:       events,
	}, nil
}
