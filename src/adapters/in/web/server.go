package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RootRoute      = "/"
	HealthRoute    = "/health"
	EncounterRoute = "/encounter"
	PathParam      = "/:"
	Location       = "location"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	router.GET(RootRoute, rootHandler)
	router.GET(HealthRoute, healthHandler)

	return &Server{
		Engine: router,
	}
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to PEG server.",
	})
}
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Healthy",
	})
}
