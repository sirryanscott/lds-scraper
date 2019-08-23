package handler

import (
	"lds-scraper/controller"
	"net/http"
)

// Login is the handler for logging in a user to churchofjesuschrist.org
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := controller.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetMembers is the handler for getting the member list from LCR
func GetMembers(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	mList, err := controller.GetMembers(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(mList)
}
