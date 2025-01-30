package models

import (
    "errors"
    "webgolang/database"
    "mime/multipart"
    "os"
    "strconv"
	"path/filepath"
	"io"
)

// User Struct
type User struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    RoleID      int    `json:"role_id"`
    ProfilePhoto string `json:"profile_photo"` // Added profile photo field
}

// GetUserByID fetches a user from the database
func GetUserByID(id int) (*User, error) {
    row := database.DB.QueryRow("SELECT id, name, email, role_id, profile_photo FROM users WHERE id=$1", id)

    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.RoleID, &user.ProfilePhoto)
    if err != nil {
        return nil, err
    }
    return &user, nil
}



func CreateUser(user User, profilePhoto multipart.File) error {
    var photoFileName string

    // If there is a profile photo, save it to the file system
    if profilePhoto != nil {
        // Read the content of the profile photo into a byte slice
        photoBytes, err := io.ReadAll(profilePhoto)
        if err != nil {
            return errors.New("failed to read profile photo")
        }

        // Generate a file name for the photo (using the user's ID or a unique name)
        photoFileName = filepath.Join("uploads", user.Name+"_"+strconv.Itoa(user.ID)+".jpg")

        // Create the file in the 'uploads' directory
        file, err := os.Create(photoFileName)
        if err != nil {
            return errors.New("failed to save profile photo")
        }
        defer file.Close()

        // Write the byte slice to the file
        _, err = file.Write(photoBytes)
        if err != nil {
            return errors.New("failed to write profile photo")
        }

        // Save the file path in the user's ProfilePhoto field
        user.ProfilePhoto = photoFileName
    }

    // Insert user data into the database including the profile photo path
    query := `INSERT INTO users (name, email, password, role_id, profile_photo) 
            VALUES ($1, $2, $3, $4, $5)`

    _, err := database.DB.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.ProfilePhoto)
    if err != nil {
        return errors.New("failed to insert user into the database")
    }

    return nil
}





