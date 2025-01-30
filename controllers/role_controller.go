package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webgolang/models"

	"github.com/gorilla/mux"
)

// GetRoleByIDHandler handles fetching a role by ID
func GetRoleByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid role ID", http.StatusBadRequest)
		return
	}

	role, err := models.GetRoleByID(id)
	if err != nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

// CreateRoleHandler handles role creation
func CreateRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := models.CreateRole(role)
	if err != nil {
		http.Error(w, "Failed to create role", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Role created successfully"})
}

