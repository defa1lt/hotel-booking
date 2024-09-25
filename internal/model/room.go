package model

type Room struct {
	ID      string  `json:"id"`
	HotelID string  `json:"hotel_id"`
	Number  string  `json:"number"`
	Type    string  `json:"type"` // Например: "single", "double", "suite"
	Price   float64 `json:"price"`
	Status  string  `json:"status"` // Например: "available", "booked"
}
