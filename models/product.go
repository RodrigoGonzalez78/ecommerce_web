package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Price       float64   `gorm:"column:price"`
	Stock       int       `gorm:"column:stock"`
	Description string    `gorm:"column:description"`
	Down        string    `gorm:"column:down"`
	Image       string    `gorm:"column:image"`
	IDCategorie uint      `gorm:"column:id_categorie"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (Product) TableName() string {
	return "products"
}

func GetProduct(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetDisabledProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Where("down = ?", "SI").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetEnabledProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Where("down = ?", "NO").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func SearchProducts(db *gorm.DB, searchTerm string, category string) ([]Product, error) {
	var products []Product
	query := db.Where("name LIKE ?", "%"+searchTerm+"%").Where("down = ?", "NO")
	if category != "Todos" {
		query = query.Where("id_categorie = ?", category)
	}
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(db *gorm.DB, id uint, data map[string]interface{}) error {
	if err := db.Model(&Product{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(db *gorm.DB, data Product) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
