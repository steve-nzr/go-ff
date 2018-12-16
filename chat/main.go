package main

import (
	"encoding/json"
	"flyff/chat/packets/incoming"
	"flyff/common/def/packet/packettype"
	"flyff/common/service/cache"
	"flyff/common/service/dotenv"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dotenv.Initialize()
	cache.Initialize()
	messaging.Initialize()

	ch := make(chan []byte)
	go messaging.Subscribe(messaging.ChatTopic, ch)

	for {
		b := <-ch
		p := new(external.PacketHandler)
		if err := json.Unmarshal(b, p); err != nil {
			fmt.Print(err)
			continue
		}

		id := p.Packet.ReadUInt32()
		switch id {
		case packettype.Chat:
			{
				incoming.Chat(p)
				break
			}
		}
	}
}
