package database

import (
	"fmt"
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetDatabase() *gorm.DB {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	fmt.Println("The url:", url)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  url,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	err = db.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	return db
}

func CloseDatabase(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection for  close")
	}
	err = database.Close()
	if err != nil {
		log.Fatal("Failed to close database connection")
	}
}
