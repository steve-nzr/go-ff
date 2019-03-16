package main

import (
	"encoding/json"
	"log"

	"github.com/Steve-Nzr/go-ff/cmd/entity/packets/incoming"
	"github.com/Steve-Nzr/go-ff/pkg/def/packet/packettype"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/database"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
	"github.com/Steve-Nzr/go-ff/pkg/service/messaging"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
