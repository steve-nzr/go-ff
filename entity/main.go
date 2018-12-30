package main

import (
	"encoding/json"
	"go-ff/common/def/packet/packettype"
	"go-ff/common/service/cache"
	"go-ff/common/service/database"
	"go-ff/common/service/dotenv"
	"go-ff/common/service/external"
	"go-ff/common/service/messaging"
	"go-ff/entity/packets/incoming"
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
			log.Print(err)
			continue
		}

		id := p.Packet.ReadUInt32()
		switch id {
		case packettype.Join:
			{
				incoming.Join(p)
				break
			}
		case packettype.Disconnect:
			{
				incoming.Disconnect(p)
				break
			}
		}
	}
}
