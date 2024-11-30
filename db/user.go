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

func CreateUser(user *models.User) error {
	result := db.Create(user)
	return result.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}
func GetPaginatedUsers(page, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var totalUsers int64

	// Calcular el offset (desplazamiento)
	offset := (page - 1) * pageSize

	// Contar el total de usuarios
	if err := db.Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		return nil, 0, err
	}

	// Obtener los usuarios con paginación
	result := db.Limit(pageSize).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, totalUsers, nil
}
