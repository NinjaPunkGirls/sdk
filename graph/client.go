package graph

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type GraphClient struct {
	firestoreClient                *firestore.Client
	nodeCollection, edgeCollection *firestore.CollectionRef
}

func NewClient(f *firestore.Client, dbName string) *GraphClient {
	return &GraphClient{
		firestoreClient: f,
		edgeCollection:  f.Collection("_graph").Doc(dbName).Collection("_edges"),
		nodeCollection:  f.Collection("_graph").Doc(dbName).Collection("_nodes"),
	}
}

type Node struct {
	client *GraphClient
	ID     string
	Data   map[string]interface{}
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

type Edge struct {
	I, O string
	P    string
	X    interface{}
	T    int64
}

func (client *GraphClient) GetNode(id string) (*Node, error) {
	doc, err := client.nodeCollection.Doc(id).Get(context.Background())
	if err != nil {
		return nil, err
	}
	node := &Node{}
	return node, doc.DataTo(node)
}

func (client *GraphClient) LinkNodes(in, out string, predicate string, data ...map[string]interface{}) error {

	edge := &Edge{
		I: in,
		O: out,
		P: predicate,
		X: data,
		T: time.Now().UTC().Unix(),
	}

	_, err := client.edgeCollection.NewDoc().Set(context.Background(), edge)
	return err
}
