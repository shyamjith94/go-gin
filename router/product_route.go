package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/controller"
)

func ProductRoute(route *gin.Engine) {
	route.POST("/product", controller.CreateProduct)
}
