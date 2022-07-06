package graph

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
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

func (client *GraphClient) NewNode(id, class string, data interface{}) (*Node, error) {
	node := &Node{
		ID:   id,
		Type: class,
		Data: data,
		Time: time.Now().UTC().Unix(),
	}
	_, err := client.nodeCollection.Doc(id).Set(context.Background(), node)
	return node, err
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
