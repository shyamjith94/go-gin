package security

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shyamjith94/go-gin/configuration"
	"github.com/shyamjith94/go-gin/constants"
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
