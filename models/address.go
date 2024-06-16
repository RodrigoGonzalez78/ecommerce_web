package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID           uint      `gorm:"primaryKey"`
	Street       string    `gorm:"column:street"`
	PostalCode   string    `gorm:"column:postal_code"`
	Neighborhood string    `gorm:"column:neighborhood"`
	City         string    `gorm:"column:city"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

func CreateAddress(db *gorm.DB, data Address) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAddress(db *gorm.DB, id uint, data map[string]interface{}) error {
	if err := db.Model(&Address{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func GetAddress(db *gorm.DB, id uint) (*Address, error) {
	var address Address
	if err := db.First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}
