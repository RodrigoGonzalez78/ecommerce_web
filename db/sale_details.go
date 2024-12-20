package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_web/models"
)

func CreateSaleDetails(data models.SaleDetails) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetSaleDetailsByIdAndProductName(id uint) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := db.Table("sale_details").
		Select("sale_details.price ,sale_details.count ,sale_details.price, products.name").
		Joins("join products on products.id = sale_details.id_product").
		Where("sale_details.id_sale = ?", id).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
