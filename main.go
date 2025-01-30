package main

import (
	"webgolang/database"
	"webgolang/routes"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	router := routes.RegisterRoutes()

	log.Println("Server started at http://localhost:9090")
	log.Fatal(http.ListenAndServe(":8080", router))
}




