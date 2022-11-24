package docs

import (
	"crypto/sha1"
	"time"
	"fmt"
	"context"

	"cloud.google.com/go/storage"
)

func Hash(b []byte) []byte {
	h := sha1.New()
	h.Write(b)	
	return h.Sum(nil)
}

type Client struct {
	Storage *storage.Client
}

func NewClient() *Client {

	client, err := storage.NewClient(context.Background())
	if err != nil {
		panic(err)
	}
	return &Client{
		client,
	}
}

func (client *Client) EmptyDocument() *Document {
	doc := &Document{
		client: client,
	}
	return doc
}

func (client *Client) NewDocument(where Place, class string, data interface{}) *Document {
	doc := &Document{
		client: client,
		Class: class,
		Time:  fmt.Sprintf("%d", time.Now().UTC().Unix()),
		Place: where,
		Data: data,
	}
	return doc
}


/*
func (doc *Document) URI() string {
	values := []interface{}{
		doc.Time.Raw,
		doc.Place.Continent,
	}

	// eg: Africa || Antarctica
	if len(doc.Place.Continent) > 0 {
		values = append(values, doc.Place.Continent)
	}

	// eg: Pacific || Indian
	if len(doc.Place.Ocean) > 0 {
		values = append(values, doc.Place.Ocean)
	}

	// eg: United States of America || United Kingdom
	if len(doc.Place.Union) > 0 {
		values = append(values, doc.Place.Union)
	}

	// eg: Wales || Cambodia
	if len(doc.Place.Country) > 0 {
		values = append(values, doc.Place.Country)
	}

	// eg: Greater London || Texas
	if len(doc.Place.CountyOrState) > 0 {
		values = append(values, doc.Place.CountyOrState)
	}

	// eg: Greater London || Texas
	if len(doc.Place.CountyOrState) > 0 {
		values = append(values, doc.Place.CountyOrState)
	}

	// eg: Greater London || Texas
	if len(doc.Place.District) > 0 {
		values = append(values, doc.Place.District)
	}

	// eg: London || St. Ives
	if len(doc.Place.TownOrCity) > 0 {
		values = append(values, doc.Place.TownOrCity)
	}

	// eg: Newham || Thanet
	if len(doc.Place.Borough) > 0 {
		values = append(values, doc.Place.Borough)
	}

	// eg: Horseguards || Canada Square
	if len(doc.Place.Road) > 0 {
		values = append(values, doc.Place.Road)
	}

	// eg: 39 || Royal Opera House
	if len(doc.Place.Building) > 0 {
		values = append(values, doc.Place.Building)
	}

	return fmt.Sprintf(
		"%v_%v",
		values...,
	)
}
*/