package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

// CreateProject creates a new project
// @Summary Create a new project
// @Description Create a new project
// @Tags project
// @Accept json
// @Produce json
// @Param project body models.Project true "Project"
// @Success 201 {object} models.Project
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/project [post]
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Project{}) {
		if err := database.DB.AutoMigrate(&models.Project{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&project); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

// GetProject retrieves all projects
// @Summary Get all projects
// @Description Get all projects
// @Tags project
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/project [get]
func GetProject(w http.ResponseWriter, r *http.Request) {
	var project []models.Project
	if result := database.DB.Find(&project); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

// GetProjectID retrieves a project by ID
// @Summary Get project by ID
// @Description Get project by ID
// @Tags project
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} models.Project
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/project/{id} [get]
func GetProjectID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var project models.Project
	if result := database.DB.First(&project, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

// UpdateProject updates a project by ID
// @Summary Update project by ID
// @Description Update project by ID
// @Tags project
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body models.Project true "Project"
// @Success 200 {object} models.Project
// @Failure 400 {string} string "Invalid ID or Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /api/v1/project/{id} [put]
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var project models.Project
	if result := database.DB.First(&project, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&project)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

// DeleteProject deletes a project by ID
// @Summary Delete project by ID
// @Description Delete project by ID
// @Tags project
// @Param id path int true "Project ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/project/{id} [delete]
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.Project{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
