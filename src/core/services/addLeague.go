package services

import (
	"PEG3000/src/adapters/out/db"
	"PEG3000/src/core/domain"
	"PEG3000/src/core/ports/in"
	"go.uber.org/zap"
)

type AddLeagueUseCase interface {
	Add(command in.AddLeagueCommand) (league domain.League, err error)
}

type AddLeagueService struct {
	db.AddLeagueDbAdapter
}

func NewAddLeagueService(adapter db.AddLeagueDbAdapter) AddLeagueService {
	return AddLeagueService{
		AddLeagueDbAdapter: adapter,
	}
}

// Add returns domain.League
func (service AddLeagueService) Add(command in.AddLeagueCommand) (league domain.League, err error) {
	zap.L().Info("Add League Service")
	params := domain.NewLeagueParams{
		DisplayName: command.DisplayName,
		OwnerId:     command.OwnerId,
		SecureId:    command.SecureId,
	}

	league = domain.NewLeague(params)
	return service.AddLeagueDbAdapter.Add(league)
}
