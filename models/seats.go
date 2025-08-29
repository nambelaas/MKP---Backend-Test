package models

import (
	"time"

	"github.com/google/uuid"
)

type SeatStatus string

const (
	Available SeatStatus = "Available"
	Hold  SeatStatus = "Hold"
	Paid  SeatStatus = "Paid"
	Cancelled  SeatStatus = "Cancelled"
	Expired  SeatStatus = "Expired"
)

type Seats struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ShowtimeId uuid.UUID `gorm:"type:uuid;not null" json:"showtime_id"`
	SeatCode string `gorm:"type:text;not null" json:"seat_code"`
	Status SeatStatus `gorm:"type:seat_status;not null;default:'Available'" json:"status"`
	HoldUntil time.Time `gorm:"type:timestamp" json:"hold_until"`
	OrderID uuid.UUID `gorm:"type:uuid" json:"order_id"`
	Showtime Showtimes `gorm:"foreignKey:ShowtimeId;references:ID;constraint:OnDelete:CASCADE"`
}

type ListSeats struct {
	SeatCode string `json:"seat_code"`
	Status   string `json:"status"`
}