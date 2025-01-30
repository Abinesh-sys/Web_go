package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"webgolang/controllers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Create a new instance of the FeatureController
	featureController := &controllers.FeatureController{}

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

	// Feature routes
	r.HandleFunc("/features", featureController.GetFeatures).Methods("GET")           // Get all features
	r.HandleFunc("/features/{id}", featureController.GetFeature).Methods("GET")      // Get a feature by ID
	r.HandleFunc("/features", featureController.CreateFeature).Methods("POST")      // Create a new feature

	// Serve Home Page
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	}).Methods("GET")

	return r
}




