package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CreateConsult(data models.Consult) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetAllNewConsults() ([]models.Consult, error) {
	var consults []models.Consult
	if err := db.Where("archived = ?", "NO").Find(&consults).Error; err != nil {
		return nil, err
	}
	return consults, nil
}

func GetAllArchivedConsults() ([]models.Consult, error) {
	var consults []models.Consult
	if err := db.Where("archived = ?", "SI").Find(&consults).Error; err != nil {
		return nil, err
	}
	return consults, nil
}

func UpdateConsult(id uint, data map[string]interface{}) error {
	if err := db.Model(&models.Consult{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
