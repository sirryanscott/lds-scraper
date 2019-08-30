package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"lds-scraper/server/utilities"
	"net/http"
	"os"
)

// User has a username and a session cookie
type User struct {
	Username string
	Session  *http.Cookie
}

// SaveUser marshals the user and saves to users.json
func SaveUser(user User) error {
	f, err := os.OpenFile(utilities.UserFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadFile(utilities.UserFile)
	if err != nil {
		return err
	}

	f, err = os.Create(utilities.UserFile)
	if err != nil {
		return err
	}

	users, err := GetUsers(data)
	if err != nil && err.Error() != "No users" {
		return err
	}

	users = UpdateUsers(users, user)

	js, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	if _, err = f.WriteString(string(js)); err != nil {
		return err
	}
	return nil
}

// GetUsers returns all of the users in users.json
func GetUsers(data []byte) ([]User, error) {
	var users []User
	err := json.Unmarshal(data, &users)

	if len(users) == 0 {
		return users, errors.New("No users")
	}

	if err != nil {
		return []User{}, err
	}

	return users, nil
}

// GetUser retrieves the user from users.json and unmarshals in
func GetUser(username string) (User, error) {
	js, err := ioutil.ReadFile(utilities.UserFile)
	if err != nil {
		return User{}, err
	}

	var users []User
	err = json.Unmarshal(js, &users)
	if err != nil {
		return User{}, err
	}

	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

//UpdateUsers updates the saved users and adds if needed
func UpdateUsers(users []User, user User) []User {
	var nUsers []User

	exists := false
	for _, u := range users {
		if u.Username == user.Username {
			nUsers = append(nUsers, user)
			exists = true
		} else {
			nUsers = append(nUsers, u)
		}
	}

	if !exists {
		nUsers = append(nUsers, user)
	}
	return nUsers

}
