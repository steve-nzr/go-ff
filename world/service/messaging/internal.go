package messaging

import (
	"encoding/json"
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/packets/in"
	"flyff/world/service/gamemap"
	"flyff/world/service/messaging/channel"
	"flyff/world/service/messaging/definitions"
	"flyff/world/service/playerstate"
	"fmt"
	"strconv"
	"time"

	"github.com/streadway/amqp"

	lock "github.com/bsm/redis-lock"
	"github.com/go-redis/redis"
)

var internalExchangeName = "packet_in"

func HandleInternalPackets() {
	err := channel.Channel.ExchangeDeclare(
		internalExchangeName, // name
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
		internalExchangeName, // exchange
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
		go handleQueueMessage(<-msgs)
	}
}

func handleQueueMessage(msg amqp.Delivery) {
	var fullmsg definitions.InternalPacket
	json.Unmarshal(msg.Body, &fullmsg)

	var pe = new(entities.PlayerEntity)
	pe.NetClientID = fullmsg.ID

	if fullmsg.Todo == definitions.AddTodo {
		go playerstate.Connection.Save(pe)

		p := net.MakePacket(net.GREETINGS).
			WriteUInt32(fullmsg.ID).
			Finalize()

		pe.Send(p)
		return
	} else if fullmsg.Todo == definitions.RemoveTodo {
		playerstate.Connection.First(pe)
		gamemap.Manager.Unregister(pe)
		go playerstate.Connection.Delete(pe)

		NetClientsMutex.Lock()
		delete(NetClients, fullmsg.ID)
		NetClientsMutex.Unlock()
		return
	} else {
		playerstate.Connection.First(pe)
	}

	if pe == nil {
		return
	}

	// Always FFFFFFF
	fullmsg.Packet.ReadUInt32()

	handlePacket(pe, &fullmsg.Packet)
}

var redisURL = "192.168.2.201:6379"
var client = redis.NewClient(&redis.Options{
	Network: "tcp",
	Addr:    redisURL,
})

func handlePacket(pe *entities.PlayerEntity, p *net.Packet) {
	protocol := p.ReadUInt32()

	playerLock, err := lock.Obtain(client, strconv.FormatUint(uint64(pe.NetClientID), 10), &lock.Options{
		RetryCount:  5,
		LockTimeout: 200 * time.Millisecond,
		RetryDelay:  50 * time.Millisecond,
	})
	if err != nil {
		fmt.Println("Can't lock main player !")
		return
	}
	defer playerLock.Unlock()
	defer playerstate.Connection.Save(pe)

	switch protocol {
	case 0xff00:
		{
			in.Join(pe, p)
		}
	case 0xffffff00:
		{
			p.ReadUInt8()
			snapshotProtocol := p.ReadUInt16()
			if snapshotProtocol == 0x00C1 {
				in.DestPos(pe, p)
			}
		}
	case 0x00FF0000:
		{
			in.Chat(pe, p)
		}
	}
}
