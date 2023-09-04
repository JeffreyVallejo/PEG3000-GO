package db

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

type FirestoreDatabase struct {
	client *firestore.Client
}

func NewDatabase(client *firestore.Client) *FirestoreDatabase {
	return &FirestoreDatabase{
		client: client,
	}
}

func NewFirebaseClient() *FirestoreDatabase {
	opt := option.WithCredentialsFile("/Users/jeff/go/src/nuzlockePrivateKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		zap.L().Fatal("failed to create new FirestoreClient", zap.Error(err))
		return nil
	}

	client, err := app.Firestore(context.Background())

	if err != nil {
		zap.L().Fatal("Client was unable to be created", zap.Error(err))
	}
	db := NewDatabase(client)

	if err != nil {
		zap.L().Fatal("Error: ", zap.Error(err))
		return nil
	}

	zap.L().Info("Successfully connected to firebase")
	return db
}
