package channel

import (
	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func Initialize() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.2.201:5672/")
	if err != nil {
		panic(err)
	}

	Channel, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	return conn
}
