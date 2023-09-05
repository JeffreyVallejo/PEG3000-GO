package db

import (
	"PEG3000/src/core/domain"
	"cloud.google.com/go/firestore"
	"context"
	"go.uber.org/zap"
)

type GetLeagueDbAdapter interface {
	Get(id string) (league domain.League, err error)
}

type GetLeagueDb struct {
	*firestore.Client
}

func NewGetLeagueDbAdapter(client *firestore.Client) GetLeagueDb {
	return GetLeagueDb{
		Client: client,
	}
}

func (db GetLeagueDb) Get(id string) (league domain.League, err error) {
	zap.L().Info("Getting League", zap.String("id", id))
	ctx := context.Background()
	docSnap, err := db.Client.Collection(LeagueCollection).Doc(id).Get(ctx)
	if err != nil {
		zap.L().Error("Getting League", zap.Error(err))
	}

	doc, err := DataToLeagueDoc(docSnap)
	if err != nil {
		return
	}

	zap.L().Info("League Fetched", zap.String("id", id))
	return doc.ToDomain(), nil
}
