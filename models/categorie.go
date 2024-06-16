package models

import (
	"time"

	"gorm.io/gorm"
)

type Categorie struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func GetCategories(db *gorm.DB) ([]Categorie, error) {
	var categories []Categorie
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
