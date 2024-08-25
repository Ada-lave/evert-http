package http_evert

import (
	"net/http"

	"github.com/Ada-lave/evert-core"
	"github.com/Ada-lave/evert-http/iternal/services"
	"github.com/gin-gonic/gin"
)

type DocController struct{
	service *services.DocService
}

func (DC *DocController) FormatDoc(c *gin.Context) {
	var request DocRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	params := evert.FormatterParams{
		AddSpacesBeetweenImageText: request.AddSpacesBeetweenImageText,
		FormatImagDescription: request.FormatImageDescription,
	}
	file, err := DC.service.FormatDocument(*request.File, params)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "some went wrong",
		})
	}

	c.Header("Content-Description", "File Transfer")
    c.Header("Content-Transfer-Encoding", "binary")
    c.Header("Content-Disposition", "attachment; filename="+request.File.Filename)
    c.Header("Content-Type", "application/octet-stream")

	c.File(file)
}

func NewDocController() *DocController {
	return &DocController{}
}