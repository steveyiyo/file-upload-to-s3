package webserver

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/image-upload/internal/s3api"
	"github.com/steveyiyo/image-upload/package/tools"
)

var S3API *s3api.S3API

type Result struct {
	Success   bool
	Message   string
	File_Name string
}

func UploadFile(c *gin.Context) {
	var r Result
	var return_check bool
	var return_message string
	return_status_code := 200

	file, header, err := c.Request.FormFile("upload_file")
	if err != nil {
		r = Result{false, "Bad Request!", ""}
		c.JSON(400, r)
		return
	}

	nsec := time.Now().UnixNano()
	extension := strings.Split(header.Filename, ".")[1]
	filename := strconv.FormatInt(nsec, 10) + "_" + tools.RandomString(5) + "." + extension

	if !os.IsNotExist(err) && tools.CreateFolder("tmp") {
		out, err := os.Create("tmp/" + filename)
		if err != nil {
			return_status_code = 400
			return_check = false
			return_message = "Error!"
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			return_status_code = 400
			return_check = false
			return_message = "Error!"
		} else {
			uploadS3_check, uploadS3_message := S3API.UploadFiletoS3("tmp/"+filename, filename)
			os.Remove("tmp/" + filename)
			if uploadS3_check {
				return_status_code = 201
				return_check = uploadS3_check
				return_message = uploadS3_message
			} else {
				return_status_code = 400
				return_check = uploadS3_check
				return_message = uploadS3_message
			}
		}
	} else {
		return_status_code = 400
		return_check = false
		return_message = "Failed to create the tmp folder!"
	}
	r = Result{return_check, return_message, filename}
	c.JSON(return_status_code, r)
}

func pageNotAvailable(c *gin.Context) {
	r := Result{false, "Page not available!", ""}
	c.JSON(404, r)
}

func Init(S3config *s3api.S3API) {
	// Define S3API
	S3API = S3config

	// Create a gin router
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Set CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// API Endpoints
	router.StaticFile("/", "./static/upload.html")
	router.NoRoute(pageNotAvailable)

	apiv1 := router.Group("/api/v1/")
	apiv1.POST("/upload", UploadFile)

	// Run
	router.Run("127.0.0.1:19725")
}
