package models

import (
	"time"
)

type Profile struct {
	ID          uint      `gorm:"primaryKey"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
