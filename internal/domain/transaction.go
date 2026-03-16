package domain

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string
	Type      string
	Amount    float64
	Category  string
	CreatedAt time.Time
}