package models

import (
	"time"

	"gorm.io/gorm"
)

type SaleDetails struct {
	ID        uint      `gorm:"primaryKey"`
	IDSale    uint      `gorm:"column:id_sale"`
	Count     int       `gorm:"column:count"`
	Price     float64   `gorm:"column:price"`
	IDProduct uint      `gorm:"column:id_product"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (SaleDetails) TableName() string {
	return "salesdetails"
}

func CreateSaleDetails(db *gorm.DB, data SaleDetails) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetSaleDetailsByIdAndProductName(db *gorm.DB, id uint) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := db.Table("salesdetails").
		Select("salesdetails.*, products.name").
		Joins("join products on products.id = salesdetails.id_product").
		Where("salesdetails.id_sale = ?", id).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
