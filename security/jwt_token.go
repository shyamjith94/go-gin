package security

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shyamjith94/go-gin/collections"
	"github.com/shyamjith94/go-gin/configuration"
	"github.com/shyamjith94/go-gin/constants"
	"github.com/shyamjith94/go-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	UserId    string
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Phone     int
	Location  string
	jwt.StandardClaims
}

// generate all tokens
func GenerteAllTokens(id string, userName string, firstName string, lastName string,
	email string, phone int, location string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		UserId:         id,
		UserName:       userName,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Phone:          phone,
		Location:       location,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(constants.JwtTokenTimeOut)).Unix()},
	}

	refreshClainms := &SignedDetails{
		UserId:         id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(constants.JwtRefrshTokenTimeOut)).Unix()},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(configuration.GetJwtKey()))
	if err != nil {
		log.Fatal(err)
		return
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClainms).SignedString([]byte(configuration.GetJwtKey()))
	if err != nil {
		log.Fatal(err)
		return
	}
	return token, refreshToken, err
}

// generate new token using reffresh token
func GenerateTokenFromRefreshToken(siginedToken string) (signedToken string, signedRefreshToken string, err error) {
	claims, msg := ValidateToken(siginedToken)
	if msg != "" {
		return "", "", errors.New(msg)
	}
	var user models.Users
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.FindOne().SetProjection(bson.M{"password": 0})
	err = collections.UserCollection.FindOne(ctx, bson.M{"userid": claims.UserId}, opts).Decode(&user)
	if err != nil {
		return "", "", errors.New("not found user")
	}
	token, refreshToken, err := GenerteAllTokens(user.UserId, user.UserName, user.FirstName,
		user.LastName, user.Email, user.Phone, user.Location)

	return token, refreshToken, err
}

// validate token
func ValidateToken(signedToken string) (claims *SignedDetails, message string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configuration.GetJwtKey()), nil
	})

	if err != nil {
		message = err.Error()
		return
	}

	claims, verify := token.Claims.(*SignedDetails)
	if !verify {
		message = "token invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		message = "token expired"
		return
	}
	return claims, message
}
