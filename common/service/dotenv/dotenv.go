package dotenv

import (
	"github.com/joho/godotenv"
)

// Initialize the service
// Loads the .env file
func Initialize() {
	godotenv.Load("../config/.env")
}
