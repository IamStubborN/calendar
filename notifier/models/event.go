package models

type Event struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
