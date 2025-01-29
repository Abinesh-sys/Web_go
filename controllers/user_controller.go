package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "webgolang/models"
)

// GetUser fetches a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from query parameters
    userID, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Fetch user from database
    user, err := models.GetUserByID(userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Return user data as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// CreateUser handles user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Decode JSON request body into a User struct
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Insert user into the database
    err = models.CreateUser(user)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}


