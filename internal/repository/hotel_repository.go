package repository

import (
	"database/sql"
	"hotel-booking/internal/model"
)

type HotelRepository interface {
	GetAll() ([]model.Hotel, error)
	GetByID(id string) (*model.Hotel, error)
	// Другие методы, например, создание, обновление, удаление отелей
}

type hotelRepository struct {
	db *sql.DB
}

func NewHotelRepository(db *sql.DB) HotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) GetAll() ([]model.Hotel, error) {
	rows, err := r.db.Query("SELECT id, name, location, description, rating FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []model.Hotel
	for rows.Next() {
		var hotel model.Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.Location, &hotel.Description, &hotel.Rating); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func (r *hotelRepository) GetByID(id string) (*model.Hotel, error) {
	var hotel model.Hotel
	err := r.db.QueryRow("SELECT id, name, location, description, rating FROM hotels WHERE id = $1", id).
		Scan(&hotel.ID, &hotel.Name, &hotel.Location, &hotel.Description, &hotel.Rating)
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}
