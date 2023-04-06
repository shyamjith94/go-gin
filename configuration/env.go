package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load ENV file")
	}
	return os.Getenv("MONGODBURL")
}
