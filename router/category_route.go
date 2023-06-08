package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/controller"
)

func CategoryRoute(route *gin.Engine) {
	route.POST("/product/category", controller.CreateCategory)

}
