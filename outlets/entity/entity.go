package entity

import (
	"time"
)

type Outlets struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Picture    string    `json:"picture"`
	Address    string    `json:"address"`
	Longitude  string    `json:"longitude"`
	Latitude   string    `json:"latitude"`
	BrandId    int       `json:"brand_id"`
	DistanceKm float32   `json:"distance_km"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"updated_at"`
}
