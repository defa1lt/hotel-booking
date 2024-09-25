package routes

import (
	"hotel-booking/internal/api/handlers"
	"hotel-booking/internal/api/middleware"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func SetupRouter(log *logrus.Logger) *mux.Router {
	r := mux.NewRouter()

	// Применение middleware
	r.Use(middleware.LoggingMiddleware(log))
	r.Use(middleware.RecoverMiddleware(log))

	// Маршруты для отелей
	r.HandleFunc("/hotels", handlers.GetHotels).Methods("GET")
	r.HandleFunc("/hotels/{id}", handlers.GetHotelByID).Methods("GET")
	r.HandleFunc("/hotels/{id}/rooms", handlers.GetRoomsByHotel).Methods("GET")

	// Маршруты для бронирований
	r.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", handlers.GetBooking).Methods("GET")
	r.HandleFunc("/bookings/{id}", handlers.UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", handlers.DeleteBooking).Methods("DELETE")

	return r
}
