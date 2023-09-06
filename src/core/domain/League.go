package domain

import (
	"time"
)

// A League is the domain model.
type League struct {
	Id          string
	DisplayName string
	OwnerId     string
	SecureId    string
	IsActive    bool
	IsRegOpen   bool
	CreatedAt   int64
}

// A NewLeagueParams serves content from a request.
type NewLeagueParams struct {
	DisplayName string
	OwnerId     string
	SecureId    string
}

// NewLeague returns a new domain.League from NewLeagueParams
// The CreatedAt field is set to time.Now()
func NewLeague(params NewLeagueParams) League {
	now := time.Now().Unix()
	return League{
		DisplayName: params.DisplayName,
		OwnerId:     params.OwnerId,
		SecureId:    params.SecureId,
		IsActive:    true,
		IsRegOpen:   true,
		CreatedAt:   now,
	}
}
