package http_evert

import (
	"net/http"

	"github.com/Ada-lave/evert-http/iternal/services"
	"github.com/gin-gonic/gin"
)

type DocController struct{
	service *services.DocService
}

func (DC *DocController) FormatDoc(c *gin.Context) {
	var request DocRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	file, err := DC.service.FormatDocument(*request.File,request.Params)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "some went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": file,
	})
}

func NewDocController() *DocController {
	return &DocController{}
}