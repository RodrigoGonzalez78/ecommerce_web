package models

import (
	"time"
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
