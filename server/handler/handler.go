package handler

import (
	"lds-scraper/server/controller"
	"log"
	"net/http"
)

// Login is the handler for logging in a user to churchofjesuschrist.org
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging in user")
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	username := r.FormValue("username")
	password := r.FormValue("password")

	err := controller.Login(username, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect credentials"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged in!"))
}

// GetMembers is the handler for getting the member list from LCR
func GetMembers(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting member list")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	username := r.FormValue("username")

	mList, err := controller.GetMembers(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("User doens't have access"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(mList)
}

// TravelFee calculates the fee for travel for myminiponyparty
func TravelFee(w http.ResponseWriter, r *http.Request) {
	log.Println("calculating travel fee for myminiponyparty")
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	keys, ok := r.URL.Query()["dest"]
	if !ok || len(keys[0]) < 1 {
		log.Println("URL param key missing")
		return
	}
	destination := keys[0]

	fee, err := controller.CalculateFee(destination)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in fee calculation"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fee)
}
