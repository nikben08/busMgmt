package models

type BusModel struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	BrandId string `json:"brand_id"`
}
