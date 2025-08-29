package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	Pending        OrderStatus = "Pending"
	PaidOrder      OrderStatus = "Paid"
	Complete       OrderStatus = "Complete"
	CancelledOrder OrderStatus = "Cancelled"
	Refunded       OrderStatus = "Refunded"
)

type Orders struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserId      uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
	ShowtimeID  uuid.UUID   `gorm:"type:uuid;not null" json:"showtime_id"`
	TotalAmount float64     `gorm:"type:numeric(12,2);not null" json:"total_amount"`
	OrderStatus OrderStatus `gorm:"type:order_status;not null;default:'Pending'" json:"order_status"`
	CreatedAt   time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
