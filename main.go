package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/steveyiyo/image-upload/internal/s3api"
	"github.com/steveyiyo/image-upload/internal/webserver"
)

var S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region string

func init() {
	// Get value from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	S3_KeyID = os.Getenv("S3_KeyID")
	S3_AppKey = os.Getenv("S3_AppKey")
	S3_bucket = os.Getenv("S3_bucket")
	S3_Endpoint = os.Getenv("S3_Endpoint")
	S3_Region = os.Getenv("S3_Region")
}

func main() {
	api1 := s3api.New(S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region)
	webserver.Init(api1)
}
