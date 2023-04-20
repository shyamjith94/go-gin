package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/security"
)

func UserRoute(router *gin.Engine) {
	// middleware
	router.Use(security.Autherization())

	// route
	router.GET("/users", GetAllUsers)
	router.GET("/user/:userId", GetUser)
}

func SignUpAndSignInRoute(router *gin.Engine) {
	// route
	router.POST("/user", CreateUser)
	router.POST("/login", LoginUser)

}
