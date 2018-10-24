package messaging

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channels map[string]*amqp.Channel

// Initialize the AMQP Connection
func Initialize() {
	conn, err := amqp.Dial(os.Getenv("AMQP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	connection = conn
	initializePubTopics()
}
