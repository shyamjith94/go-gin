package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/configuration"
	home "github.com/shyamjith94/go-gin/homepage"
	"github.com/shyamjith94/go-gin/router"
)

func main() {
	server := gin.Default()
	server.Use(gin.Recovery())
	server.Use(gin.Logger())

	// middleware

	// db connection
	configuration.ConnectDataBase()

	// routes
	home.HomeRoute(server)
	router.SignUpAndSignInRoute(server)
	router.UserRoute(server)
	router.ProductRoute(server)
	router.CategoryRoute(server)

	// run server
	server.Run(":8000")

}
