package models

type BusSeat struct {
	ID       string `json:"id,omitempty"`
	BusID    string `json:"bus_id,omitempty"`
	VoyageID string `json:"voyage_id,omitempty"`
	SeatNo   int    `json:"seat_no,omitempty"`
	Sex      int    `json:"sex,omitempty"`
}
