package models

import (
	"time"
)

type Consult struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Description string    `gorm:"column:description"`
	Attended    string    `gorm:"column:attended"`
	Archived    string    `gorm:"column:archived"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
