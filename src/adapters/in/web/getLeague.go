package web

import (
	"PEG3000/src/core/domain"
	"PEG3000/src/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetLeagueWeb struct {
	services.GetLeagueUseCase
}

func (server *Server) WithGetLeagueAdapter(service services.GetLeagueUseCase) *Server {
	adapter := NewGetLeagueWebAdapter(service)
	server.GET(LeagueRoute+PathParam+Id, adapter.GetLeagueHandler)
	return server
}

func NewGetLeagueWebAdapter(service services.GetLeagueUseCase) GetLeagueWeb {
	return GetLeagueWeb{GetLeagueUseCase: service}
}

func (adapter *GetLeagueWeb) GetLeagueHandler(c *gin.Context) {
	id := c.Param(Id)
	League, err := adapter.Get(id)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	response := GetLeagueResponse{League: League}
	c.JSON(http.StatusOK, response)
}

type GetLeagueResponse struct {
	domain.League `json:"League" binding:"required"`
}
