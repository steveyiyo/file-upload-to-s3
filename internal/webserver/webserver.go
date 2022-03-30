package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/image-upload/internal/api"
)

func pageNotAvailable(c *gin.Context) {
	c.String(404, "Page not available")
}

func Init() {

	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.POST("/api/v1/upload", api.Uploadfile)

	router.NoRoute(pageNotAvailable)

	router.Run("127.0.0.1:19725")

}
