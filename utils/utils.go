package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnviromentalVariables(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(".env file is missing")
	}
	return os.Getenv(key)
}


