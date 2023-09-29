package models

import "time"

type Cerpen struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(100)" json:"title"`
	Body  string `gorm:"type:varchar(200)" json:"body"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CerpenResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
