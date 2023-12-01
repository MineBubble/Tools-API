package models

import "github.com/lib/pq"

type Tool struct {
	ID          uint           `json:"id" gorm:""primaryKey"`
	Title       string         `json:"title"`
	Link        string         `json:"link"`
	Description string         `json:"description"`
	Tags        pq.StringArray `json:"tags" gorm:"type:varchar[]"`
}
