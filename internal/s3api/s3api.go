package s3api

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3API struct {
	bucket    string
	awsConfig *aws.Config
	session   *session.Session
	client    *s3.S3
}

// New S3 client
func New(S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region string) *S3API {

	// Create S3 session
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(S3_KeyID, S3_AppKey, ""),
		Endpoint:         aws.String(S3_Endpoint),
		Region:           aws.String("us-west-002"),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	NewConfig := &S3API{
		bucket:    S3_bucket,
		awsConfig: s3Config,
		session:   newSession,
		client:    s3Client,
	}

	// fmt.Println(S3_KeyID, S3_AppKey, S3_bucket, S3_Endpoint, S3_Region)
	return NewConfig
}

// Create Bucket
func (s3Client *S3API) CreateBucket(bucket_name string) (bool, string) {
	// Pre define return values
	returnSuccess := false
	returnMessage := "NULL"

	// Define environment variables
	bucket := aws.String(bucket_name)

	// Define bucket name
	cparams := &s3.CreateBucketInput{
		Bucket: bucket, // Required
	}

	// If bucket not exist, create it.
	_, err := s3Client.client.CreateBucket(cparams)
	if err != nil {
		returnSuccess = false
		returnMessage = err.Error()
	} else {
		returnSuccess = true
		returnMessage = "Successfully to created bucket: " + bucket_name
	}
	return returnSuccess, returnMessage
}

// Download file
func (s3Client *S3API) DownloadFile(bucket_name, filename string) (bool, string) {
	// Pre define return values
	var err error
	returnSuccess := false
	returnMessage := "NULL"

	// Define environment variables
	bucket := aws.String(bucket_name)
	key := aws.String(filename)

	//Get Object
	_, err = s3Client.client.GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})

	if err != nil {
		returnSuccess = false
		returnMessage = "Failed to download file" + err.Error()
	} else {
		returnSuccess = true
		returnMessage = "Successfully to download file: " + filename
	}
	return returnSuccess, returnMessage
}

// Upload file to S3
func (s3Client *S3API) UploadFiletoS3(path, filename string) (bool, string) {
	var err error
	// Pre define return values
	returnSuccess := false
	returnMessage := "NULL"

	// Define environment variables
	bucket := aws.String(s3Client.bucket)
	key := aws.String(path)

	// If bucket not exist, create it.
	// returnSuccess, returnMessage = CreateBucket(S3_bucket)

	_, err = s3Client.client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("S3 Compatible API"),
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		returnSuccess = false
		returnMessage = "Failed to upload object" + s3Client.bucket + path + err.Error()
	} else {
		returnSuccess = true
		returnMessage = "Successfully to uploaded file: " + filename
	}
	return returnSuccess, returnMessage
}
