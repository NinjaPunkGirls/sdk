package graph

import (
	"context"
	"fmt"
	"time"
)

func (client *GraphClient) LinkNodes(in, out string, predicate string, data ...map[string]interface{}) error {

	if len(in) == 0 || len(out) == 0 {
		return fmt.Errorf("one or more inputs for predicate %s are length zero", predicate)
	}

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
