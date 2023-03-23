package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	conncetion := "host=localhost user=root dbname=mobile_02 password=root port=5440 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conncetion))
	log.Println("Database connection successful")
	if err != nil {
		log.Panic("Database connection error")
	}
}
