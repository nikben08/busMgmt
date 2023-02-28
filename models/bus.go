package models

type Bus struct {
	ID            string `json:"id,omitempty"`
	PlateNumber   string `json:"plate_number"`
	NumberOfSeats int    `json:"number_of_seats"`
	Model         string `json:"model"`
	Type          string `json:"type"`
}
