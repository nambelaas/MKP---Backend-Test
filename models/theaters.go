package models

import (
	"time"

	"github.com/google/uuid"
)

type Theaters struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	City      string    `gorm:"type:text;not null" json:"city"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
