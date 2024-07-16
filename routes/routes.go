package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/backend/controllers"
	"github.com/sahilchauhan0603/backend/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitializeRoutes(router *mux.Router) {
    
	// Auth route	
	router.HandleFunc("/api/v1/validateToken", middleware.ValidateTokenHandler).Methods("POST")

	// Handle preflight requests for the /api/v1 endpoints
    router.PathPrefix("/api/v1").Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        w.WriteHeader(http.StatusNoContent)
    })

	// Authentication routes
	// @Summary Microsoft login
	// @Description Initiate Microsoft login
	// @Tags Authentication
	// @Router /login [get]
	router.HandleFunc("/login", controllers.HandleMicrosoftLogin).Methods("GET")

	// @Summary Microsoft login callback
	// @Description Handle Microsoft login callback
	// @Tags Authentication
	// @Router /callback [get]
	router.HandleFunc("/callback", controllers.HandleMicrosoftCallback).Methods("GET")

	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)


	// Uploader routes
	uploaderRouter := router.PathPrefix("/api/v1").Subrouter()
	uploaderRouter.Use(middleware.JWTVerify)

	// @Summary Create a new uploader
	// @Description Create a new uploader
	// @Tags Uploaders
	// @Accept json
	// @Produce json
	// @Param uploader body models.Uploader true "Uploader information"
	// @Success 201 {object} models.Uploader
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/upload [post]
	uploaderRouter.HandleFunc("/upload", controllers.CreateUploader).Methods("POST")

	// @Summary Get all uploaders
	// @Description Retrieve all uploaders
	// @Tags Uploaders
	// @Produce json
	// @Success 200 {array} models.Uploader
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /api/v1/upload [get]
	uploaderRouter.HandleFunc("/upload", controllers.GetUploader).Methods("GET")

	// @Summary Get uploader by ID
	// @Description Retrieve uploader by ID
	// @Tags Uploaders
	// @Produce json
	// @Param id path int true "Uploader ID"
	// @Success 200 {object} models.Uploader
	// @Failure 404 {string} string "Uploader not found"
	// @Router /api/v1/upload/{id} [get]
	uploaderRouter.HandleFunc("/upload/{id}", controllers.GetUploaderID).Methods("GET")

	// @Summary Update uploader
	// @Description Update uploader by ID
	// @Tags Uploaders
	// @Accept json
	// @Produce json
	// @Param id path int true "Uploader ID"
	// @Param uploader body models.Uploader true "Uploader information"
	// @Success 200 {object} models.Uploader
	// @Failure 404 {string} string "Uploader not found"
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/upload/{id} [put]
	uploaderRouter.HandleFunc("/upload/{id}", controllers.UpdateUploader).Methods("PUT")

	// @Summary Delete uploader
	// @Description Delete uploader by ID
	// @Tags Uploaders
	// @Param id path int true "Uploader ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Uploader not found"
	// @Router /api/v1/upload/{id} [delete]
	uploaderRouter.HandleFunc("/upload/{id}", controllers.DeleteUploader).Methods("DELETE")

	// Project routes
	projectRouter := router.PathPrefix("/api/v1").Subrouter()
	projectRouter.Use(middleware.JWTVerify)

	// @Summary Create a new project
	// @Description Create a new project
	// @Tags Projects
	// @Accept json
	// @Produce json
	// @Param project body models.Project true "Project information"
	// @Success 201 {object} models.Project
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/project [post]
	projectRouter.HandleFunc("/project", controllers.CreateProject).Methods("POST")

	// @Summary Get all projects
	// @Description Retrieve all projects
	// @Tags Projects
	// @Produce json
	// @Success 200 {array} models.Project
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /api/v1/project [get]
	projectRouter.HandleFunc("/project", controllers.GetProject).Methods("GET")

	// @Summary Get project by ID
	// @Description Retrieve project by ID
	// @Tags Projects
	// @Produce json
	// @Param id path int true "Project ID"
	// @Success 200 {object} models.Project
	// @Failure 404 {string} string "Project not found"
	// @Router /api/v1/project/{id} [get]
	projectRouter.HandleFunc("/project/{id}", controllers.GetProjectID).Methods("GET")

	// @Summary Update project
	// @Description Update project by ID
	// @Tags Projects
	// @Accept json
	// @Produce json
	// @Param id path int true "Project ID"
	// @Param project body models.Project true "Project information"
	// @Success 200 {object} models.Project
	// @Failure 404 {string} string "Project not found"
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/project/{id} [put]
	projectRouter.HandleFunc("/project/{id}", controllers.UpdateProject).Methods("PUT")

	// @Summary Delete project
	// @Description Delete project by ID
	// @Tags Projects
	// @Param id path int true "Project ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Project not found"
	// @Router /api/v1/project/{id} [delete]
	projectRouter.HandleFunc("/project/{id}", controllers.DeleteProject).Methods("DELETE")

	// Admin routes
	adminRouter := router.PathPrefix("/api/v1").Subrouter()
	adminRouter.Use(middleware.JWTVerify)

	// @Summary Create a new admin
	// @Description Create a new admin
	// @Tags Admins
	// @Accept json
	// @Produce json
	// @Param admin body models.Admin true "Admin information"
	// @Success 201 {object} models.Admin
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/admin [post]
	adminRouter.HandleFunc("/admin", controllers.CreateAdmin).Methods("POST")

	// @Summary Get all admins
	// @Description Retrieve all admins
	// @Tags Admins
	// @Produce json
	// @Success 200 {array} models.Admin
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /api/v1/admin [get]
	adminRouter.HandleFunc("/admin", controllers.GetAdmin).Methods("GET")

	// @Summary Get admin by ID
	// @Description Retrieve admin by ID
	// @Tags Admins
	// @Produce json
	// @Param id path int true "Admin ID"
	// @Success 200 {object} models.Admin
	// @Failure 404 {string} string "Admin not found"
	// @Router /api/v1/admin/{id} [get]
	adminRouter.HandleFunc("/admin/{id}", controllers.GetAdminID).Methods("GET")

	// @Summary Update admin
	// @Description Update admin by ID
	// @Tags Admins
	// @Accept json
	// @Produce json
	// @Param id path int true "Admin ID"
	// @Param admin body models.Admin true "Admin information"
	// @Success 200 {object} models.Admin
	// @Failure 404 {string} string "Admin not found"
	// @Failure 400 {string} string "Bad Request"
	// @Router /api/v1/admin/{id} [put]
	adminRouter.HandleFunc("/admin/{id}", controllers.UpdateAdmin).Methods("PUT")

	// @Summary Delete admin
	// @Description Delete admin by ID
	// @Tags Admins
	// @Param id path int true "Admin ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Admin not found"
	// @Router /api/v1/admin/{id} [delete]
	adminRouter.HandleFunc("/admin/{id}", controllers.DeleteAdmin).Methods("DELETE")

}
