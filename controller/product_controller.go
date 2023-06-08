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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	user, err := utils.GetUser(c)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: err.Error(), Data: nil})
		return
	}
	// tax calculation and user id
	product.UserId = user.UserId
	product.TaxAmount = product.CalculateTax()
	product.Id = primitive.NewObjectID()
	product.ProductId = product.Id.Hex()
	product.CreatedAt = time.Now().Local()
	product.UpdatedAt = time.Now().Local()
	_, err = collections.ProductCollection.InsertOne(ctx, &product)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated,
		Message: constants.Success, Data: &product})
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	var product models.Product

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collections.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent,
			Message: err.Error(), Data: nil})
		return
	}
	defer result.Close(ctx)

	for result.Next(ctx) {
		if err := result.Decode(&product); err != nil {
			c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent,
				Message: err.Error(), Data: nil})
			return
		}
		products = append(products, product)
	}
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK,
		Message: constants.Success, Data: &products})
}

func GetProduct(c *gin.Context) {
	var product models.Product
	productId := c.Param("ProductId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collections.ProductCollection.FindOne(ctx, bson.M{"productid": productId}).Decode(&product)
	if err != nil {
		c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent,
			Message: err.Error(), Data: nil})
		return
	}
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK,
		Message: constants.Success, Data: &product})
}
