package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load ENV file")
	}
	return os.Getenv("MONGODBURL")
}

func GetJwtKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load ENV file")
	}
	return os.Getenv("SECURITY_KEY")
}
