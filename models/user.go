package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Down      bool      `gorm:"column:down"`
	IDAddress uint      `gorm:"column:id_address"`
	IDProfile uint      `gorm:"column:id_profile"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (User) TableName() string {
	return "users"
}

func GetUser(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *gorm.DB, id uint, data map[string]interface{}) error {
	if err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
