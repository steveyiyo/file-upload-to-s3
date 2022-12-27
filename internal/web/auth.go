package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steveyiyo/file-upload-to-s3/internal/jwt"
)

func authLoginPage(c *gin.Context) {
	checkSession := VerifySessionAuth(c)

	if checkSession.jwtSuccess {
		c.Redirect(302, "/")
	} else {
		c.HTML(200, "login.tmpl", gin.H{})
	}
}

func authLogin(c *gin.Context) {
	var jwtToken string

	// Check to oauth2

	// If oauth2 success, redirect to /
	jwtToken, _ = jwt.GenerateToken("steveyiyo")
	c.SetCookie("jwt", jwtToken, 86400, "/", "", false, true)
	c.Redirect(302, "/")
}

// Process Cookie to get JWT Token
func getJWTToken(cookies []*http.Cookie) string {
	jwtToken := ""
	for _, cookie := range cookies {
		if cookie.Name == "jwt" {
			jwtToken = cookie.Value
		}
	}
	return jwtToken
}

// Verify Session Auth
type sessionAuthData struct {
	jwtSuccess bool
	username   string
}

func VerifySessionAuth(c *gin.Context) sessionAuthData {
	jwtToken := getJWTToken(c.Request.Cookies())
	jwtSuccess := false
	username := ""
	jwtCheck := false

	if jwtToken != "" {
		jwtCheck, username = jwt.ValidateToken(jwtToken)
		if jwtCheck {
			jwtSuccess = true
		} else {
			jwtSuccess = false
		}
	}
	returnResult := sessionAuthData{jwtSuccess, username}

	return returnResult
}
