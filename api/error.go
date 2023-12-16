package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func SendErrorResponse(c *gin.Context, code int, error string) {
	c.JSON(code, Error{Error: error, Status: code})
}

func SendSimpleErrorResponse(c *gin.Context, code int) {
	c.JSON(code, Error{Error: http.StatusText(code), Status: code})
}
