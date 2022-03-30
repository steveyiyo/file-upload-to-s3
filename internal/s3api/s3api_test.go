package s3api_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region string
var s3Client *s3.S3

// Initialize S3 client
func Init() {
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

	// Create S3 session
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(S3_KeyID, S3_AppKey, ""),
		Endpoint:         aws.String(S3_Endpoint),
		Region:           aws.String("us-west-002"),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)
	s3Client = s3.New(newSession)
}

func CreateBucket(bucket_name string) (bool, string) {
	// Pre define return values
	return_success := false
	return_message := "NULL"

	// Define environment variables
	bucket := aws.String(bucket_name)

	// Define bucket name
	cparams := &s3.CreateBucketInput{
		Bucket: bucket, // Required
	}

	// If bucket not exist, create it.
	_, err := s3Client.CreateBucket(cparams)
	if err != nil {
		return_success = false
		return_message = err.Error()
	} else {
		return_success = true
		return_message = "Successfully to created bucket: " + bucket_name
	}
	return return_success, return_message
}

func DownloadFile(bucket_name, filename string) (bool, string) {
	// Pre define return values
	var err error
	return_success := false
	return_message := "NULL"

	// Define environment variables
	bucket := aws.String(bucket_name)
	key := aws.String(filename)

	//Get Object
	_, err = s3Client.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})

	if err != nil {
		return_success = false
		return_message = "Failed to download file" + err.Error()
	} else {
		return_success = true
		return_message = "Successfully to download file: " + filename
	}
	return return_success, return_message
}

// Upload file to S3
func UploadFiletoS3(path, filename string) (bool, string) {
	var err error

	// Pre define return values
	return_success := false
	return_message := "NULL"

	// Define environment variables
	fmt.Println(S3_bucket)
	bucket := aws.String(S3_bucket)
	key := aws.String(path)

	// If bucket not exist, create it.
	// return_success, return_message = CreateBucket(S3_bucket)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("S3 Compatible API"),
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return_success = false
		return_message = "Failed to upload object" + S3_bucket + path + err.Error()
		fmt.Println(path)
	} else {
		return_success = true
		return_message = "Successfully to uploaded file: " + filename
	}
	return return_success, return_message
}

func TestMain(m *testing.M) {
	Init()
	success, message := UploadFiletoS3("testfile.txt", "testfile.txt")
	fmt.Println(success, message)
}
