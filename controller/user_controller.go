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
	"github.com/shyamjith94/go-gin/security"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func CreateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		var user Users
// 		// defer cancel()
// 		if err := c.BindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
// 				Message: "error", Data: nil})
// 			return
// 		}
// 		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
// 			Message: "fine", Data: &user})
// 	}
// }

// creating new user
func CreateUser(c *gin.Context) {
	var user models.Users
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// check body has data
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: constants.NoBody, Data: nil})
		return
	}
	// validate json
	if validatorErr := validate.Struct(&user); validatorErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: validatorErr.Error(), Data: nil})
		return
	}
	// hashing password
	hashedPass, err := security.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: "Can't encrypt password", Data: nil})
		return
	}
	// create object in mongo
	user.Password = hashedPass
	user.Id = primitive.NewObjectID()
	user.UserId = user.Id.Hex()
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()
	_, err = collections.UserCollection.InsertOne(ctx, &user)
	if err != nil {
		c.JSON(http.StatusNotImplemented, response.Response{Status: http.StatusNotImplemented,
			Message: err.Error(), Data: nil})
		return
	}
	user.Password = ""
	c.JSON(http.StatusCreated, response.Response{Status: http.StatusCreated,
		Message: constants.Success, Data: &user})
}

// get all user
func GetAllUsers(c *gin.Context) {
	var users []models.Users
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ignore password field
	opts := options.Find().SetProjection(bson.M{"password": 0})
	result, err := collections.UserCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent,
			Message: err.Error(), Data: nil})
		return
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var user models.Users
		if err := result.Decode(&user); err != nil {
			c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent,
				Message: err.Error(), Data: nil})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK,
		Message: constants.Success, Data: &users})
}

// Get single user
func GetUser(c *gin.Context) {
	var user models.Users
	userId := c.Param("userId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.FindOne().SetProjection(bson.M{"password": 0})
	// objId, _ := primitive.ObjectIDFromHex(userId)
	err := collections.UserCollection.FindOne(ctx, bson.M{"userid": userId}, opts).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent, Message: err.Error(),
			Data: nil})
		return
	}
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK, Message: constants.Success,
		Data: &user})

}

// Login user
func LoginUser(c *gin.Context) {
	var user models.Users
	var loginDetails models.Login

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.BindJSON(&loginDetails)
	// check body has data
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: constants.NoBody, Data: nil})
		return
	}
	// validate json
	if validatorErr := validate.Struct(&loginDetails); validatorErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: validatorErr.Error(), Data: nil})
		return
	}
	//  get record from db based name
	err = collections.UserCollection.FindOne(ctx, bson.M{"username": loginDetails.UserName}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNoContent, response.Response{Status: http.StatusNoContent,
			Message: err.Error(), Data: nil})
		return
	}
	// verfify given registered pass
	passwordIsChecked, message := security.VerifyPassword(user.Password, loginDetails.Password)
	if !passwordIsChecked {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: message, Data: nil})
		return
	}
	// generate token
	token, refreshToken, err := security.GenerteAllTokens(user.UserId, user.UserName, user.FirstName, user.LastName,
		user.Email, user.Phone, user.Location)
	user.AccessToken = token
	user.RefreshToken = refreshToken
	user.Password = ""
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.Response{Status: http.StatusInternalServerError,
				Message: err.Error(), Data: nil})
			return
		}
	}
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK,
		Message: constants.Success, Data: &user})
}

// genrate new token from refresh token
func GenerateNewToken(c *gin.Context) {
	var token models.Token
	err := c.BindJSON(&token)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: constants.NoBody, Data: nil})
		return
	}
	if validatorErr := validate.Struct(&token); validatorErr != nil {
		c.JSON(http.StatusBadRequest, response.Response{Status: http.StatusBadRequest,
			Message: validatorErr.Error(), Data: nil})
		return
	}
	accessToken, refreshToken, err := security.GenerateTokenFromRefreshToken(token.RefreshToken)
	if err != nil {
		c.JSON(http.StatusFailedDependency, response.Response{Status: http.StatusFailedDependency,
			Message: err.Error(), Data: nil})
		return
	}
	token.RefreshToken = refreshToken
	token.AccessToken = accessToken
	c.JSON(http.StatusOK, response.Response{Status: http.StatusOK,
		Message: constants.Success, Data: &token})
}
