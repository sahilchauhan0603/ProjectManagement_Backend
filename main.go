package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	database "github.com/sahilchauhan0603/backend/config"
	routes "github.com/sahilchauhan0603/backend/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.DatabaseConnector()

	router := mux.NewRouter()
	routes.InitializeRoutes(router)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
    }
	fmt.Printf("Server is running on port %s", port)
	log.Fatalf("Failed to start server: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
