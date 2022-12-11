package getenvironment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Cant load .env file")
	}
	return os.Getenv(key)
}
