package models

import (
	"time"

	"gorm.io/gorm"
)

type Consult struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Description string    `gorm:"column:description"`
	Attended    string    `gorm:"column:attended"`
	Archived    string    `gorm:"column:archived"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func CreateConsult(db *gorm.DB, data Consult) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetAllNewConsults(db *gorm.DB) ([]Consult, error) {
	var consults []Consult
	if err := db.Where("archived = ?", "NO").Find(&consults).Error; err != nil {
		return nil, err
	}
	return consults, nil
}

func GetAllArchivedConsults(db *gorm.DB) ([]Consult, error) {
	var consults []Consult
	if err := db.Where("archived = ?", "SI").Find(&consults).Error; err != nil {
		return nil, err
	}
	return consults, nil
}

func UpdateConsult(db *gorm.DB, id uint, data map[string]interface{}) error {
	if err := db.Model(&Consult{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
