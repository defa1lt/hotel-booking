package handlers

import (
	"encoding/json"
	"net/http"

	"hotel-booking/internal/model"
	"hotel-booking/internal/service"
	"hotel-booking/internal/utils"

	"github.com/gorilla/mux"
)

var bookingService service.BookingService

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	// Валидация данных бронирования
	if err := booking.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Создание бронирования через сервис
	createdBooking, err := bookingService.CreateBooking(&booking)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Ошибка сервера")
		return
	}

	// Формирование ответа в формате HAL
	halResponse := utils.NewHALResponse(createdBooking)
	halResponse.AddLink("self", "/bookings/"+createdBooking.ID)
	halResponse.AddLink("hotel", "/hotels/"+createdBooking.HotelID)
	halResponse.AddLink("room", "/rooms/"+createdBooking.RoomID)
	halResponse.AddLink("user", "/users/"+createdBooking.UserID)
	halResponse.AddLink("cancel", map[string]interface{}{
		"href":   "/bookings/" + createdBooking.ID,
		"method": "DELETE",
	})
	halResponse.AddLink("update", map[string]interface{}{
		"href":   "/bookings/" + createdBooking.ID,
		"method": "PUT",
	})

	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(halResponse)
}

func GetBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	booking, err := bookingService.GetBookingByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Бронирование не найдено")
		return
	}

	halResponse := utils.NewHALResponse(booking)
	halResponse.AddLink("self", "/bookings/"+booking.ID)
	halResponse.AddLink("hotel", "/hotels/"+booking.HotelID)
	halResponse.AddLink("room", "/rooms/"+booking.RoomID)
	halResponse.AddLink("user", "/users/"+booking.UserID)
	halResponse.AddLink("cancel", map[string]interface{}{
		"href":   "/bookings/" + booking.ID,
		"method": "DELETE",
	})
	halResponse.AddLink("update", map[string]interface{}{
		"href":   "/bookings/" + booking.ID,
		"method": "PUT",
	})

	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(halResponse)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var booking model.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Неверный формат запроса")
		return
	}

	// Обновление бронирования через сервис
	updatedBooking, err := bookingService.UpdateBooking(id, &booking)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Ошибка обновления бронирования")
		return
	}

	halResponse := utils.NewHALResponse(updatedBooking)
	halResponse.AddLink("self", "/bookings/"+updatedBooking.ID)
	halResponse.AddLink("hotel", "/hotels/"+updatedBooking.HotelID)
	halResponse.AddLink("room", "/rooms/"+updatedBooking.RoomID)
	halResponse.AddLink("user", "/users/"+updatedBooking.UserID)
	halResponse.AddLink("cancel", map[string]interface{}{
		"href":   "/bookings/" + updatedBooking.ID,
		"method": "DELETE",
	})
	halResponse.AddLink("update", map[string]interface{}{
		"href":   "/bookings/" + updatedBooking.ID,
		"method": "PUT",
	})

	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(halResponse)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := bookingService.DeleteBooking(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Ошибка удаления бронирования")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
