package db

import (
	"PEG3000/src/core/domain"
	"cloud.google.com/go/firestore"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type GetWildPokemonDbAdaptor interface {
	Get(location string) (wildPokemon domain.WildPokemonList, err error)
}

type GetWildPokemonDb struct {
	*firestore.Client
}

func NewGetWildPokemonDbAdapter(client *firestore.Client) GetWildPokemonDb {
	return GetWildPokemonDb{
		Client: client,
	}
}

func (db GetWildPokemonDb) Get(location string) (wildPokemon domain.WildPokemonList, err error) {
	zap.L().Info("Getting Wild Pokemon", zap.String("location", location))
	ctx := context.Background()
	docSnap, err := db.Client.Collection(WildPokemonCollection).Doc(location).Get(ctx)
	if err != nil {
		zap.L().Error("Getting Wildpokemon", zap.Error(err))
	}

	doc, err := DataToWildPokemonDoc(docSnap)
	if err != nil {
		return
	}

	wildPokemon.Pokemon = doc.ToDomain()
	zap.L().Info("WildPokemon Fetched", zap.String("location", location))

	return wildPokemon, nil
}
