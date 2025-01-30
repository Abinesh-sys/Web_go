package models

import (
	"webgolang/database"
	
)

type Feature struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

// CreateFeature inserts a new feature into the database
func CreateFeature(feature Feature) error {
	query := "INSERT INTO features (name, description) VALUES ($1, $2) RETURNING id"
	err := database.DB.QueryRow(query, feature.Name, feature.Description).Scan(&feature.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllFeatures retrieves all features from the database
func GetAllFeatures() ([]Feature, error) {
	rows, err := database.DB.Query("SELECT id, name, description FROM features")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var features []Feature
	for rows.Next() {
		var feature Feature
		if err := rows.Scan(&feature.ID, &feature.Name, &feature.Description); err != nil {
			return nil, err
		}
		features = append(features, feature)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return features, nil
}

// GetFeature retrieves a single feature by ID
func GetFeature(id int) (Feature, error) {
	var feature Feature
	err := database.DB.QueryRow("SELECT id, name, description FROM features WHERE id = $1", id).Scan(&feature.ID, &feature.Name, &feature.Description)
	if err != nil {
		return feature, err
	}
	return feature, nil
}



