package db

import (
	"PEG3000/src/adapters/core/domain"
	"go.uber.org/zap"
)

type GetWildPokemonDbAdaptor interface {
	Get(location string) (wildPokemon domain.WildPokemon, err error)
}

type GetWildPokemonDb struct {
	Client *FirestoreDatabase
}

func NewGetWildPokemonDbAdapter(client *FirestoreDatabase) GetWildPokemonDb {
	return GetWildPokemonDb{
		Client: client,
	}
}

func (db GetWildPokemonDb) Get(location string) (wildPokemon domain.WildPokemon, err error) {
	zap.L().Info("GET")
	return domain.WildPokemon{}, nil
}
