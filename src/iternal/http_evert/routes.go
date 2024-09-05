package http_evert

import "github.com/gin-gonic/gin"

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /accounts/{id} [get]
func InitRouter(e *gin.RouterGroup) {
	group := e.Group("format")

	docController := NewDocController()
	group.POST("/docx", docController.FormatDoc)
}