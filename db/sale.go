package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CreateSale(data models.Sale) (*models.Sale, error) {
	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
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
		Select("sales.*, users.name as user_name, users.last_name as last_name, users.email as email").
		Joins("join users on users.id = sales.id_user").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func UserSales(id uint) ([]models.Sale, error) {
	var sales []models.Sale
	err := db.Table("sales").
		Select("id, id_user, total_price, date").
		Where("id_user = ?", id).
		Scan(&sales).Error
	if err != nil {
		return nil, err
	}
	return sales, nil
}
