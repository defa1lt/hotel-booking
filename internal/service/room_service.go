package service

import (
	"hotel-booking/internal/model"
	"hotel-booking/internal/repository"
)

type RoomService interface {
	GetRoomsByHotelID(hotelID string) ([]model.Room, error)
	// Другие методы: создание, обновление, удаление номеров
}

type roomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) RoomService {
	return &roomService{
		repo: repo,
	}
}

func (s *roomService) GetRoomsByHotelID(hotelID string) ([]model.Room, error) {
	return s.repo.GetByHotelID(hotelID)
}
