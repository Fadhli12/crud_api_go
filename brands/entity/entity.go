package entity

import "time"

type Brands struct {
	ID        int       `json:"id"`
	Name      string    `json:"title"`
	Logo      string    `json:"image"`
	Banner    string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
