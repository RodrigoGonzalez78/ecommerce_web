package models

import (
	"time"
)

type Categorie struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
