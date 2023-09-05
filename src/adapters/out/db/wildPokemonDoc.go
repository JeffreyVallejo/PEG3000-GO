package db

import (
	"PEG3000/src/core/domain"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"go.uber.org/zap"
	"log"
)

type WildPokemonDocList struct {
	*firestore.DocumentRef `firestore:"-"` // Explicitly tells converter to ignore.
	AllPokemon             map[string]interface{}
}

func DataToWildPokemonDoc(docSnap *firestore.DocumentSnapshot) (doc WildPokemonDocList, err error) {
	if doc.AllPokemon = docSnap.Data(); err != nil {
		zap.L().Error("DataTo WildPokemonDoc", zap.Error(err))
		return
	}
	doc.DocumentRef = docSnap.Ref
	return
}

func (doc WildPokemonDocList) ToDomain() []domain.WildPokemon {

	var temp []domain.WildPokemon

	for _, val := range doc.AllPokemon {
		jsonData, err := json.Marshal(val)
		if err != nil {
			log.Fatal(err)
		}

		var mon domain.WildPokemon
		err = json.Unmarshal(jsonData, &mon)
		if err != nil {
			log.Fatal(err)
		}

		temp = append(temp, mon)
	}
	return temp
}
