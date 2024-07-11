package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	database "github.com/sahilchauhan0603/backend/config"
	models "github.com/sahilchauhan0603/backend/models"
)

func CreateUploader(w http.ResponseWriter, r *http.Request) {
	var uploader models.Uploader
	if err := json.NewDecoder(r.Body).Decode(&uploader); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Uploader{}) {
		if err := database.DB.AutoMigrate(&models.Uploader{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if result := database.DB.Create(&uploader); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uploader)
}

func GetUploader(w http.ResponseWriter, r *http.Request) {

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Uploader{}) {
		if err := database.DB.AutoMigrate(&models.Uploader{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var uploader []models.Uploader
	if result := database.DB.Find(&uploader); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploader)
}


func GetUploaderID(w http.ResponseWriter, r *http.Request) {

	// Check if table exists or create it if it doesn't
	if !database.DB.Migrator().HasTable(&models.Uploader{}) {
		if err := database.DB.AutoMigrate(&models.Uploader{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var uploader models.Uploader
	if result := database.DB.First(&uploader, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploader)
}


func UpdateUploader(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var uploader models.Uploader
	if result := database.DB.First(&uploader, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&uploader); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	database.DB.Save(&uploader)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploader)
}


func DeleteUploader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if result := database.DB.Delete(&models.Uploader{}, id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
