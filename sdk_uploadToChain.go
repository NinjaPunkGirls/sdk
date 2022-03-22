package npgsdk

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golangdaddy/relysia-client"
	"github.com/kr/pretty"
)

// UploadToChain makes the given file available on the blockchain via the cloud storage
func (sdk *SDK) UploadToChain(walletID, bucketName, filename string, objectBytes []byte) (*relysia.UploadResponse, error) {

	ctx := context.Background()

	log.Println("FILENAME", filename)

	bucket, err := sdk.Storage.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	obj := bucket.Object(filename)
	w := obj.NewWriter(ctx)
	w.Write(objectBytes)
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	uploadedObjectURI := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, filename)
	log.Println(uploadedObjectURI)

	time.Sleep(time.Second)

	upload, err := sdk.Relysia.UploadReference(
		walletID,
		filename,
		uploadedObjectURI,
		"",
	)
	if err != nil {
		return nil, err
	}
	pretty.Println(upload)

	return upload, nil
}
