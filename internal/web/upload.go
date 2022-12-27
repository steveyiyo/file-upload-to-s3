package web

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/file-upload-to-s3/internal/s3api"
	"github.com/steveyiyo/file-upload-to-s3/pkg/tools"
)

var S3API *s3api.S3API

type Result struct {
	Success   bool
	Message   string
	File_Name string
	File_URL  string
}

func UploadFile(c *gin.Context) {
	checkSession := VerifySessionAuth(c)

	if checkSession.jwtSuccess {
		var r Result
		var returnCheck bool
		var returnMessage string
		var returnFileURL string
		returnStatusCode := 200

		file, header, err := c.Request.FormFile("upload_file")
		if err != nil {
			r = Result{false, "Bad Request!", "", ""}
			c.JSON(400, r)
			return
		}

		nsec := time.Now().UnixNano()
		extension := strings.Split(header.Filename, ".")[1]
		filename := fmt.Sprintf("%s_%s.%s", strconv.FormatInt(nsec, 10), tools.RandomString(5), extension)

		if !os.IsNotExist(err) && tools.CreateFolder("tmp") {
			out, err := os.Create("tmp/" + filename)
			if err != nil {
				returnStatusCode = 400
				returnCheck = false
				returnMessage = err.Error()
				returnFileURL = ""
			}

			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				returnStatusCode = 400
				returnCheck = false
				returnMessage = err.Error()
				returnFileURL = ""
			} else {
				uploadS3Check, uploadS3Message := S3API.UploadFiletoS3("tmp/"+filename, filename)
				os.Remove("tmp/" + filename)
				if uploadS3Check {
					returnStatusCode = 201
					returnCheck = uploadS3Check
					returnMessage = uploadS3Message
					returnFileURL = fmt.Sprintf("https://s3.us-west-002.backblazeb2.com/steveyi-drive/tmp/%s", filename)
				} else {
					returnStatusCode = 400
					returnCheck = uploadS3Check
					returnMessage = uploadS3Message
					returnFileURL = ""
				}
			}
		} else {
			returnStatusCode = 400
			returnCheck = false
			returnMessage = "Failed to create the tmp folder!"
		}
		r = Result{returnCheck, returnMessage, filename, returnFileURL}
		c.JSON(returnStatusCode, r)
	} else {
		r := Result{false, "Session expired!", "", ""}
		c.JSON(401, r)
	}
}
