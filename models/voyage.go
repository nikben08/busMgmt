package models

type Voyage struct {
	ID    string  `json:"id,omitempty"`
	Fee   float32 `json:"fee,omitempty"`
	From  string  `json:"from,omitempty"`
	To    string  `json:"to,omitempty"`
	Date  string  `json:"date,omitempty"`
	BusID string  `json:"bus_id,omitempty"`
}

type VoyageResponse struct {
	ID    string      `json:"id,omitempty"`
	Fee   float32     `json:"fee,omitempty"`
	From  string      `json:"from,omitempty"`
	To    string      `json:"to,omitempty"`
	Date  string      `json:"date,omitempty"`
	BusID string      `json:"bus_id,omitempty"`
	Bus   BusResponse `json:"bus"`
}

type BusResponse struct {
	ID            string        `json:"id,omitempty"`
	PlateNumber   string        `json:"plate_number"`
	NumberOfSeats int           `json:"number_of_seats"`
	Model         string        `json:"model"`
	Type          string        `json:"type"`
	Seat          []BusSeat     `json:"seat"`
	Property      []BusProperty `json:"property"`
}

type BusSeatResponse struct {
	ID       string `json:"id,omitempty"`
	BusID    string `json:"bus_id,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	VoyageID string `json:"voyage_id,omitempty"`
	SeatNo   int    `json:"seat_no,omitempty"`
	Sex      int    `json:"sex,omitempty"`
}
