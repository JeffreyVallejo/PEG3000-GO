package dto

import "PEG3000/src/core/domain"

type League struct {
	Id          string `json:"id" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
	OwnerId     string `json:"ownerId" binding:"required"`
	SecureId    string `json:"secureId" binding:"required"`
	IsActive    bool   `json:"isActive" binding:"required"`
	IsRegOpen   bool   `json:"isRegOpen" binding:"required"`
	CreatedAt   int64  `json:"createdAt" binding:"required"`
}

func NewLeague(league domain.League) League {
	return League{
		Id:          league.Id,
		DisplayName: league.DisplayName,
		OwnerId:     league.OwnerId,
		SecureId:    league.SecureId,
		IsActive:    league.IsActive,
		IsRegOpen:   league.IsRegOpen,
	}
}

func (league League) ToDomain() domain.League {
	return domain.League{
		Id:          league.Id,
		DisplayName: league.DisplayName,
		OwnerId:     league.OwnerId,
		SecureId:    league.SecureId,
		IsActive:    league.IsActive,
		IsRegOpen:   league.IsRegOpen,
		CreatedAt:   league.CreatedAt,
	}
}
