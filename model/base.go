package model

import "time"

type Base struct {
	ID        int       `json:"id" gorm:"PrimaryKey"`
	Active    bool      `json:"active" gorm:"default: true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
