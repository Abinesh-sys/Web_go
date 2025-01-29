package controllers

import (
	"net/http"
)

func CreateRole(w http.ResponseWriter, r *http.Request) {
	// Implement Role Creation Logic Here
	w.Write([]byte("Role Created"))
}
