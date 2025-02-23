package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	UseSSL          bool
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Endpoint = os.Getenv("ENDPOINT")
	AccessKeyID = os.Getenv("ACCESS_KEY_ID")
	SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	BucketName = os.Getenv("BUCKET_NAME")
	UseSSL, err = strconv.ParseBool(os.Getenv("USE_SSL"))
	if err != nil {
		log.Fatalf("Error parsing USE_SSL: %v", err)
	}
}
