package api

import (
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/image-upload/package/tools"
)

type Result struct {
	Success bool
	Message string
}

func Uploadfile(c *gin.Context) {
	file, _ := c.FormFile("filename") // get file from form input name 'file'
	file.Filename = tools.RandomString(8)
	c.SaveUploadedFile(file, "tmp/"+file.Filename) // save file to tmp folder in current directory

	var result Result
	result = Result{true, "File uploaded successfully. " + file.Filename}
	c.JSON(200, result)
}
