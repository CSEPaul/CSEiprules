package libs

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func logmaker() *zap.Logger {
	logger := CreateLogger()
	defer logger.Sync()
	return logger
}

// Reads the token from the environment variable
func TokenReader() string {
	token := os.Getenv("TOKEN")
	if token == "" {
		logger := logmaker()
		logger.Error("No token found in environment variable")
		fmt.Println("No token found in environment variable")
		fmt.Println("Please set the TOKEN environment variable")
		return ""
	}
	return token

}

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
