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
	"github.com/shyamjith94/go-gin/utils"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check body has data
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: constants.NoBody, Data: nil})
		return
	}

	// validate json
	if validateErr := validate.Struct(&product); validateErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: validateErr.Error(), Data: nil})
		return
	}

	// create product db
	user, err := utils.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: err.Error(), Data: nil})
		return
	}
	// tax calculation and user id
	product.UserId = user.UserId
	product.TaxAmount = product.CalculateTax()
	_, err = collections.ProductCollection.InsertOne(ctx, &product)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated,
		Message: constants.Success, Data: &product})
}
