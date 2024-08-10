package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CreateAddress(data *models.Address) (uint, error) {
	// Crea la dirección en la base de datos
	if err := db.Create(data).Error; err != nil {
		return 0, err
	}

	// Retorna el ID de la dirección creada
	return data.ID, nil
}

func UpdateAddress(id uint, data map[string]interface{}) error {
	if err := db.Model(&models.Address{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func GetAddress(id uint) (*models.Address, error) {
	var address models.Address
	if err := db.First(&address, id).Error; err != nil {
		return nil, err
	}
	return &address, nil
}
