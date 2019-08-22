package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	session         *http.Cookie
	loginurl        = "https://signin.churchofjesuschrist.org/login.html"
	memberList      = "https://lcr.churchofjesuschrist.org/services/umlu/report/member-list"
	formData        = url.Values{}
	filenameMembers = "member-list"
)

func main() {
	http.HandleFunc("/login/", login)
	http.HandleFunc("/members/", getMembers)
	http.HandleFunc("/export/", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	formData.Add("username", username)
	formData.Add("password", password)

	fmt.Println(formData)
	session = getSession()
}

func getSession() *http.Cookie {
	resp, err := http.PostForm(loginurl, formData)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "ObSSOCookie" {
			fmt.Println(cookie)
			return cookie
		}
	}
	return nil

}

func getMembers(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	req, err := http.NewRequest("GET", memberList, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(session)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
