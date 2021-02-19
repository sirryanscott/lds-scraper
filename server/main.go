package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lds-scraper/server/handler"
	"lds-scraper/server/utilities"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/login/", handler.Login)
	http.HandleFunc("/members/", handler.GetMembers)
	http.HandleFunc("/export/", handler.Login)        //FIXME: not the right handler
	http.HandleFunc("/travelfee/", handler.TravelFee) // handler for MyMiniPonyParty travel fee calculator

	//loadWardList()
	//printSingles()

	log.Println("****Starting Server****")
	log.Fatal(http.ListenAndServe(":7070", nil))
}

func loadWardList() error {
	type HeadOfHousehold struct {
		NameFormats struct {
			Name string `json:"listPreferredLocal"`
		} `json:"nameFormats"`
		HouseholdMember struct {
			Household struct {
				Address struct {
					AddressLines []string `json:"addressLines"`
				} `json:"address"`
				HeadName string `json:"directoryPreferredLocal"`
			} `json:"household"`
		} `json:"householdMember"`
		IsHead                   bool   `json:"isHead"`
		PriesthoodOffice         string `json:"priesthoodOffice"`
		PriesthoodTeacherOrAbove bool   `json:"priesthoodTeacherOrAbove"`
	}

	js, err := ioutil.ReadFile(utilities.Households)
	if err != nil {
		log.Println("error in file reading")
		return err
	}

	var allHouseholds []HeadOfHousehold
	var foxwood1 []HeadOfHousehold
	var foxwood1Priesthood []HeadOfHousehold
	var foxwood2 []HeadOfHousehold
	var foxwood2Priesthood []HeadOfHousehold

	err = json.Unmarshal(js, &allHouseholds)
	if err != nil {
		log.Println("error in unmarshal")
		return err
	}

	for _, head := range allHouseholds {
		isOldWard := true
		streetAddress := head.HouseholdMember.Household.Address.AddressLines[0]
		for _, road := range utilities.Roads {
			if strings.Contains(strings.ToLower(streetAddress), road) {
				if head.IsHead {
					foxwood1 = append(foxwood1, head)
				}
				if head.PriesthoodTeacherOrAbove {
					foxwood1Priesthood = append(foxwood1Priesthood, head)
				}
				isOldWard = false
			}

		}
		if isOldWard {
			if head.IsHead {
				foxwood2 = append(foxwood2, head)
			}
			if head.PriesthoodTeacherOrAbove {
				foxwood2Priesthood = append(foxwood2Priesthood, head)
			}
		}
	}

	for _, home := range foxwood1 {
		fmt.Println(home.HouseholdMember.Household.HeadName)
		//log.Println(home.HouseholdMember.Household.Address.AddressLines[0])
	}

	fmt.Println("Foxwood 1 count: " + strconv.Itoa(len(foxwood1)))
	fmt.Println()
	for _, home := range foxwood2 {
		fmt.Println(home.HouseholdMember.Household.HeadName)
		//log.Println(home.HouseholdMember.Household.Address.AddressLines[0])
	}
	fmt.Println("Foxwood 2 count: " + strconv.Itoa(len(foxwood2)))
	fmt.Println()

	totalHouseholds := len(foxwood1) + len(foxwood2)
	fmt.Println("Total Households: " + strconv.Itoa(totalHouseholds))

	fmt.Println("\nFoxwood 1 Priesthood")

	for _, home := range foxwood1Priesthood {
		fmt.Println(home.NameFormats.Name + "\t" + home.PriesthoodOffice)
	}

	fmt.Println("\nFoxwood 2 Priesthood")

	for _, home := range foxwood2Priesthood {
		fmt.Println(home.NameFormats.Name + "\t" + home.PriesthoodOffice)
	}

	return nil
}

func printSingles() error {
	type HeadOfHousehold struct {
		NameFormats struct {
			Name string `json:"listPreferredLocal"`
		} `json:"nameFormats"`
		HouseholdMember struct {
			Household struct {
				Address struct {
					AddressLines []string `json:"addressLines"`
				} `json:"address"`
				HeadName string `json:"directoryPreferredLocal"`
			} `json:"household"`
		} `json:"householdMember"`
		IsHead                   bool   `json:"isHead"`
		PriesthoodOffice         string `json:"priesthoodOffice"`
		PriesthoodTeacherOrAbove bool   `json:"priesthoodTeacherOrAbove"`
		IsSingleAdult            bool   `json:"isSingleAdult"`
	}

	js, err := ioutil.ReadFile(utilities.Households)
	if err != nil {
		log.Println("error in file reading")
		return err
	}

	var allHouseholds []HeadOfHousehold

	err = json.Unmarshal(js, &allHouseholds)
	if err != nil {
		log.Println("error in unmarshal")
		return err
	}

	for _, head := range allHouseholds {
		if head.IsSingleAdult {
			fmt.Print(head.NameFormats.Name + "\t\t\t" + head.HouseholdMember.Household.Address.AddressLines[0])
			if head.PriesthoodTeacherOrAbove {
				fmt.Print("\t\t\t" + head.PriesthoodOffice + "\n")
			} else {
				fmt.Print("\n")
			}
		}
	}

	return nil
}
