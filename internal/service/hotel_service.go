package service

import (
	"hotel-booking/internal/model"
	"hotel-booking/internal/repository"
)

type HotelService interface {
	GetAllHotels() ([]model.Hotel, error)
	GetHotelByID(id string) (*model.Hotel, error)
	// Другие методы: создание, обновление, удаление отелей
}

type hotelService struct {
	repo repository.HotelRepository
}

func NewHotelService(repo repository.HotelRepository) HotelService {
	return &hotelService{
		repo: repo,
	}
}

func (s *hotelService) GetAllHotels() ([]model.Hotel, error) {
	return s.repo.GetAll()
}

func (s *hotelService) GetHotelByID(id string) (*model.Hotel, error) {
	return s.repo.GetByID(id)
}
