package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/steveyiyo/file-upload-to-s3/pkg/tools"
)

var jwtKey = []byte(tools.RandomString(32))

type Token struct {
	jwt.StandardClaims
	UserID string
}

// Generate Token
func GenerateToken(userId string) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()

	value := Token{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		UserID: userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, value)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Validate Token
func ValidateToken(tokenString string) (bool, string) {
	var claims Token
	token, _ := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	var returnSuccess bool
	var returnMessage string

	if !token.Valid {
		returnSuccess = false
		returnMessage = ""
	} else {
		returnSuccess = true
		returnMessage = claims.UserID
	}

	return returnSuccess, returnMessage
}
