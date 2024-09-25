package repository

import (
	"database/sql"
	"hotel-booking/internal/model"
)

type BookingRepository interface {
	Create(booking *model.Booking) (*model.Booking, error)
	GetByID(id string) (*model.Booking, error)
	Update(id string, booking *model.Booking) (*model.Booking, error)
	Delete(id string) error
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *model.Booking) (*model.Booking, error) {
	query := `INSERT INTO bookings (id, hotel_id, room_id, user_id, start_date, end_date, status)
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, booking.ID, booking.HotelID, booking.RoomID, booking.UserID, booking.StartDate, booking.EndDate, booking.Status)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *bookingRepository) GetByID(id string) (*model.Booking, error) {
	var booking model.Booking
	query := `SELECT id, hotel_id, room_id, user_id, start_date, end_date, status
              FROM bookings WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&booking.ID, &booking.HotelID, &booking.RoomID, &booking.UserID, &booking.StartDate, &booking.EndDate, &booking.Status)
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) Update(id string, booking *model.Booking) (*model.Booking, error) {
	query := `UPDATE bookings SET hotel_id = $1, room_id = $2, user_id = $3, start_date = $4, end_date = $5, status = $6
              WHERE id = $7`
	_, err := r.db.Exec(query, booking.HotelID, booking.RoomID, booking.UserID, booking.StartDate, booking.EndDate, booking.Status, id)
	if err != nil {
		return nil, err
	}
	booking.ID = id
	return booking, nil
}

func (r *bookingRepository) Delete(id string) error {
	query := `DELETE FROM bookings WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
