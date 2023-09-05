package domain

import (
	"PEG3000/src/adapters/in/web/dto"
	"math/rand"
)

type WildPokemonList struct {
	Pokemon []WildPokemon
}

type WildPokemon struct {
	Name           string `json:"name"`
	AppearanceRate int    `json:"appearance rate"`
	Exclusive      string `json:"exclusive"`
}

func (pokemon *WildPokemonList) GenerateEncounter() dto.Encounter {
	return dto.Encounter{
		Name: pokemon.Pokemon[rand.Intn(len(pokemon.Pokemon))].Name,
	}
}
