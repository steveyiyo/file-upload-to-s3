package web

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/file-upload-to-s3/internal/s3api"
)

func indexPage(c *gin.Context) {
	checkSession := VerifySessionAuth(c)

	if checkSession.jwtSuccess {
		c.HTML(200, "upload.tmpl", gin.H{})
	} else {
		c.Redirect(302, "/auth/login")
	}
}

func pageNotAvailable(c *gin.Context) {
	r := Result{false, "Page not available!", "", ""}
	c.JSON(404, r)
}

func Init(S3config *s3api.S3API) {
	// Define S3API
	S3API = S3config

	// Create a gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Set CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// API Endpoints
	router.LoadHTMLGlob("static/*")
	router.GET("/", indexPage)
	router.GET("/auth/login", authLoginPage)
	router.POST("/auth/login", authLogin)
	router.NoRoute(pageNotAvailable)

	apiv1 := router.Group("/api/v1/")
	apiv1.POST("/upload", UploadFile)

	// Run
	hostRun := "0.0.0.0:29572"
	webUrl := fmt.Sprintf("http://%s", hostRun)
	log.Println(webUrl)
	router.Run(hostRun)
}
