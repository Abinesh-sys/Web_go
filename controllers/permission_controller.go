package controllers

import (
	"encoding/json"
	"net/http"
	"webgolang/models"
)

func GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	permissions, err := models.GetPermissions()
	if err != nil {
		http.Error(w, "Error fetching permissions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

func CreatePermissionHandler(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := models.CreatePermission(permission); err != nil {
		http.Error(w, "Error creating permission", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Permission created successfully"})
}

