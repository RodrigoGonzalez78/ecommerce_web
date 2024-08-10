package models

type Address struct {
	ID           uint   `gorm:"primaryKey"`
	Street       string `gorm:"column:street"`
	PostalCode   string `gorm:"column:postal_code"`
	Neighborhood string `gorm:"column:neighborhood"`
	City         string `gorm:"column:city"`
}
