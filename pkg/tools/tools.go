package tools

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// Random string
func RandomString(length int) string {
	rand.Seed(time.Now().Unix())

	var output strings.Builder

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJULMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return (output.String())
}

// Create folder
func CreateFolder(FolderName string) bool {
	return_check := false
	if _, err := os.Stat(FolderName); os.IsNotExist(err) {
		err = os.Mkdir(FolderName, 0755)
		if err != nil {
			// fmt.Println(err)
			return_check = false
		} else {
			return_check = true
		}
	} else {
		return_check = true
	}
	return return_check
}
