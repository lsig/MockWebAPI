package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lsig/OtpWebAPI/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDatabase() (error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New("Error connecting to database")
	}

	return nil

}

func Migrate() error {
	err := database.AutoMigrate(&model.User{})

	if err != nil {
		return errors.New("Schema migration failed")
	}

	return nil
}

func GetDB() *gorm.DB {
	return database
}
