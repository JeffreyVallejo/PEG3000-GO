package main

import (
	"PEG3000/src/adapters/in/web"
	"PEG3000/src/adapters/out/db"
	"PEG3000/src/core/services"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	zap.L().Info("Starting Server . . .")

	client := db.NewFirebaseClient()
	getWildPokemon := db.NewGetWildPokemonDbAdapter(client)
	getLeague := db.NewGetLeagueDbAdapter(client)
	addLeague := db.NewAddLeagueDbAdapter(client)

	server := web.NewServer().
		WithGetWildPokemonAdapter(services.NewGetWildPokemonService(getWildPokemon)).
		WithGetLeagueAdapter(services.NewGetLeagueService(getLeague)).
		WithAddLeagueAdapter(services.NewAddLeagueService(addLeague))

	if err := server.Run(); err != nil {
		zap.L().Fatal("listen", zap.Error(err))
	}
}
