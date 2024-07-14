package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/sahilchauhan0603/backend/docs" // Swagger docs
	// "github.com/swaggo/http-swagger"             // Importing the Swagger handler

	database "github.com/sahilchauhan0603/backend/config"
	routes "github.com/sahilchauhan0603/backend/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	database.DatabaseConnector()

	// Create a new router
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	// Set the port for the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatalf("Failed to start server: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
