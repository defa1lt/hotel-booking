package handlers

import (
	"encoding/json"
	"net/http"

	"hotel-booking/internal/service"
	"hotel-booking/internal/utils"

	"github.com/gorilla/mux"
)

var roomService service.RoomService

func GetRoomsByHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hotelID := vars["id"]

	rooms, err := roomService.GetRoomsByHotelID(hotelID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Ошибка получения номеров")
		return
	}

	halResponse := utils.NewHALResponse(rooms)
	halResponse.AddLink("self", "/hotels/"+hotelID+"/rooms")
	halResponse.AddLink("hotel", "/hotels/"+hotelID)

	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(halResponse)
}
