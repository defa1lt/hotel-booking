package service

import (
	"github.com/google/uuid"
	"hotel-booking/internal/model"
	"hotel-booking/internal/repository"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) (*model.Booking, error)
	GetBookingByID(id string) (*model.Booking, error)
	UpdateBooking(id string, booking *model.Booking) (*model.Booking, error)
	DeleteBooking(id string) error
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}

func (s *bookingService) CreateBooking(booking *model.Booking) (*model.Booking, error) {
	// Генерация уникального ID для бронирования
	booking.ID = uuid.New().String()
	booking.Status = "confirmed"
	// Дополнительная бизнес-логика: проверка доступности номера и т.д.
	return s.repo.Create(booking)
}

func (s *bookingService) GetBookingByID(id string) (*model.Booking, error) {
	return s.repo.GetByID(id)
}

func (s *bookingService) UpdateBooking(id string, booking *model.Booking) (*model.Booking, error) {
	// Дополнительная бизнес-логика: проверка возможности обновления бронирования
	return s.repo.Update(id, booking)
}

func (s *bookingService) DeleteBooking(id string) error {
	// Дополнительная бизнес-логика: проверка возможности удаления бронирования
	return s.repo.Delete(id)
}
