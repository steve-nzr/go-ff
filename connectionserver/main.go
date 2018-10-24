package main

import (
	"encoding/json"
	"flyff/common/service/dotenv"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/connectionserver/service/connectionmanager"
	"log"
)

func onConnectedHandler(ch <-chan *external.Client) {
	for {
		c := <-ch
		if c == nil {
			continue
		}

		connectionmanager.Add(c)
		c.SendGreetings()
	}
}

func onDisconnectedHandler(ch <-chan *external.Client) {
	for {
		c := <-ch
		if c == nil {
			continue
		}

		connectionmanager.Remove(c)
		messaging.Publish(messaging.EntityTopic, &external.PacketHandler{
			ClientID: c.ID,
			Packet:   external.MakePacket(external.Disconnect).FinalizeForInternal(),
		})
	}
}

func onMessageHandler(ch <-chan *external.PacketHandler) {
	for {
		p := <-ch
		if p == nil {
			continue
		}

		// Always FFFFFFF
		p.Packet.ReadUInt32()

		id := p.Packet.ReadUInt32()
		p.Packet.Offset -= (32 / 8)

		switch id {
		case 0xFF00:
			{
				messaging.Publish(messaging.EntityTopic, p)
			}
		case 0x00FF0000:
			{
				messaging.Publish(messaging.ChatTopic, p)
			}
		default:
			{
				log.Printf("Unknown packet '0x%x'", id)
			}
		}
	}
}

func onInternalMessageHandler(ch <-chan []byte) {
	for {
		d := <-ch
		if d == nil {
			continue
		}

		var p external.PacketEmitter
		err := json.Unmarshal(d, &p)
		if err != nil {
			continue
		}

		p.Packet.Finalize()

		for _, id := range p.To {
			c := connectionmanager.Get(id)
			if c == nil {
				continue
			}

			c.SendFinalized(p.Packet)
		}
	}
}

func main() {
	// Initializers ----
	dotenv.Initialize()
	messaging.Initialize()

	// Internal ----
	onInternalMessage := make(chan []byte)
	go messaging.Subscribe(messaging.ConnectionTopic, onInternalMessage)
	go onInternalMessageHandler(onInternalMessage)

	// External ----
	onConnected := make(chan *external.Client)
	onDisconnected := make(chan *external.Client)
	onMessage := make(chan *external.PacketHandler)

	go onConnectedHandler(onConnected)
	go onDisconnectedHandler(onDisconnected)
	go onMessageHandler(onMessage)

	server := external.Create("0.0.0.0:5400")
	server.OnConnected(onConnected)
	server.OnDisconnected(onDisconnected)
	server.OnMessage(onMessage)
	server.Start()
}
