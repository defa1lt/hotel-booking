package repository

import (
	"database/sql"
	"hotel-booking/internal/model"
)

type RoomRepository interface {
	GetByHotelID(hotelID string) ([]model.Room, error)
	// Другие методы: создание, обновление, удаление номеров
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) GetByHotelID(hotelID string) ([]model.Room, error) {
	rows, err := r.db.Query("SELECT id, hotel_id, number, type, price, status FROM rooms WHERE hotel_id = $1", hotelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(&room.ID, &room.HotelID, &room.Number, &room.Type, &room.Price, &room.Status); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
