package models

import (
	"time"
)

type Item struct {
	Id          uint      `json:"item_id"`
	Name        string    `json:"item_name"`
	Description string    `json:"item_description"`
	DateFound   time.Time `json:"date_found"`
	WhereFound  string    `json:"where_found"`
	Status      string    `json:"item_status"`
}
