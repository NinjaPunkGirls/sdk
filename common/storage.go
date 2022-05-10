package common

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

func (app *App) GetObject(bucket *storage.BucketHandle, objectName string) ([]byte, error) {
	r, err := bucket.Object(objectName).NewReader(context.Background())
	if err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	if err := r.Close(); err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	return b, nil
}
