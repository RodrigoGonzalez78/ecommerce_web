package db

import (
	"log"

	"github.com/RodrigoGonzalez78/ecommerce_web/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConnection() {

	var dbError error
	db, dbError = gorm.Open(postgres.Open(config.Cnf.DatabaseURL), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("Base de datos conectada!!")
	}

}
