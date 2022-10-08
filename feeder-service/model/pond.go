package model

import "time"

type Pond struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Uuid      string    `json:"uuid" gorm:"unique"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
