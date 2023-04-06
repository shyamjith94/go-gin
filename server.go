package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/configuration"
)

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())

	// db connection
	configuration.ConnectDataBase()

	// routes

	// run server
	server.Run(":8080")

}
