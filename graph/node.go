package graph

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"
)

type Node struct {
	client *GraphClient
	ID     string
	Class  string
	// parent is the ID but a slice after split with "_"
	Parent []string    `firestore:"Parent,omitempty"`
	Data   interface{} `firestore:"Data,omitempty"`
	Auto   []string    `firestore:"Auto,omitempty"`
	Time   int64
}

func (node *Node) Global() string {
	return fmt.Sprintf("%s_%s", node.Class, node.ID)
}

func (node *Node) In(predicate string) ([]string, error) {
	return node.traverse("O", predicate)
}
func (node *Node) Out(predicate string) ([]string, error) {
	return node.traverse("I", predicate)
}

func (node *Node) traverse(direction, predicate string) ([]string, error) {

	var opposite string = "I"
	if direction == "I" {
		opposite = "O"
	}

	results := []string{}

	println(predicate, direction, "==", node.ID, opposite)

	iter := node.client.edgeCollection.Collection(predicate).Where(direction, "==", node.Global()).Select(opposite).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		m := map[string]interface{}{}
		if err := doc.DataTo(&m); err != nil {
			return nil, err
		}
		results = append(
			results,
			m[opposite].(string),
		)
	}

	return results, nil
}
