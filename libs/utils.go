package libs

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// ////////////////// Token Functions //////////////////

// Reads the token from the environment variable
func TokenReader() string {
	token := os.Getenv("TOKEN")
	if token == "" {
		logmaker().Error("No token found in environment variable")
		return ""
	}
	return token
}

// Token Length checker
func TokenLengthChecker(token string) bool {
	if len(token) != 32 {
		logmaker().Error("Token length is not the required 32 characters")
		return false
	}
	return true
}

// Checks if the character is alphanumeric
func isAlphanumeric(char rune) bool {
	return (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}

// Token Alphanumeric checker
func TokenAlphanumericChecker(token string) bool {
	for _, char := range token {
		if !isAlphanumeric(char) {
			logmaker().Error("Token contains non-alphanumeric characters")
			return false
		}
	}
	return true
}

// ///////////////// JSON Functions //////////////////

func JsonWriter(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		logmaker().Error("Error creating file: ", zap.String("error", err.Error()))
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		logmaker().Error("Error writing to file: ", zap.String("error", err.Error()))
		return
	}
}

// ///////////////// HTTP Functions //////////////////

func GetData(url string) ([]byte, error) {
	// This function should be implemented to make an HTTP GET request to the given URL
	// and return the response body as a byte slice.
	// For now, we'll just return nil and an error.
	return nil, fmt.Errorf("GetData not implemented")
}
