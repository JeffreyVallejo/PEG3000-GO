package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error" binding:"required"`
}

func InternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ErrorMessage{Error: err.Error()})
}
