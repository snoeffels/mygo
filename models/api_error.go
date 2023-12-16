package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func SendErrorResponse(c *gin.Context, code int, error string) {
	c.JSON(code, APIError{Error: error, Status: code})
}

func SendSimpleErrorResponse(c *gin.Context, code int) {
	c.JSON(code, APIError{Error: http.StatusText(code), Status: code})
}
