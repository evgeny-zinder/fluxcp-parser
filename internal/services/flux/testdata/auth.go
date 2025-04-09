package testdata

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	User       string
	Password   string
	ServerName string
	ServerURL  string
)

func init() {
	_ = godotenv.Load()

	User = os.Getenv("USER_NAME")
	Password = os.Getenv("PASSWORD")
	ServerName = os.Getenv("SERVER_NAME")
	ServerURL = os.Getenv("SERVER_URL")
}
