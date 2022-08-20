package database

import (
	"fmt"
	"log"
	"os"

	"github.com/tech-thinker/linkly/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// IsConnected returns the connection status
	IsConnected bool
	// IsSQLite returns the connection status
	IsSQLite bool
	// DB returns the database connection
	DB *gorm.DB
)

func GetDB() *gorm.DB {
	var err error
	// Get ENV variables
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	if DB == nil {
		if dbHost == "" {
			fmt.Println("Environment variable DB_HOST is null.")
			return nil
		}
		if dbName == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil
		}
		if dbUser == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil
		}
		if dbPassword == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil
		}

		if dbPort == "" {
			dbPort = "5432"
		}
	}

	// Connect to database
	dest := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dest), &gorm.Config{})
	if err == nil {
		IsConnected = true
	} else {
		log.Println("failed to connect database")
	}

	// if unable to connect to database, create sqlite database
	if !IsConnected {
		// Create sqlite connection
		DB, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
		if err == nil {
			IsConnected = true
			IsSQLite = true
		} else {
			log.Println("failed to connect database")
		}
	}

	// Migrate the schema
	DB.AutoMigrate(&models.URL{})
	// Link the database to the models
	DB.AutoMigrate(&models.Link{})
	// Domain the database to the models
	DB.AutoMigrate(&models.Domain{})

	return DB
}
