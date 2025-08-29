package models

import (
	"time"

	"github.com/google/uuid"
)

type Movies struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title           string    `gorm:"type:text;not null" json:"title"`
	DurationMinutes int       `gorm:"type:integer" json:"duration_minutes"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
