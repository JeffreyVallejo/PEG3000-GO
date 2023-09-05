package web

import (
	"PEG3000/src/adapters/in/web/dto"
	"PEG3000/src/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetWildPokemonWeb struct {
	services.GetWildPokemonUseCases
}

type GetWildPokemonResponse struct {
	Encounter dto.Encounter `json:"WildPokemon" binding:"required"`
}

func (server *Server) WithGetWildPokemonAdapter(service services.GetWildPokemonUseCases) *Server {
	adapter := NewGetWildPokemonWebAdapter(service)
	server.GET(EncounterRoute+PathParam+Location, adapter.GetWildPokemonHandler)
	return server
}

func NewGetWildPokemonWebAdapter(service services.GetWildPokemonUseCases) GetWildPokemonWeb {
	return GetWildPokemonWeb{
		GetWildPokemonUseCases: service,
	}
}

// Entry Point
func (adapter *GetWildPokemonWeb) GetWildPokemonHandler(c *gin.Context) {
	location := c.Param(Location)
	encounter, err := adapter.Get(location)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}
	response := GetWildPokemonResponse{Encounter: encounter}
	c.JSON(http.StatusOK, response)
}
