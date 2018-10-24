package messaging

import "log"

const (
	// ConnectionTopic name
	ConnectionTopic string = "connection"
	// EntityTopic name
	EntityTopic string = "entity"
)

var pubexchangesname = []string{
	ConnectionTopic,
	EntityTopic,
}

func initializePubTopics() {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range pubexchangesname {
		err := ch.ExchangeDeclare(
			name,    // name
			"topic", // type
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)

		if err != nil {
			log.Fatal(err)
		}
	}
}
