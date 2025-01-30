package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"webgolang/controllers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Serve Static Files
	r.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	// User Routes
	r.HandleFunc("/users", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")

	// Role Routes
	r.HandleFunc("/roles/{id}", controllers.GetRoleByIDHandler).Methods("GET")
	r.HandleFunc("/roles", controllers.CreateRoleHandler).Methods("POST")

	// Permission Routes
	r.HandleFunc("/permissions", controllers.GetPermissionsHandler).Methods("GET")
	r.HandleFunc("/permissions", controllers.CreatePermissionHandler).Methods("POST")

	// Serve Home Page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	}).Methods("GET")

	return r
}




