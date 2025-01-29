package controllers

import (
	"net/http"
)

func ManageFeatures(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Manage Features"))
}
