package controllers

import (
	"encoding/json"
	"net/http"
	"webgolang/models"
	"github.com/gorilla/mux"
	"strconv" // Import strconv for string-to-int conversion
)

// FeatureController struct
type FeatureController struct {}

// CreateFeature handles POST requests to create a new feature.
func (fc *FeatureController) CreateFeature(w http.ResponseWriter, r *http.Request) {
	var feature models.Feature

	// Decode the incoming JSON body
	if err := json.NewDecoder(r.Body).Decode(&feature); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert feature into the database
	if err := models.CreateFeature(feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back the created feature as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFeatures handles GET requests to fetch all features.
func (fc *FeatureController) GetFeatures(w http.ResponseWriter, r *http.Request) {
	features, err := models.GetAllFeatures()  // Assuming models.GetAllFeatures() is a function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return features in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(features); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFeature handles GET requests to fetch a single feature by ID.
func (fc *FeatureController) GetFeature(w http.ResponseWriter, r *http.Request) {
	// Get the feature ID from the URL
	idStr := mux.Vars(r)["id"]

	// Convert the string ID to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// Fetch the feature by ID from the database
	feature, err := models.GetFeature(id)
	if err != nil {
		http.Error(w, "Feature not found", http.StatusNotFound)
		return
	}

	// Return the feature as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}








