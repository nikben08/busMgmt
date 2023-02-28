package models

type User struct {
	ID          string `json:"id,omitempty"`
	Username    string `json:"username,omitempty"`
	Hash        string `json:"hash,omitempty"`
	AccessLevel string `json:"access_level,omitempty"`
}
