package main

import (
	"lds-scraper/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login/", handler.Login)
	http.HandleFunc("/members/", handler.GetMembers)
	http.HandleFunc("/export/", handler.Login) //FIXME: not the right handler

	log.Println("****Starting Server****")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
