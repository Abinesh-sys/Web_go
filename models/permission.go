package models

import (
	"webgolang/database"
)

type Permission struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FeatureID int    `json:"feature_id"`
	RoleID    int    `json:"role_id"`
}

func GetPermissions() ([]Permission, error) {
	rows, err := database.DB.Query("SELECT id, name, feature_id, role_id FROM permissions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []Permission
	for rows.Next() {
		var permission Permission
		if err := rows.Scan(&permission.ID, &permission.Name, &permission.FeatureID, &permission.RoleID); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func CreatePermission(permission Permission) error {
	query := "INSERT INTO permissions (name, feature_id, role_id) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, permission.Name, permission.FeatureID, permission.RoleID)
	return err
}
