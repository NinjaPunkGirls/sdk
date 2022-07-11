package graph

import (
	"context"
	"time"
)

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
