package model

type Hotel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}
