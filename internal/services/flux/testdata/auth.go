package testdata

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/eznd-go/flux/pkg/projectpath"
)

var (
	User       string
	Password   string
	ServerName string
	ServerURL  string
)

func init() {
	_ = godotenv.Load(projectpath.Root + "/.env")
	User = os.Getenv("USER_NAME")
	Password = os.Getenv("PASSWORD")
	ServerName = os.Getenv("SERVER_NAME")
	ServerURL = os.Getenv("SERVER_URL")
}
