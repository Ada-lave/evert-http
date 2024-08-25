package http_evert

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DocController struct{}

func (DC *DocController) FormatDoc(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": "dddd",
	})
}

func NewDocController() *DocController {
	return &DocController{}
}