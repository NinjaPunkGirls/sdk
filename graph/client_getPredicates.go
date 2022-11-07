package graph

import (
	"context"

	"google.golang.org/api/iterator"
)

func (client *GraphClient) GetPredicates(globalID string) ([]*PredicateStat, error) {

	class, id, err := client.SplitID(globalID)
	if err != nil {
		return nil, err
	}

	results := []*PredicateStat{}

	iter := client.nodeCollection.Collection(class).Doc(id).Collection("p").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		ps := &PredicateStat{}
		if err := doc.DataTo(ps); err != nil {
			return nil, err
		}
		ps.Key = id
		results = append(
			results,
			ps,
		)
	}

	return results, nil
}
