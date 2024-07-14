package database

import (
	"database/sql"
	"log"
	"os"

	models "github.com/sahilchauhan0603/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// DatabaseConnector initializes the database connection and migrates the necessary tables.
//
// This function retrieves environment variables for the database connection,
// checks if the database exists, creates it if it doesn't, and performs
// auto-migration for the Admin, Project, and Uploader models.
//
// It is expected that the following environment variables are set:
// - DB_NAME: Name of the database
// - DB_USER: Database username
// - DB_PASS: Database password
// - DB_HOST: Database host
// - DB_PORT: Database port
//
// If any required environment variable is missing or if the database connection
// or migration fails, the application will log an error and exit.
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

	// First, check if the database exists and create it if it doesn't
	serverDSN := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/"
	db, err := sql.Open("mysql", serverDSN)
	if err != nil {
		log.Fatal("failed to connect to MySQL server: ", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("failed to create database: ", err)
	}

	// Now, connect to the newly created or existing database
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// First migrate the admin table
	err = DB.AutoMigrate(&models.Admin{})
	if err != nil {
		log.Fatalf("failed to migrate Admin: %v", err)
	}

	// Then migrate the related tables
	err = DB.AutoMigrate(&models.Project{}, &models.Uploader{})
	if err != nil {
		log.Fatalf("failed to migrate related tables: %v", err)
	}
}
