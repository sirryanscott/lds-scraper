package controller

import (
	"errors"
	"fmt"
	"lds-scraper/utilities"
	"log"
	"net/http"
	"net/url"
)

//Login validates the credentials and gets the session key
func Login(username, password string) error {
	err := validateCredentials(username, password)
	if err != nil {
		return err
	}

	session := getSession(username, password)
	if session == nil {
		return errors.New("Unable to get session")
	}
	fmt.Println(session)

	return nil
}

func validateCredentials(username, password string) error {
	if username == "" || password == "" {
		return errors.New("Username and password cannot be blank")
	}

	return nil
}

func getSession(username, password string) *http.Cookie {
	formData := url.Values{}
	formData.Add("username", username)
	formData.Add("password", password)

	resp, err := http.PostForm(utilities.LoginURL, formData)
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
