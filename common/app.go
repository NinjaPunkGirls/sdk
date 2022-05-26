package common

import (
	"context"
	"encoding/hex"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/storage"

	firebase "firebase.google.com/go"
	"github.com/fxamacker/cbor/v2"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc `json:"-"`
}

type App struct {
	Gin       *gin.Engine
	Storage   *storage.Client
	Firestore *firestore.Client
	cbor      cbor.EncMode
	routes    []Route
}

func NewApp(projectID string) *App {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	fapp, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	firestoreClient, err := fapp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	storageClient, err := fapp.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	app := &App{
		Gin:       gin.Default(),
		Storage:   storageClient,
		Firestore: firestoreClient,
	}
	app.UseCBOR()

	return app
}

func (app *App) TimeNow() time.Time {
	return time.Now().UTC()
}

func (app *App) SeedDigest(input string) string {
	return hex.EncodeToString(app.SHA256([]byte(os.Getenv("SEED")), []byte(input)))
}

func (app *App) UseCBOR() {
	// setup CBOR encoer
	cb, err := cbor.CanonicalEncOptions().EncMode()
	if err != nil {
		log.Fatalln(err)
	}
	app.cbor = cb
}

func (app *App) Close() {
	app.Firestore.Close()
}
