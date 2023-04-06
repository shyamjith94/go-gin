package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/response"
)

func HomeRoute(router *gin.Engine) {
	router.GET("/", homePageControl)
}

func homePageControl(c *gin.Context) {
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: "server running", Data: nil})
}
