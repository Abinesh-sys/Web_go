package models

import (
    "webgolang/database"
    "errors"
)

// User Struct
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    RoleID   int    `json:"role_id"`
}

// GetUserByID fetches a user from the database
func GetUserByID(id int) (*User, error) {
    row := database.DB.QueryRow("SELECT id, name, email, role_id FROM users WHERE id=$1", id)

    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.RoleID)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user User) error {
    query := `INSERT INTO users (name, email, password, role_id) VALUES ($1, $2, $3, $4)`

    _, err := database.DB.Exec(query, user.Name, user.Email, user.Password, user.RoleID)
    if err != nil {
        return errors.New("failed to insert user into the database")
    }
    return nil
}
