package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Down      string `gorm:"column:down;default:null"`
	IDAddress *uint  `gorm:"column:id_address;default:null"`
	IDProfile uint   `gorm:"column:id_profile"`
}
