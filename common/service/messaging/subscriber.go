package messaging

import (
	"log"
)

// Subscribe to all messages in the given topic
// Warning, this is a blocking function
// Consider call it in a Goroutine
func Subscribe(topic string, msgs chan<- []byte) {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		"",    // name (random)
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		topic,  // exchange
		false,
		nil)
	if err != nil {
		log.Fatal(err)
	}

	consumeChan, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := <-consumeChan
		msgs <- msg.Body
	}
}
