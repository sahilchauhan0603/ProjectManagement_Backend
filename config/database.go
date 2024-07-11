package database

import (
	"log"
	"os"

	models "github.com/sahilchauhan0603/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnector() {

	// Retrieve environment variables for the database
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbName == "" || dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Missing required database environment variables")
	}

	// Now, connect to the newly created or existing database
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(&models.Uploader{})
	if err != nil {
		log.Fatalf("failed to migrate Uploader: %v", err)
	}

	// Then migrate the related tables
	err = DB.AutoMigrate(&models.Project{}, &models.Admin{})
	if err != nil {
		log.Fatalf("failed to migrate related tables: %v", err)
	}
}
