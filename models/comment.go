package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Cerpen_id int       `gorm:"foreignKey:ID" json:"cerpen_id"`
	Name      string    `gorm:"type:varchar(200)" json:"name"`
	Body      string    `gorm:"type:varchar(200)" json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Cerpen_id string    `json:"cerpen_id"`
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
