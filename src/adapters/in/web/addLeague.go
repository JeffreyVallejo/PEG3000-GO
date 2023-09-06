package web

import (
	"PEG3000/src/adapters/in/web/dto"
	"PEG3000/src/core/ports/in"
	"PEG3000/src/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AddLeagueWeb struct {
	services.AddLeagueUseCase
}

type AddLeagueRequest struct {
	DisplayName string `json:"displayName" binding:"required"`
	OwnerId     string `json:"ownerId" binding:"required"`
	SecureId    string `json:"secureId" binding:"required"`
}

type AddLeagueResponse struct {
	dto.League `json:"league" binding:"required"`
}

// WithAddLeagueAdapter returns a Server
func (server *Server) WithAddLeagueAdapter(service services.AddLeagueUseCase) *Server {
	//Create a new adapter
	adapter := NewAddLeagueWebAdapter(service)
	//Describe the endpoint
	server.POST(LeagueRoute, adapter.AddLeagueHandler)
	return server
}

// NewAddLeagueWebAdapter returns a new instance of AddLeagueWeb
func NewAddLeagueWebAdapter(service services.AddLeagueUseCase) AddLeagueWeb {
	// Construct a new AddLeagueWeb
	return AddLeagueWeb{
		AddLeagueUseCase: service,
	}
}

// AddLeagueHandler handles the request then returns a response.
func (adapter *AddLeagueWeb) AddLeagueHandler(c *gin.Context) {
	request, err := NewAddLeagueRequest(c)
	if err != nil {
		return
	}

	command := request.ToCommand()
	league, err := adapter.Add(command)
	if err != nil {
		InternalServerErrorResponse(c, err)
		return
	}

	response := AddLeagueResponse{League: dto.NewLeague(league)}
	c.JSON(http.StatusCreated, response)
}

// NewAddLeagueRequest returns AddLeagueRequest and error.
// JSON is bound to AddLeageRequest then Validated.
func NewAddLeagueRequest(c *gin.Context) (request *AddLeagueRequest, err error) {
	// Bind the JSON to AddLeagueRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		BadRequestResponse(c, err)
		return
	}

	// Clean and Validate request data
	if err = request.CleanAndValidate(); err != nil {
		BadRequestResponse(c, err)
		return
	}
	return
}

// CleanAndValidate validates request input and returns error if invalid.
func (request *AddLeagueRequest) CleanAndValidate() error {
	//Clean white spaces
	request.DisplayName = strings.TrimSpace(request.DisplayName)
	if request.DisplayName == "" {
		return ValidationNameError()
	}

	request.OwnerId = strings.TrimSpace(strings.ToLower(request.OwnerId))
	request.SecureId = strings.TrimSpace(request.SecureId)
	return nil
}

// ToCommand returns in.AddLeagueCommand
func (request *AddLeagueRequest) ToCommand() in.AddLeagueCommand {
	return in.AddLeagueCommand{
		DisplayName: request.DisplayName,
		OwnerId:     request.OwnerId,
		SecureId:    request.SecureId,
	}
}
