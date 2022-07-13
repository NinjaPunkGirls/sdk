package graph

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/iterator"
)

type Node struct {
	client *GraphClient
	ID     string
	Class  string
	Keys   []string
	Values []string
	Auto   []string
	Time   int64
}

func (node *Node) Global() string {
	return fmt.Sprintf("%s_%s", node.Class, node.ID)
}

func (node *Node) In(predicate string) ([]*Node, error) {
	return node.traverse("O", predicate)
}
func (node *Node) Out(predicate string) ([]*Node, error) {
	return node.traverse("I", predicate)
}

func (node *Node) traverse(direction, predicate string) ([]*Node, error) {

	results := []*Node{}

	iter := node.client.edgeCollection.Collection(predicate).Where(direction, "==", node.ID).Documents(context.Background())
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
