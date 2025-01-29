package routes

import (
	"webgolang/controllers"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/user", controllers.GetUser)
	router.HandleFunc("/create_role", controllers.CreateRole)
	router.HandleFunc("/manage_permissions", controllers.ManagePermissions)
	router.HandleFunc("/manage_features", controllers.ManageFeatures)
	return router
}
