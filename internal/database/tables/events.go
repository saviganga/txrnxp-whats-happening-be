package tables

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WhatsHappening represents an event or happening in the system.
type WhatsHappening struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	Name        string    `gorm:"type:varchar(50);not null" json:"name"`
	Image       string    `gorm:"type:varchar(150)" json:"image"`
	EventType   string    `gorm:"type:varchar(50)" json:"event_type"`
	Country     string    `gorm:"type:varchar(50);default:'NGA'" json:"country"`
	Description string    `gorm:"type:varchar(50)" json:"description"`
	Address     string    `gorm:"type:text" json:"address"`
	Category    string    `gorm:"type:varchar(50)" json:"category"`
	Duration    string    `gorm:"type:varchar(50)" json:"duration"`
	StartTime   time.Time `gorm:"type:timestamp with time zone;not null" json:"start_time"`
	EndTime     time.Time `gorm:"type:timestamp with time zone;not null" json:"end_time"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}

// BeforeCreate is a GORM hook that generates a new UUID before creating a record.
func (event *WhatsHappening) BeforeCreate(tx *gorm.DB) (err error) {
	event.ID = uuid.New()
	return
}
