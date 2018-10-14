package messaging

import (
	"encoding/json"
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/packets/in"
	"flyff/world/service/gamemap"
	"flyff/world/service/messaging/channel"
	"flyff/world/service/messaging/definitions"
	"fmt"

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
		select {
		case msg := <-msgs:
			{
				var fullmsg definitions.InternalPacket
				json.Unmarshal(msg.Body, &fullmsg)

				var pe *entities.PlayerEntity
				if fullmsg.Todo == definitions.AddTodo {
					pe = new(entities.PlayerEntity)
					pe.NetClientID = fullmsg.ID

					WorldClientsMutex.Lock()
					WorldClients[fullmsg.ID] = pe
					WorldClientsMutex.Unlock()

					p := net.MakePacket(net.GREETINGS).
						WriteUInt32(fullmsg.ID).
						Finalize()

					pe.Send(p)
					continue
				} else if fullmsg.Todo == definitions.RemoveTodo {
					WorldClientsMutex.RLock()
					pe = WorldClients[fullmsg.ID]
					WorldClientsMutex.RUnlock()

					gamemap.Manager.Unregister(pe)

					WorldClientsMutex.Lock()
					delete(WorldClients, fullmsg.ID)
					WorldClientsMutex.Unlock()

					NetClientsMutex.Lock()
					delete(NetClients, fullmsg.ID)
					NetClientsMutex.Unlock()
					continue
				} else {
					WorldClientsMutex.RLock()
					pe = WorldClients[fullmsg.ID]
					WorldClientsMutex.RUnlock()
				}

				if pe == nil {
					continue
				}

				// Always FFFFFFF
				fullmsg.Packet.ReadUInt32()

				go handlePacket(pe, &fullmsg.Packet)

				fmt.Println("MESSAGE IN !")
			}
		}
	}
}

var redisURL = "192.168.2.201:6379"
var client = redis.NewClient(&redis.Options{
	Network: "tcp",
	Addr:    redisURL,
})

func handlePacket(pe *entities.PlayerEntity, p *net.Packet) {
	protocol := p.ReadUInt32()

	switch protocol {
	case 0xff00:
		{
			lock, err := lock.Obtain(client, "lock.foo", &lock.Options{
				RetryCount: 3,
			})
			if err != nil {
				fmt.Println("Can't lock !")
			} else {
				fmt.Println("LOCKED")
				lock.Unlock()
			}

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
