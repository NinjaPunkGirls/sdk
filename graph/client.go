package graph

import (
	"context"
	"encoding/json"
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

func (client *GraphClient) NewNode(class, id string, data interface{}) (*Node, error) {

	var payload map[string]interface{}

	switch v := data.(type) {
	case map[string]interface{}:
		payload = v
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, &payload); err != nil {
			return nil, err
		}
	}

	var keys, values []string
	for key, value := range payload {
		b, err := json.Marshal(value)
		if err != nil {
			panic(err)
		}
		keys = append(keys, key)
		values = append(values, string(b))
	}
	autoKeys := []string{}
	for _, value := range payload {
		switch v := value.(type) {
		case string:
			for x := 0; x < 10; x++ {
				if x > 4 && x < len(v) {
					autoKeys = append(autoKeys, v[:x])
				}
			}
		}
	}

	node := &Node{
		ID:     id,
		Class:  class,
		Keys:   keys,
		Values: values,
		Auto:   autoKeys,
		Time:   time.Now().UTC().Unix(),
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
