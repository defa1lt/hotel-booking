package model

import (
	"fmt"
	"time"
)

type Booking struct {
	ID        string    `json:"id"`
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"` // Например: "confirmed", "cancelled"
}

// Валидация данных бронирования
func (b *Booking) Validate() error {
	if b.HotelID == "" || b.RoomID == "" || b.UserID == "" {
		return fmt.Errorf("отсутствуют обязательные поля")
	}
	if b.StartDate.After(b.EndDate) {
		return fmt.Errorf("дата начала бронирования должна быть раньше даты окончания")
	}
	// Дополнительные проверки...
	return nil
}
