package models

import (
	"time"
)

type Sale struct {
	ID         uint      `gorm:"primaryKey"`
	IDUser     uint      `gorm:"column:id_user"`
	TotalPrice float64   `gorm:"column:total_price"`
	Date       time.Time `gorm:"column:date"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
