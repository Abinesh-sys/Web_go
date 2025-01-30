package models

import (
	"webgolang/database"
)

// Feature struct represents a feature entity in the system.
type Feature struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// GetFeatures retrieves all features from the database.
func GetFeatures() ([]Feature, error) {
	rows, err := database.DB.Query("SELECT id, name, description, is_active FROM features")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var features []Feature
	for rows.Next() {
		var feature Feature
		if err := rows.Scan(&feature.ID, &feature.Name, &feature.Description, &feature.IsActive); err != nil {
			return nil, err
		}
		features = append(features, feature)
	}
	return features, nil
}

// CreateFeature inserts a new feature into the database.
func CreateFeature(feature Feature) error {
	query := "INSERT INTO features (name, description, is_active) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, feature.Name, feature.Description, feature.IsActive)
	return err
}

// GetFeature retrieves a feature by its ID from the database.
func GetFeature(id int) (*Feature, error) {
	row := database.DB.QueryRow("SELECT id, name, description, is_active FROM features WHERE id = $1", id)

	var feature Feature
	if err := row.Scan(&feature.ID, &feature.Name, &feature.Description, &feature.IsActive); err != nil {
		return nil, err
	}
	return &feature, nil
}


