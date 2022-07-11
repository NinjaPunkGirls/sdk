package graph

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (client *GraphClient) GetNodes(class string) ([]*Node, error) {

	results := []*Node{}

	iter := client.nodeCollection.Collection(class).OrderBy("Time", firestore.Desc).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		node := &Node{}
		if err := doc.DataTo(node); err != nil {
			log.Println(err)
			continue
		}
		results = append(
			results,
			node,
		)
	}

	return results, nil
}
