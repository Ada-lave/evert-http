package main

import (
	"net/http"
	"github.com/Ada-lave/evert-http/iternal/http_evert"
	"github.com/gin-gonic/gin"
)

func setupEngine() *gin.Engine {

	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return server
}

func main() {
	server := setupEngine()
	api := server.Group("/api")

	http_evert.InitRouter(api)

	server.Run()


}