package graph

import (
	"context"

	"google.golang.org/api/iterator"
)

func (client *GraphClient) GetPredicates() ([]string, error) {

	results := []string{}

	iter := client.edgeCollection.Collections(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(
			results,
			doc.ID,
		)
	}

	return results, nil
}
