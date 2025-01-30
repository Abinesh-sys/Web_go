package controllers

import (
	"encoding/json"
	"net/http"
	"webgolang/models"
	"github.com/gorilla/mux"
	"strconv"
)

// FeatureController struct to handle feature API requests.
type FeatureController struct{}

// GetFeatures handles GET requests to fetch all features.
func (fc *FeatureController) GetFeatures(w http.ResponseWriter, r *http.Request) {
	features, err := models.GetFeatures()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(features); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetFeature handles GET requests to fetch a single feature by ID.
func (fc *FeatureController) GetFeature(w http.ResponseWriter, r *http.Request) {
	// Get the feature ID from the URL
	idStr := mux.Vars(r)["id"] // Get the ID as a string

	// Convert the string ID to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	feature, err := models.GetFeature(id)
	if err != nil {
		http.Error(w, "Feature not found", http.StatusNotFound)
		return
	}

	// Convert to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateFeature handles POST requests to create a new feature.
func (fc *FeatureController) CreateFeature(w http.ResponseWriter, r *http.Request) {
	var feature models.Feature
	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&feature); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the feature into the database
	if err := models.CreateFeature(feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the created feature back as the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(feature); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



