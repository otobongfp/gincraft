package models

import (
	"time"
)

// Example represents an example model
type Example struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} 