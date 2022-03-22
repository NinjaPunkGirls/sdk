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

func InitSDK(gcpProjectID, authToken string) *SDK {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: gcpProjectID}
	fapp, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		firestoreClient, err := fapp.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		defer firestoreClient.Close()
	*/
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
	Relysia *relysia.Client
	Storage *storage.Client
	//	firestoreClient *firestore.Client
}

func (sdk *SDK) HashSHA1(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	return h.Sum(nil)
}

/*
func (sdk *SDK) AuthenticatedClient(c *gin.Context) (string, *relysia.Client, error) {
	if len(c.Request.Header["Authentication"]) == 0 {
		return "", nil, fmt.Errorf("no authentication header found")
	}
	token := c.Request.Header["Authentication"][0]
	return hex.EncodeToString(sdk.Hash([]byte(token))), sdk.relysia.WithToken(token), nil
}

func (sdk *SDK) AuthenticatedUser(c *gin.Context) (*UserSession, error) {
	ctx := context.Background()

	token, relysiaClient, err := sdk.AuthenticatedClient(c)
	if err != nil {
		return nil, err
	}

	doc, err := sdk.firestoreClient.Collection("sessions").Doc(token).Get(ctx)
	if err != nil {
		return nil, err
	}

	session := &UserSession{}
	if err := doc.DataTo(session); err != nil {
		return nil, err
	}
	session.sdk = sdk
	session.client = relysiaClient

	return session, nil
}
*/
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
