package docs

import (
	"bytes"
	"strings"
	"encoding/hex"
	"encoding/json"
	"strconv"
)

type PlaceInput struct {
	Continent     string
	Ocean         string
	Union         string
	Country       string
	CountyOrState string
	District      string
	TownOrCity    string
	Borough       string
	Road          string
	Building      string
	Apartment     string
}

func ExamplePlace() PlaceInput {
	return PlaceInput{
		Continent:     "EUROPE",
		Union:         "UNITED KINGDOM",
		Country:       "ENGLAND",
		CountyOrState: "KENT",
		District:      "THANET",
		TownOrCity:    "MARGATE",
		Borough:       "CLIFTONVILLE",
		Road:          "DALBY SQUARE",
		Building:      "14",
	}
}

type Place struct {
	ID      string
	Details []string
}

func (place *Place) ToID() string {
	b := Hash(place.ToJSON())
	return hex.EncodeToString(b)
}

func (place *Place) ToJSON() []byte {
	b, err := json.Marshal(place.Details)
	if err != nil {
		panic("JSON place")
	}
	return b
}

func (place *Place) ParentHashes() []string {
	hashes := []string{}
	var currentHash string
	for n, v := range place.Details {

		currentHash = hex.EncodeToString(
			Hash(
				bytes.Join(
					[][]byte{
						[]byte(currentHash),
						[]byte(Hash([]byte(strconv.Itoa(n) + " " + v))),
					},
					nil,
				),
			),
		)
		hashes = append(hashes, currentHash)
	}
	return hashes
}

func NewPlace(input PlaceInput) Place {

	values := []string{
		input.Continent,
		input.Ocean,
		input.Union,
		input.Country,
		input.CountyOrState,
		input.District,
		input.TownOrCity,
		input.Borough,
		input.Road,
		input.Building,
		input.Apartment,
	}

	place := Place{
		Details: make([]string, len(values)),
	}
	for n, v := range values {
		if len(v) > 0 {
			place.Details[n] = v
		}
		/*
			if place.Details[n] != nil {
				println(n, *place.Details[n])
			}
		*/
	}

	place.ID = place.ToID()

	return place
}

func (self Place) URI() string {
	hashes := self.ParentHashes()
	return "where/" + strings.Join(hashes, "/")
}