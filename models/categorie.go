package models

type Categorie struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}
