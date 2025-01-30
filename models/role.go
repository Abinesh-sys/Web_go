package models

import (
	"webgolang/database"
)

// Role Struct
type Role struct {
    ID              int    `json:"id"`
    RoleName        string `json:"role_name"`
    RoleIcon        string `json:"role_icon"`
    CanSelfRegister bool   `json:"can_self_register"`
}

func GetRoleByID(id int) (*Role, error) {
    row := database.DB.QueryRow("SELECT id, role_name, role_icon, can_self_register FROM roles WHERE id=$1", id)

    var role Role
    err := row.Scan(&role.ID, &role.RoleName, &role.RoleIcon, &role.CanSelfRegister)
    if err != nil {
        return nil, err
    }
    return &role, nil
}
func CreateRole(role Role) error {
    query := `INSERT INTO roles (role_name, role_icon, can_self_register) VALUES ($1, $2, $3)`
    _, err := database.DB.Exec(query, role.RoleName, role.RoleIcon, role.CanSelfRegister)
    if err != nil {
        return err
    }
    return nil
}


