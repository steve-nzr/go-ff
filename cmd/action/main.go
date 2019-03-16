package main

import (
	"encoding/json"
	"log"

	"github.com/Steve-Nzr/go-ff/cmd/action/packets/incoming"
	"github.com/Steve-Nzr/go-ff/pkg/def/packet/packettype"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
	"github.com/Steve-Nzr/go-ff/pkg/service/messaging"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cache.Initialize()
	messaging.Initialize()

	log.Println("Server is up !")

	ch := make(chan []byte)
	go messaging.Subscribe(messaging.ActionTopic, ch)

	for {
		b := <-ch
		p := new(external.PacketHandler)
		if err := json.Unmarshal(b, p); err != nil {
			log.Print(err)
			continue
		}

		id := p.Packet.ReadUInt32()
		switch id {
		case packettype.Doequip:
			{
				incoming.EquipItem(p)
				break
			}
		case packettype.Moveitem:
			{
				incoming.MoveItem(p)
				break
			}
		case packettype.Dropitem:
			{
				incoming.DropItem(p)
				break
			}
		}
	}
}
