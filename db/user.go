package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
	"gorm.io/gorm"
)

func CheckExistUser(email string) (bool, models.User, error) {
	var user models.User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, user, nil
		}
		return false, user, result.Error
	}

	return true, user, nil
}

func GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id uint, data map[string]interface{}) error {
	if err := db.Model(&models.User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser crea un nuevo usuario en la base de datos
func CreateUser(user *models.User) error {
	result := db.Create(user)
	return result.Error
}
