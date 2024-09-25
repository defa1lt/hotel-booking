package main

import (
	"database/sql"
	"log"
	"net/http"

	"hotel-booking/internal/api/routes"
	"hotel-booking/internal/repository"
	"hotel-booking/internal/service"
	"hotel-booking/pkg/config"
	"hotel-booking/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация логгера
	log := logger.NewLogger(cfg.LogLevel)

	// Подключение к базе данных
	db, err := sql.Open("postgres", "host="+cfg.DBHost+" port="+cfg.DBPort+" user="+cfg.DBUser+" "+
		"password="+cfg.DBPassword+" dbname="+cfg.DBName+" sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка проверки подключения к базе данных: %v", err)
	}

	// Инициализация репозиториев
	hotelRepo := repository.NewHotelRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	bookingRepo := repository.NewBookingRepository(db)

	// Инициализация сервисов
	hotelService := service.NewHotelService(hotelRepo)
	roomService := service.NewRoomService(roomRepo)
	bookingService := service.NewBookingService(bookingRepo)

	// Передача сервисов в обработчики
	handlers.InitializeHandlers(hotelService, roomService, bookingService)

	// Инициализация маршрутов
	router := routes.SetupRouter(log)

	// Запуск сервера
	log.Infof("Запуск сервера на порту %s", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
