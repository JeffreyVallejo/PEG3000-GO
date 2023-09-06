package db

import (
	"PEG3000/src/core/domain"
	"cloud.google.com/go/firestore"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

// AddLeagueDbAdapter describes how to interact with the Db.
type AddLeagueDbAdapter interface {
	Add(league domain.League) (domain.League, error)
}

// AddLeagueDb persists items to Cloud Firestore
// See https://cloud.google.com/firestore/docs.
type AddLeagueDb struct {
	*firestore.Client
}

// NewAddLeagueDbAdapter returns a new firebase-port to cloud Firestore for adding New LeagueDoc Documents.
// See the db package for details on creating a suitable
// db.Client: https://godoc.org/cloud.google.com/go/firestore.
func NewAddLeagueDbAdapter(client *firestore.Client) AddLeagueDb {
	return AddLeagueDb{
		Client: client,
	}
}

// Add creates a document to add to firebase then returns domain.League from the document.
func (out AddLeagueDb) Add(league domain.League) (domain.League, error) {
	zap.L().Info("Adding League")

	// Create a new LeagueDoc from the league domain to persist to the DB
	doc := NewLeagueDoc(league)
	docRef, _, err := out.Client.Collection(LeagueCollection).Add(context.Background(), doc)
	if err != nil {
		zap.L().Error("Adding League", zap.Error(err))
	}
	// Set the DocumentRef ID from firebase to the doc.
	doc.DocumentRef = docRef
	zap.L().Info("League Added", zap.String("id", doc.ID))

	//Convert the doc to the domain model and return.
	return doc.ToDomain(), nil
}
