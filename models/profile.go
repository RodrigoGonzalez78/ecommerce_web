package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID          uint      `gorm:"primaryKey"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func GetProfiles(db *gorm.DB) ([]Profile, error) {
	var profiles []Profile
	if err := db.Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}
