package messaging

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// Publish a messager to the given topic
// The data parameter needs to be a pointer to type
func Publish(topic string, data interface{}) {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.Publish(topic, "", false, false, amqp.Publishing{
		Body: bytes,
	})
	if err != nil {
		log.Fatal(err)
	}
}
