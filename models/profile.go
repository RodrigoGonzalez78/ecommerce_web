package models

type Profile struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"column:description"`
}
