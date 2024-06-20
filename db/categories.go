package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func GetCategories() ([]models.Categorie, error) {
	var categories []models.Categorie
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
