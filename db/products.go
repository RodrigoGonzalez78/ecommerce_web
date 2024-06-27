package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func GetProduct(id uint) (*models.Product, error) {
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetDisabledProducts() ([]models.Product, error) {
	var products []models.Product
	if err := db.Where("down = ?", "SI").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetEnabledProducts() ([]models.Product, error) {
	var products []models.Product
	if err := db.Where("down = ?", "NO").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func SearchProducts(searchTerm string, category string) ([]models.Product, error) {
	var products []models.Product
	query := db.Where("name LIKE ?", "%"+searchTerm+"%").Where("down = ?", "NO")
	if category != "Todos" {
		query = query.Where("id_categorie = ?", category)
	}
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(id uint, data map[string]interface{}) error {
	if err := db.Model(&models.Product{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(data models.Product) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
