package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error" binding:"required"`
}

func InternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ErrorMessage{Error: err.Error()})
}

func BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorMessage{Error: err.Error()})
}

func ValidationNameError() error {
	return fmt.Errorf("validation error: displayName is empty or only had white spaces")
}
