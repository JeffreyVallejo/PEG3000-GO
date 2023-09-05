package services

import (
	"PEG3000/src/adapters/in/web/dto"
	"PEG3000/src/adapters/out/db"
	"go.uber.org/zap"
)

type GetWildPokemonUseCases interface {
	Get(location string) (Encounter dto.Encounter, err error)
}

type GetWildPokemonService struct {
	db.GetWildPokemonDbAdaptor
}

func NewGetWildPokemonService(adapter db.GetWildPokemonDbAdaptor) GetWildPokemonService {
	return GetWildPokemonService{
		GetWildPokemonDbAdaptor: adapter,
	}
}

func (service GetWildPokemonService) Get(location string) (encounter dto.Encounter, err error) {
	zap.L().Info("Get Wild Pokemon Service")
	wildPokemon, err := service.GetWildPokemonDbAdaptor.Get(location)
	return wildPokemon.GenerateEncounter(), err
}
