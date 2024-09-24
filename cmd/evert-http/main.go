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

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
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
	server.Use(CORSMiddleware())
	gin.SetMode(gin.ReleaseMode)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := server.Group("/api")

	http_evert.InitRouter(api)

	server.Run()
}