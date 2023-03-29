package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "localhost"
	DB_PORT     = 5432
	DB_NAME     = "dblearngorm"
	// DEBUG_MODE  = true // true/false
)

var (
	db  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	// db.Debug().AutoMigrate(models.Books{})
	return db
}
