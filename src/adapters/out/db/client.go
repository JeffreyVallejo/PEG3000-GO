package db

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

const (
	WildPokemonCollection = "Encounter-Pokemon"
	LeagueCollection      = "League"
)

type FirestoreDatabase struct {
	client *firestore.Client
}

func NewDatabase(client *firestore.Client) *FirestoreDatabase {
	return &FirestoreDatabase{
		client: client,
	}
}

func NewFirebaseClient() *firestore.Client {
	opt := option.WithCredentialsFile("/Users/jeff/GitHub/PEG3000-GO/nuzlockePrivateKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		zap.L().Fatal("failed to create new FirestoreClient", zap.Error(err))
		return nil
	}

	client, err := app.Firestore(context.Background())

	if err != nil {
		zap.L().Fatal("Client was unable to be created", zap.Error(err))
	}

	zap.L().Info("Successfully connected to firebase")
	return client
}
