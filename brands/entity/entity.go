package entity

import "time"

type Brands struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Logo      string    `json:"logo"`
	Banner    string    `json:"banner"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
