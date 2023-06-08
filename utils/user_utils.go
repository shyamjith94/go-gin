package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/models"
)

func GetUser(c *gin.Context) (user *models.Users, err error) {
	if _, found := c.Keys["User"]; found {
		return c.Keys["User"].(*models.Users), nil
	}
	return user, errors.New("Auth token not have userinfo")
}
