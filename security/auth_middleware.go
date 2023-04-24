package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/models"
	"github.com/shyamjith94/go-gin/response"
)

func Autherization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Users
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
		user.UserId = claims.UserId
		user.UserName = claims.UserName
		user.FirstName = claims.FirstName
		user.LastName = claims.LastName
		user.Email = claims.Email
		user.Location = claims.Location
		c.Set("User", &user)
		c.Next()
	}
}
