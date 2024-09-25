package handlers

import (
	"encoding/json"
	"net/http"

	"hotel-booking/internal/service"
	"hotel-booking/internal/utils"

	"github.com/gorilla/mux"
)

var hotelService service.HotelService

func GetHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := hotelService.GetAllHotels()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Ошибка получения отелей")
		return
	}

	halResponse := utils.NewHALResponse(hotels)
	halResponse.AddLink("self", "/hotels")

	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(halResponse)
}

func GetHotelByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	hotel, err := hotelService.GetHotelByID(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Отель не найден")
		return
	}

	halResponse := utils.NewHALResponse(hotel)
	halResponse.AddLink("self", "/hotels/"+hotel.ID)
	halResponse.AddLink("rooms", "/hotels/"+hotel.ID+"/rooms")

	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(halResponse)
}
