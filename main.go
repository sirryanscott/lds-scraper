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
	login("sirryanscott", "maVoitan1")
	getMembers()
}

func login(username, password string) {
	formData.Add("username", username)
	formData.Add("password", password)
	fmt.Println(formData)
	session = getSession()
}

func getSession() *http.Cookie {
	resp, err := http.PostForm(loginurl, formData)
	if err != nil {
		log.Fatal(err)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "ObSSOCookie" {
			fmt.Println(cookie)
			return cookie
		}
	}
	return nil

}

func getMembers() {
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

	ioutil.WriteFile(filenameMembers, body, 0600)
}
