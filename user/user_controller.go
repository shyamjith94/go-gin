package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/response"
)

// var authCollection *mongo.Collection = configuration.GetCollection(configuration.DbClient, "users")
// var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user Users
		// defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
				Message: "error", Data: nil})
			return
		}
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: "fine", Data: &user})
	}
}
