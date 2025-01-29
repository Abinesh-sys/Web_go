package models

type User  struct {
	ID     int
	Name   string
	Email  string
	RoleID int
	Role   Role
}

func GetUserByID(id int) (*User, error) {

}

func CreateUser(user User) error {
	
}