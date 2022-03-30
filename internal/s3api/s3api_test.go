package s3api

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Get value from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	S3_KeyID := os.Getenv("S3_KeyID")
	S3_AppKey := os.Getenv("S3_AppKey")
	S3_bucket := os.Getenv("S3_bucket")
	S3_Endpoint := os.Getenv("S3_Endpoint")
	S3_Region := os.Getenv("S3_Region")
	new1 := New(S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region)
	success, message := new1.UploadFiletoS3("aaa.txt", "aaa.txt")
	fmt.Println(success, message)
}
