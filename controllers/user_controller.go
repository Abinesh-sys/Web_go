package controllers

import (
	"encoding/json"
	"net/http"
	"webgolang/models"
	"strconv"
)

// GetUser handles retrieving a user by their ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from query parameter
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert user ID from string to int
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user details from the database
	user, err := models.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with user details in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser handles user creation along with the profile photo upload
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data (for file upload)
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve form data
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	roleID := r.FormValue("role_id")

	// Convert roleID to integer
	roleIDInt, err := strconv.Atoi(roleID)
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	// Retrieve the profile photo from the form (if any)
	profilePhoto, _, err := r.FormFile("profile_photo")
	if err != nil && err.Error() != "http: no such file" {
		http.Error(w, "Error reading profile photo", http.StatusBadRequest)
		return
	}

	// Create a user object
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		RoleID:   roleIDInt,
	}

	// Save the user and profile photo (if provided)
	err = models.CreateUser(user, profilePhoto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}





