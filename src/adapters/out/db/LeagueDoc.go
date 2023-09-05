package db

import (
	"PEG3000/src/core/domain"
	"cloud.google.com/go/firestore"
	"go.uber.org/zap"
)

type LeagueDoc struct {
	*firestore.DocumentRef `firestore:"-"`
	DisplayName            string `firestore:"displayName"`
	OwnerId                string `firestore:"ownerId"`
	SecureId               string `firestore:"secureId"`
	IsActive               bool   `firestore:"isActive"`
	IsRegOpen              bool   `firestore:"isRegOpen"`
}

func NewLeagueDoc(league domain.League) LeagueDoc {
	return LeagueDoc{
		DisplayName: league.DisplayName,
		OwnerId:     league.OwnerId,
		SecureId:    league.SecureId,
		IsActive:    league.IsActive,
		IsRegOpen:   league.IsRegOpen,
	}
}

func DataToLeagueDoc(docSnap *firestore.DocumentSnapshot) (doc LeagueDoc, err error) {
	if err = docSnap.DataTo(&doc); err != nil {
		zap.L().Error("DataTo LeagueDoc", zap.Error(err))
	}
	doc.DocumentRef = docSnap.Ref
	return
}

func (doc LeagueDoc) ToDomain() domain.League {
	return domain.League{
		Id:          doc.DocumentRef.ID,
		DisplayName: doc.DisplayName,
		OwnerId:     doc.OwnerId,
		SecureId:    doc.SecureId,
		IsActive:    doc.IsActive,
		IsRegOpen:   doc.IsRegOpen,
	}
}
