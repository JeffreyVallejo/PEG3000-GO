package main

import (
	"PEG3000/src/adapters/in/web"
	"PEG3000/src/adapters/out/db"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	zap.L().Info("Starting Server . . .")

	client := db.NewFirebaseClient()
	getWildPokemon := db.NewGetWildPokemonDbAdapter(client)
	server := web.NewServer().
		WithGetWildPokemonAdapter(services.NewGetWildPokemonService(getWildPokemon))

	if err := server.Run(); err != nil {
		zap.L().Fatal("listen", zap.Error(err))
	}
}
