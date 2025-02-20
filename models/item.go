package models

type Item struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	CreatedBy   string `json:"created_by"`
}
