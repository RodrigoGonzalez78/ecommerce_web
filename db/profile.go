package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func GetProfiles() ([]models.Profile, error) {
	var profiles []models.Profile
	if err := db.Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}
