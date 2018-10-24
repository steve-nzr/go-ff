package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

// Initialize the service
// Loads the .env file
func Initialize() {
	godotenv.Load(os.Getenv("GOPATH") + "/src/flyff/.env")
}
