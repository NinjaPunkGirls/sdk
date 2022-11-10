package graph

import (
	"context"
	"errors"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
)

type GraphClient struct {
	firestoreClient *firestore.Client
	nodeCollection  *firestore.DocumentRef
	edgeCollection  *firestore.DocumentRef
}

func NewClient(f *firestore.Client, dbName string) *GraphClient {
	return &GraphClient{
		firestoreClient: f,
		edgeCollection:  f.Collection("_edges").Doc(dbName),
		nodeCollection:  f.Collection("_nodes").Doc(dbName),
	}
}

func (client *GraphClient) SplitID(id string) (string, string, error) {
	s := strings.Split(id, "_")
	if len(s) != 2 {
		return "", "", errors.New("malformed global ID: " + id)
	}
	return s[0], s[1], nil
}

func (client *GraphClient) NewNode(node *Node) (*Node, error) {

	for key, array := range node.Auto {
		filter := map[string]bool{}
		for _, item := range array {
			filter[item] = true
		}
		node.Auto[key] = []string{}
		for item, _ := range filter {
			node.Auto[key] = append(node.Auto[key], item)
		}
	}

	node.Time = time.Now().UTC().Unix()

	if _, err := client.nodeCollection.Collection(node.Class).Doc(node.ID).Set(context.Background(), node); err != nil {
		return nil, err
	}

	if len(node.Parent) == 2 {
		parentClass := node.Parent[0]
		parent := node.Parent[1]

		ps := &PredicateStat{}
		if doc, err := client.nodeCollection.Collection(parentClass).Doc(parent).Collection("p").Doc(node.Class).Get(context.Background()); err == nil {
			if err := doc.DataTo(ps); err == nil {
				return nil, err
			}
		}
		// increment the counter for this predicate
		ps.Value++
		ps.Key = node.Class
		// update the stat document
		if _, err := client.nodeCollection.Collection(parentClass).Doc(parent).Collection("p").Doc(node.Class).Set(context.Background(), ps); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func (client *GraphClient) GetNode(globalID string) (*Node, error) {
	class, id, err := client.SplitID(globalID)
	if err != nil {
		return nil, err
	}
	doc, err := client.nodeCollection.Collection(class).Doc(id).Get(context.Background())
	if err != nil {
		return nil, err
	}
	node := &Node{}
	if err := doc.DataTo(node); err != nil {
		return nil, err
	}
	node.client = client
	return node, nil
}
