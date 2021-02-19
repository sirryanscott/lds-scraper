package models

// DistanceMatrix represents the distance from google maps
type DistanceMatrix struct {
	Rows []Rows `json:"rows"`
}

// Rows in json
type Rows struct {
	Elements []Elements `json:"elements"`
}

// Elements in the json
type Elements struct {
	Distance struct {
		Value int `json:"value"`
	} `json:"distance"`
}
