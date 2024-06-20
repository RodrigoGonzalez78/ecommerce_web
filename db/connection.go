package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "user=postgres password=12345678 dbname=taskM host=localhost port=5432 sslmode=disable"

// Conexion de a la base de datos
var db *gorm.DB

func DBConnection() {

	var dbError error
	db, dbError = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("Base de datos conectada!!")
	}

}
