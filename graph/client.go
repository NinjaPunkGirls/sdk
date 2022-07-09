package graph

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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

func (client *GraphClient) NewNode(class, id string, data interface{}) (*Node, error) {
	node := &Node{
		ID:    id,
		Class: class,
		Data:  data,
		Time:  time.Now().UTC().Unix(),
	}
	_, err := client.nodeCollection.Collection(class).Doc(id).Set(context.Background(), node)
	return node, err
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
	return node, doc.DataTo(node)
}

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

func (client *GraphClient) LinkNodes(in, out string, predicate string, data ...map[string]interface{}) error {

	edge := &Edge{
		I: in,
		O: out,
		P: predicate,
		X: data,
		T: time.Now().UTC().Unix(),
	}

	_, err := client.edgeCollection.Collection(predicate).NewDoc().Set(context.Background(), edge)
	return err
}
