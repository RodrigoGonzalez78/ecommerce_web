package models

type SaleDetails struct {
	ID        uint    `gorm:"primaryKey"`
	IDSale    uint    `gorm:"column:id_sale"`
	Count     int     `gorm:"column:count"`
	Price     float64 `gorm:"column:price"`
	IDProduct uint    `gorm:"column:id_product"`
}
