package routes

import (
	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/backend/controllers"
	"github.com/sahilchauhan0603/backend/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitializeRoutes initializes the API routes and applies middleware
func InitializeRoutes(router *mux.Router) {

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
	uploaderRouter := router.PathPrefix("/upload").Subrouter()
	uploaderRouter.Use(middleware.JWTVerify)

	// @Summary Create a new uploader
	// @Description Create a new uploader
	// @Tags Uploaders
	// @Accept json
	// @Produce json
	// @Param uploader body models.Uploader true "Uploader information"
	// @Success 201 {object} models.Uploader
	// @Failure 400 {string} string "Bad Request"
	// @Router /uploader/upload [post]
	uploaderRouter.HandleFunc("", controllers.CreateUploader).Methods("POST")

	// @Summary Get all uploaders
	// @Description Retrieve all uploaders
	// @Tags Uploaders
	// @Produce json
	// @Success 200 {array} models.Uploader
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /uploader/upload [get]
	uploaderRouter.HandleFunc("", controllers.GetUploader).Methods("GET")

	// @Summary Get uploader by ID
	// @Description Retrieve uploader by ID
	// @Tags Uploaders
	// @Produce json
	// @Param id path int true "Uploader ID"
	// @Success 200 {object} models.Uploader
	// @Failure 404 {string} string "Uploader not found"
	// @Router /uploader/upload/{id} [get]
	uploaderRouter.HandleFunc("/{id}", controllers.GetUploaderID).Methods("GET")

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
	// @Router /uploader/{id} [put]
	uploaderRouter.HandleFunc("/{id}", controllers.UpdateUploader).Methods("PUT")

	// @Summary Delete uploader
	// @Description Delete uploader by ID
	// @Tags Uploaders
	// @Param id path int true "Uploader ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Uploader not found"
	// @Router /uploader/upload/{id} [delete]
	uploaderRouter.HandleFunc("/{id}", controllers.DeleteUploader).Methods("DELETE")

	// Project routes
	projectRouter := router.PathPrefix("/project").Subrouter()
	projectRouter.Use(middleware.JWTVerify)

	// @Summary Create a new project
	// @Description Create a new project
	// @Tags Projects
	// @Accept json
	// @Produce json
	// @Param project body models.Project true "Project information"
	// @Success 201 {object} models.Project
	// @Failure 400 {string} string "Bad Request"
	// @Router /project [post]
	projectRouter.HandleFunc("", controllers.CreateProject).Methods("POST")

	// @Summary Get all projects
	// @Description Retrieve all projects
	// @Tags Projects
	// @Produce json
	// @Success 200 {array} models.Project
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /project [get]
	projectRouter.HandleFunc("", controllers.GetProject).Methods("GET")

	// @Summary Get project by ID
	// @Description Retrieve project by ID
	// @Tags Projects
	// @Produce json
	// @Param id path int true "Project ID"
	// @Success 200 {object} models.Project
	// @Failure 404 {string} string "Project not found"
	// @Router /project/{id} [get]
	projectRouter.HandleFunc("/{id}", controllers.GetProjectID).Methods("GET")

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
	// @Router /project/{id} [put]
	projectRouter.HandleFunc("/{id}", controllers.UpdateProject).Methods("PUT")

	// @Summary Delete project
	// @Description Delete project by ID
	// @Tags Projects
	// @Param id path int true "Project ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Project not found"
	// @Router /project/{id} [delete]
	projectRouter.HandleFunc("/{id}", controllers.DeleteProject).Methods("DELETE")

	// Admin routes
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.JWTVerify)

	// @Summary Create a new admin
	// @Description Create a new admin
	// @Tags Admins
	// @Accept json
	// @Produce json
	// @Param admin body models.Admin true "Admin information"
	// @Success 201 {object} models.Admin
	// @Failure 400 {string} string "Bad Request"
	// @Router /admin [post]
	adminRouter.HandleFunc("", controllers.CreateAdmin).Methods("POST")

	// @Summary Get all admins
	// @Description Retrieve all admins
	// @Tags Admins
	// @Produce json
	// @Success 200 {array} models.Admin
	// @Failure 500 {string} string "Internal Server Error"
	// @Router /admin [get]
	adminRouter.HandleFunc("", controllers.GetAdmin).Methods("GET")

	// @Summary Get admin by ID
	// @Description Retrieve admin by ID
	// @Tags Admins
	// @Produce json
	// @Param id path int true "Admin ID"
	// @Success 200 {object} models.Admin
	// @Failure 404 {string} string "Admin not found"
	// @Router /admin/{id} [get]
	adminRouter.HandleFunc("/{id}", controllers.GetAdminID).Methods("GET")

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
	// @Router /admin/{id} [put]
	adminRouter.HandleFunc("/{id}", controllers.UpdateAdmin).Methods("PUT")

	// @Summary Delete admin
	// @Description Delete admin by ID
	// @Tags Admins
	// @Param id path int true "Admin ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Admin not found"
	// @Router /admin/{id} [delete]
	adminRouter.HandleFunc("/{id}", controllers.DeleteAdmin).Methods("DELETE")

}
