package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

const firebaseConfigFile = "firebaseConfig.json"

var (
	Ctx        context.Context
	Client     *firestore.Client
	AuthClient *auth.Client
)

func InitializeFirebase() {
	Ctx = context.Background()
	opt := option.WithCredentialsFile(firebaseConfigFile)
	app, err := firebase.NewApp(Ctx, nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization error: %v\n", err)
	}

	Client, err = app.Firestore(Ctx)
	if err != nil {
		log.Fatalf("Firestore initialization error: %v\n", err)
	}
	// Initialize Firebase Authentication
	AuthClient, err = app.Auth(Ctx)
	if err != nil {
		log.Fatalf("Firebase Auth initialization error: %v\n", err)
	}
}
