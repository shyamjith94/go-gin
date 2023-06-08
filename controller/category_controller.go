package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shyamjith94/go-gin/collections"
	"github.com/shyamjith94/go-gin/constants"
	"github.com/shyamjith94/go-gin/models"
	"github.com/shyamjith94/go-gin/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: constants.NoBody, Data: nil})
		return
	}
	if validateErr := validate.Struct(&category); validateErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: validateErr.Error(), Data: nil})
		return
	}
	category.Id = primitive.NewObjectID()
	category.CategoryId = category.Id.Hex()
	_, err = collections.CategoryCollection.InsertOne(ctx, &category)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated,
		Message: constants.Success, Data: &category})
}
