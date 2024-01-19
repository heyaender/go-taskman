package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	env := godotenv.Load(".env")
	if env != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}
