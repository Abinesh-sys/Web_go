package controllers

import (
	"net/http"
)

func ManagePermissions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Manage Permissions"))
}
