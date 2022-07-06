package graph

import (
	"context"
	"log"

	"google.golang.org/api/iterator"
)

type Node struct {
	client *GraphClient
	ID     string
	Type   string
	Data   interface{}
	Time   int64
}

func (node *Node) In(predicate string) ([]*Node, error) {
	return node.traverse("O", predicate)
}
func (node *Node) Out(predicate string) ([]*Node, error) {
	return node.traverse("I", predicate)
}

func (node *Node) traverse(direction, predicate string) ([]*Node, error) {

	results := []*Node{}

	iter := node.client.edgeCollection.Where("I", "==", node.ID).Where("P", "==", predicate).Documents(context.Background())
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
