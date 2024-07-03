package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CreateSale(data models.Sale) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func FindSale(id uint) (*models.Sale, error) {
	var sale models.Sale
	if err := db.First(&sale, id).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}

func AllSales() ([]map[string]interface{}, error) {
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

func UserSales(id uint) ([]map[string]interface{}, error) {
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
