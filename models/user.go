package models

import (
	"webgolang/database"
)

type User  struct {
	ID     int
	Name   string
	Email  string
	RoleID int
}

func GetUserByID(id int) (*User, error) {
	row := database.DB.QueryRow("SELECT id, name, email, role_id FROM users WHERE id=$1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.RoleID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}