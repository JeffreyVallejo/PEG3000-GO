package services

import (
	"PEG3000/src/adapters/out/db"
	"PEG3000/src/core/domain"
	"go.uber.org/zap"
)

type GetLeagueUseCase interface {
	Get(id string) (League domain.League, err error)
}

type GetLeagueService struct {
	db.GetLeagueDbAdapter
}

func GetNewLeagueService(adapter db.GetLeagueDbAdapter) GetLeagueService {
	return GetLeagueService{
		GetLeagueDbAdapter: adapter,
	}
}

func (service GetLeagueService) Get(id string) (league domain.League, err error) {
	zap.L().Info("Get League Service")
	return service.GetLeagueDbAdapter.Get(id)

}
