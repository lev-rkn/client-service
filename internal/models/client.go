package models

type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
