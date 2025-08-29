package models

import (
	"time"

	"github.com/google/uuid"
)

type Showtimes struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MovieId   uuid.UUID `gorm:"type:uuid;not null" json:"movie_id"`
	TheaterId uuid.UUID `gorm:"type:uuid;not null" json:"theater_id"`
	StartAt   time.Time `gorm:"type:timestamp;not null" json:"start_at"`
	BasePrice float64   `gorm:"type:numeric(12,2);not null" json:"base_price"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Movie     Movies    `gorm:"foreignKey:MovieId;references:ID;constraint:OnDelete:CASCADE"`
	Theater   Theaters  `gorm:"foreignKey:TheaterId;references:ID;constraint:OnDelete:CASCADE"`
}

type ListShowtimes struct {
	ID          uuid.UUID `json:"id"`
	MovieName   string    `json:"movie_name"`
	TheaterName string    `json:"theater_name"`
	StartAt     time.Time `json:"start_at"`
	BasePrice   float64   `json:"base_price"`
	Seats       []ListSeats   `json:"seats"`
}

type ShowtimeInput struct {
	MovieId   uuid.UUID `json:"movie_id" binding:"required,uuid"`
	TheaterId uuid.UUID `json:"theater_id" binding:"required,uuid"`
	StartAt   string `json:"start_at" binding:"required"`
	BasePrice float64   `json:"base_price" binding:"required,gt=0"`
	Seats     []Seats   `json:"seats" binding:"required,dive"`
}
