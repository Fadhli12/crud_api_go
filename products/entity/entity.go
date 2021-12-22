package entity

import (
	"time"
)

type Products struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	Price     int       `json:"price"`
	BrandId   int       `json:"brand_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
