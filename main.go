package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/steveyiyo/file-upload-to-s3/internal/s3api"
	"github.com/steveyiyo/file-upload-to-s3/internal/webserver"
)

var S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region string

func init() {
	// Get value from .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	S3_KeyID = os.Getenv("S3_KeyID")
	S3_AppKey = os.Getenv("S3_AppKey")
	S3_bucket = os.Getenv("S3_bucket")
	S3_Endpoint = os.Getenv("S3_Endpoint")
	S3_Region = os.Getenv("S3_Region")

	if S3_KeyID == "" || S3_AppKey == "" || S3_bucket == "" || S3_Endpoint == "" || S3_Region == "" {
		log.Fatal("Error to loading environment.")
	}
}

func main() {
	api1 := s3api.New(S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region)
	webserver.Init(api1)
}
