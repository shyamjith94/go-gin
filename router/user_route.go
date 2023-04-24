package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/controller"
	"github.com/shyamjith94/go-gin/security"
)

func UserRoute(route *gin.Engine) {
	// middleware
	route.Use(security.Autherization())

	// route
	route.GET("/users", controller.GetAllUsers)
	route.GET("/user/:userId", controller.GetUser)
}

func SignUpAndSignInRoute(route *gin.Engine) {
	// route
	route.POST("/user", controller.CreateUser)
	route.POST("/login", controller.LoginUser)

}
