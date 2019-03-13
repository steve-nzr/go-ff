package messaging

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channels map[string]*amqp.Channel

var messageExpiration = "5000"

// Initialize the AMQP Connection
func Initialize() {
	conn, err := amqp.Dial(os.Getenv("AMQP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	connection = conn
	initializePubTopics()

	messageExpiration = os.Getenv("MESSAGE_EXPIRATION")
}
