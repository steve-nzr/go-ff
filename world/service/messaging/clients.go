package messaging

import (
	"encoding/json"
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/service/messaging/channel"
	"flyff/world/service/messaging/definitions"
	"sync"

	"github.com/streadway/amqp"
)

type netClients map[uint32]*net.Client
type worldClients map[uint32]*entities.PlayerEntity

var NetClients = make(netClients)
var NetClientsMutex sync.RWMutex
var WorldClients = make(worldClients)
var WorldClientsMutex sync.RWMutex

func HandleNewClient(c *net.Client) {
	NetClientsMutex.Lock()
	NetClients[c.ID] = c
	NetClientsMutex.Unlock()

	fullmsg := new(definitions.InternalPacket)
	fullmsg.ID = c.ID
	fullmsg.Todo = definitions.AddTodo

	bytes, err := json.Marshal(fullmsg)
	if err != nil {
		return
	}

	channel.Channel.Publish(
		"packet_in", // exchange
		"#",         // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
}

func HandleRemoveClient(c *net.Client) {
	fullmsg := new(definitions.InternalPacket)
	fullmsg.ID = c.ID
	fullmsg.Todo = definitions.RemoveTodo

	bytes, err := json.Marshal(fullmsg)
	if err != nil {
		return
	}

	channel.Channel.Publish(
		"packet_in", // exchange
		"#",         // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		})
}

func HandleNewMessage(nc *net.Client, packet *net.Packet) {
	fullmsg := new(definitions.InternalPacket)
	fullmsg.ID = nc.ID
	fullmsg.Packet = *packet

	bytes, err := json.Marshal(fullmsg)
	if err != nil {
		return
	}

	channel.Channel.Publish(
		"packet_in", // exchange
		"#",         // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bytes,
		},
	)
}
