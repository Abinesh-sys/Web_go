package routes

import (
	"webgolang/controllers"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Serve Static Files
	router.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	// User Routes
	router.HandleFunc("/users", controllers.GetUser)
	router.HandleFunc("/user", controllers.CreateUser)

	// Serve Home Page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})

	return router
}

