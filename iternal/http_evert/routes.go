package http_evert

import "github.com/gin-gonic/gin"

func InitRouter(e *gin.RouterGroup) {
	group := e.Group("format")

	docController := NewDocController()
	group.POST("/docx", docController.FormatDoc)
}