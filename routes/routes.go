package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/backend/controllers"
	"github.com/sahilchauhan0603/backend/middleware"
)

func InitializeRoutes(router *mux.Router) {

	// Apply middleware to all routes
	router.Use(middleware.JWTVerify)

	// routes for Uploader Table
	router.HandleFunc("/upload", controllers.CreateUploader).Methods("POST")
	router.HandleFunc("/upload", controllers.GetUploader).Methods("GET")
	router.HandleFunc("/upload/{id}", controllers.GetUploaderID).Methods("GET")
	router.HandleFunc("/uploader/{id}", controllers.UpdateUploader).Methods("PUT")
	router.HandleFunc("/upload/{id}", controllers.DeleteUploader).Methods("DELETE")

	// routes for Project Table
	router.HandleFunc("/project", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/project", controllers.GetProject).Methods("GET")
	router.HandleFunc("/project/{id}", controllers.GetProjectID).Methods("GET")
	router.HandleFunc("/project/{id}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/project/{id}", controllers.DeleteProject).Methods("DELETE")

	// routes for Admin Table
	router.HandleFunc("/admin", controllers.CreateAdmin).Methods("POST")
	router.HandleFunc("/admin", controllers.GetAdmin).Methods("GET")
	router.HandleFunc("/admin/{id}", controllers.GetAdminID).Methods("GET")
	router.HandleFunc("/admin/{id}", controllers.UpdateAdmin).Methods("PUT")
	router.HandleFunc("/admin/{id}", controllers.DeleteAdmin).Methods("DELETE")

	// Authentication routes
	router.HandleFunc("/login", controllers.HandleMicrosoftLogin).Methods("GET")
	router.HandleFunc("/callback", controllers.HandleMicrosoftCallback).Methods("GET")
}
