package main

import (
	"encoding/json"
	"flyff/common/service/cache"
	"flyff/common/service/database"
	"flyff/common/service/dotenv"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/entity/packets/incoming"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dotenv.Initialize()
	database.Initialize()
	cache.Initialize()
	messaging.Initialize()

	ch := make(chan []byte)
	go messaging.Subscribe(messaging.EntityTopic, ch)

	for {
		b := <-ch
		p := new(external.PacketHandler)
		if err := json.Unmarshal(b, p); err != nil {
			fmt.Print(err)
			continue
		}

		id := p.Packet.ReadUInt32()
		switch id {
		case 0xff00:
			{
				incoming.Join(p)
				break
			}
		case uint32(external.Disconnect):
			{
				incoming.Disconnect(p)
				break
			}
		}
	}
}
