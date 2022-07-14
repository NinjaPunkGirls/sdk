package graph

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (client *GraphClient) GetNodes(class string) ([]string, error) {

	results := []string{}

	iter := client.nodeCollection.Collection(class).OrderBy("Time", firestore.Desc).Select().Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var id string = doc.Ref.ID
		results = append(
			results,
			id,
		)
	}

	return results, nil
}
