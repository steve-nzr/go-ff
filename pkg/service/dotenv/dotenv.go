package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Initialize the service
// Loads the .env file
func Initialize() {
	filename, ok := os.LookupEnv("ENV_FILE")
	if ok == false {
		filename = ".env"
	}
	err := godotenv.Load("./config/" + filename)
	if err != nil {
		log.Print(err.Error())
	}
}
