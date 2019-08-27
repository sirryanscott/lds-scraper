package controller

import (
	"io/ioutil"
	"lds-scraper/server/models"
	"lds-scraper/server/utilities"
	"log"
	"net/http"
)

// GetMembers grabs the member list from LCR for the user
func GetMembers(username string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", utilities.MemberListURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	//FIXME: need to put in a package for saving/retreiving users/session keys
	user, err := models.GetUser(username)
	if err != nil {
		return nil, err
	}
	req.AddCookie(user.Session)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}