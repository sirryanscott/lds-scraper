package controller

import (
	"errors"
	"fmt"
	"lds-scraper/server/models"
	"lds-scraper/server/utilities"
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
		fmt.Println("Unable to get session")
		return errors.New("Unable to get session")
	}
	fmt.Println(session)

	user := models.User{Username: username, Session: session}
	err = models.SaveUser(user)
	if err != nil {
		return err
	}

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
	fmt.Println(resp)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	fmt.Println(len(resp.Cookies()))
	for _, cookie := range resp.Cookies() {
		fmt.Println(cookie)
		if cookie.Name == "ObSSOCookie" {
			fmt.Println(cookie)
			return cookie
		}
	}
	fmt.Println("Cookie not found")
	return nil

}
