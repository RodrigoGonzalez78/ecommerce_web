package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	ID         uint      `gorm:"primaryKey"`
	IDUser     uint      `gorm:"column:id_user"`
	TotalPrice float64   `gorm:"column:total_price"`
	Date       time.Time `gorm:"column:date"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (Sale) TableName() string {
	return "sales"
}

func CreateSale(db *gorm.DB, data Sale) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func FindSale(db *gorm.DB, id uint) (*Sale, error) {
	var sale Sale
	if err := db.First(&sale, id).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}

func AllSales(db *gorm.DB) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := db.Table("sales").
		Select("sales.*, users.name, users.last_name, users.email").
		Joins("join users on users.id = sales.id_user").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func UserSales(db *gorm.DB, id uint) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := db.Table("sales").
		Select("sales.*, users.name, users.last_name, users.email").
		Joins("join users on users.id = sales.id_user").
		Where("users.id = ?", id).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
