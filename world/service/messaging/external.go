package messaging

import (
	"encoding/json"
	"flyff/world/service/messaging/channel"
	"flyff/world/service/messaging/definitions"
)

var externalExchangeName = "packet_out"

func HandleExternalPackets() {
	err := channel.Channel.ExchangeDeclare(
		externalExchangeName, // name
		"topic",              // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)

	q, err := channel.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = channel.Channel.QueueBind(
		q.Name,               // queue name
		"#",                  // routing key
		externalExchangeName, // exchange
		false,
		nil)

	msgs, err := channel.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case m := <-msgs:
			{
				var packetout definitions.ExternalPacket
				json.Unmarshal(m.Body, &packetout)

				for _, id := range packetout.To {
					NetClientsMutex.RLock()
					c := NetClients[id]
					NetClientsMutex.RUnlock()

					if c == nil {
						continue
					}

					c.SendFinalized(packetout.Packet)
				}
			}
		}
	}
}
