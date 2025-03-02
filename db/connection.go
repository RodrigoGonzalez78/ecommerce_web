package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "user=postgres password=12345678 dbname=ecommerce host=localhost port=5432 sslmode=disable"

// Conexion de a la base de datos
var db *gorm.DB

func DBConnection() {


	dsn := os.Getenv("DATABASE_URL")
	
	if dsn == "" {
		log.Fatal("DATABASE_URL no está configurada")
	}

	var dbError error
	db, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("Base de datos conectada!!")
	}

}
