package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
