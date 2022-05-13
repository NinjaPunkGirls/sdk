package npgsdk

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"io"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"github.com/golangdaddy/relysia-client"
)

// InitSDK constructs the context for the user's interactions
func InitSDK(gcpProjectID, authToken string) *SDK {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: gcpProjectID}
	fapp, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	storageClient, err := fapp.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &SDK{
		Relysia: relysia.NewClient().WithToken(authToken),
		Storage: storageClient,
	}
}

type SDK struct {
	Relysia      *relysia.Client
	Storage      *storage.Client
	insecureMode bool
}

func (self *SDK) Insecure() {
	self.insecureMode = true
	self.Relysia.Insecure()
}

func (sdk *SDK) HashSHA1(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	return h.Sum(nil)
}

func (sdk *SDK) GetHTTP(url string, dst interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}
	return nil
}
