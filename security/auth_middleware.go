package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/response"
)

func Autherization() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
				Message: "No Access Token", Data: nil})
			c.Abort()
			return
		}
		claims, err := ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
				Message: err, Data: nil})
			c.Abort()
			return
		}
		c.Set("UserId", claims.UserId)
		c.Set("UserName", claims.UserName)
		c.Set("firstname", claims.FirstName)
		c.Set("FirstName", claims.LastName)
		c.Set("Email", claims.Email)
		c.Set("Location", claims.Location)
		c.Next()
	}
}
