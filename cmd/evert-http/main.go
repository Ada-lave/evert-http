package main

import (
	"net/http"
	"github.com/Ada-lave/evert-http/src/iternal/http_evert"
	"github.com/gin-gonic/gin"
	docs "github.com/Ada-lave/evert-http/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func setupEngine() *gin.Engine {

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return server
}


// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := setupEngine()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := server.Group("/api")

	http_evert.InitRouter(api)

	server.Run()


}