package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/steveyiyo/file-upload-to-s3/internal/s3api"
	"github.com/steveyiyo/file-upload-to-s3/internal/web"
)

var S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region, S3_URL_PATH string

func init() {

	// Get value from .env
	err := godotenv.Load()

	S3_KeyID = os.Getenv("S3_KEYID")
	S3_AppKey = os.Getenv("S3_APPKEY")
	S3_bucket = os.Getenv("S3_BUCKET")
	S3_Endpoint = os.Getenv("S3_ENDPOINT")
	S3_Region = os.Getenv("S3_REGION")
	S3_URL_PATH = os.Getenv("S3_URL_PATH")

	if !checkEnvironment() {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func checkEnvironment() bool {
	if S3_KeyID == "" || S3_AppKey == "" || S3_bucket == "" || S3_Endpoint == "" || S3_Region == "" || S3_URL_PATH == "" {
		return false
	}
	return true
}

func main() {
	api1 := s3api.New(S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region)
	web.Init(api1)
}
