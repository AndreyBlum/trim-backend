package domain

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}
