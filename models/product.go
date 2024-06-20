package models

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"column:name"`
	Price       float64 `gorm:"column:price"`
	Stock       int     `gorm:"column:stock"`
	Description string  `gorm:"column:description"`
	Down        string  `gorm:"column:down"`
	Image       string  `gorm:"column:image"`
	IDCategorie uint    `gorm:"column:id_categorie"`
}
