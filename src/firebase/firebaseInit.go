package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const (
	firebaseConfigFile = "firebaseConfig.json"
)

var (
	ctx    context.Context
	client *firestore.Client
)

func InitFirebase() {
	// Initialize Firebase
	ctx = context.Background()
	opt := option.WithCredentialsFile(firebaseConfigFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization error: %v\n", err)
	}

	// Initialize Firestore
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Firestore initialization error: %v\n", err)
	}
	defer client.Close()
}
