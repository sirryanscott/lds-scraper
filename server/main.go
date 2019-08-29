package main

import (
	"log"
	"net/http"
	"ryan/lds-scraper/server/handler"
)

func main() {
	http.HandleFunc("/login/", handler.Login)
	http.HandleFunc("/members/", handler.GetMembers)
	http.HandleFunc("/export/", handler.Login) //FIXME: not the right handler

	log.Println("****Starting Server****")
	log.Fatal(http.ListenAndServe(":7070", nil))
}
