package controller

import (
	"encoding/json"
	"io/ioutil"
	"lds-scraper/server/models"
	"lds-scraper/server/utilities"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
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

// CalculateFee calculates the fee if distance is outside 20 mile radius
func CalculateFee(destination string) ([]byte, error) {
	destination = strings.ReplaceAll(destination, " ", "+")
	log.Println(destination)

	url := utilities.GoogleMapsBeginURL + destination + utilities.GoogleMapsEndURL
	log.Println(url)

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

	var distance models.DistanceMatrix
	err = json.Unmarshal(body, &distance)

	if err != nil {
		log.Println("error in json decoding")
	}

	dist := float64(distance.Rows[0].Elements[0].Distance.Value)

	var kmToMi = 1609.0 // km to miles
	var minMiles = 10.0
	var travelFee = 2.5
	dist /= kmToMi
	log.Print(strconv.FormatFloat(dist, 'f', 0, 64) + " miles total")

	dist -= minMiles
	fee := math.Ceil(dist * travelFee)

	log.Print(strconv.FormatFloat(dist, 'f', 0, 64) + " miles used for calculation")

	if fee < 0 {
		fee = 0
	}

	log.Print("fee: $" + strconv.FormatFloat(fee, 'f', 0, 64))

	return []byte(strconv.FormatFloat(fee, 'f', 0, 64)), nil
}
