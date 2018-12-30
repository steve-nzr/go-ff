package main

import (
	"encoding/json"
	"go-ff/action/packets/incoming"
	"go-ff/common/def/packet/packettype"
	"go-ff/common/service/cache"
	"go-ff/common/service/dotenv"
	"go-ff/common/service/external"
	"go-ff/common/service/messaging"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dotenv.Initialize()
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
