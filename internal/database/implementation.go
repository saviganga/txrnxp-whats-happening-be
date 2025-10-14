package database

import (
	"fmt"
	"time"
	"txrnxp-whats-happening/api/v1/dto"
	"txrnxp-whats-happening/internal/database/tables"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) CreateEvent(event dto.EventRequest) (tables.WhatsHappening, error) {

	id, _ := uuid.Parse(event.ID)
	whatsHappening := tables.WhatsHappening{
		ID:          id,
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

// func (r *GormRepository) GetWhatsHappeningEvents(page int) (PaginatedResponse, error) {
// 	const pageSize = 10
// 	var events []tables.WhatsHappening
// 	var totalCount int64

// 	// Ensure page starts at 1
// 	if page < 1 {
// 		page = 1
// 	}

// 	// Get today's midnight (local time)
// 	today := time.Now().Truncate(24 * time.Hour)

// 	// Count total rows starting from today
// 	if err := r.db.Model(&tables.WhatsHappening{}).
// 		Where("start_time >= ?", today).
// 		Count(&totalCount).Error; err != nil {
// 		return PaginatedResponse{}, fmt.Errorf("unable to count events: %w", err)
// 	}

// 	// Fetch paginated events starting from today
// 	if err := r.db.Model(&tables.WhatsHappening{}).
// 		Where("start_time >= ?", today).
// 		Order("start_time ASC").
// 		Limit(pageSize).
// 		Offset((page - 1) * pageSize).
// 		Find(&events).Error; err != nil {
// 		return PaginatedResponse{}, fmt.Errorf("unable to fetch events: %w", err)
// 	}

// 	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize)) // ceil division

// 	return PaginatedResponse{
// 		Page:       page,
// 		PageSize:   pageSize,
// 		TotalCount: totalCount,
// 		TotalPages: totalPages,
// 		Data:       events,
// 	}, nil
// }

func (r *GormRepository) GetWhatsHappeningEvents(page int, filters map[string]string) (PaginatedResponse, error) {
	const pageSize = 10
	var events []tables.WhatsHappening
	var totalCount int64
	const DateLayout = "02-01-2006" // dd-mm-yyyy

	// Ensure page starts at 1
	if page < 1 {
		page = 1
	}

	// Get today's midnight (local time)
	today := time.Now().Truncate(24 * time.Hour)

	// Base query: events starting from today
	query := r.db.Model(&tables.WhatsHappening{}).Where("start_time >= ?", today)

	// Apply filters dynamically
	for key, value := range filters {
		switch key {
		case "name":
			query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+value+"%")
		case "address":
			query = query.Where("LOWER(address) LIKE LOWER(?)", "%"+value+"%")
		case "category":
			query = query.Where("LOWER(category) LIKE LOWER(?)", "%"+value+"%")
			// query = query.Where("category = ?", value)
		// case "start_time":
		// 	if t, err := time.Parse(time.RFC3339, value); err == nil {
		// 		query = query.Where("start_time >= ?", t)
		// 	}
		// case "end_time":
		// 	if t, err := time.Parse(time.RFC3339, value); err == nil {
		// 		query = query.Where("end_time <= ?", t)
		// 	}
		case "start_time":
			if t, err := time.Parse(DateLayout, value); err == nil {
				query = query.Where("start_time >= ?", t)
			}
		case "end_time":
			if t, err := time.Parse(DateLayout, value); err == nil {
				query = query.Where("end_time <= ?", t)
			}
		}
	}

	// Count total rows
	if err := query.Count(&totalCount).Error; err != nil {
		return PaginatedResponse{}, fmt.Errorf("unable to count events: %w", err)
	}

	// Fetch paginated events
	if err := query.Order("start_time ASC").
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

func (r *GormRepository) GetWhatsHappeningEvent(eventID string) (tables.WhatsHappening, error) {

	var event tables.WhatsHappening

	err := r.db.Model(&tables.WhatsHappening{}).
		Where("id = ?", eventID).
		First(&event).Error
	if err != nil {
		return event, fmt.Errorf("unable to fetch events: %w", err)
	}

	return event, nil

}

func (r *GormRepository) UploadEventImage(event tables.WhatsHappening, imageURL string) error {
	err := r.db.Debug().Model(&tables.WhatsHappening{}).Where("id = ?", event.ID).Update("image", imageURL).Error
	if err != nil {
		return fmt.Errorf("failed to update event image: %w", err)
	}

	return nil
}
