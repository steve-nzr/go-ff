package main

import (
	"encoding/json"
	"flyff/common/service/cache"
	"flyff/common/service/dotenv"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/moving/packets/incoming"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dotenv.Initialize()
	cache.Initialize()
	messaging.Initialize()

	ch := make(chan []byte)
	go messaging.Subscribe(messaging.MovingTopic, ch)

	for {
		b := <-ch
		p := new(external.PacketHandler)
		if err := json.Unmarshal(b, p); err != nil {
			fmt.Print(err)
			continue
		}

		id := p.Packet.ReadUInt32()
		switch id {
		case 0xffffff00:
			{
				p.Packet.ReadUInt8()
				snapshotProtocol := p.Packet.ReadUInt16()
				if snapshotProtocol == 0x00C1 {
					incoming.DestPos(p)
				}
			}
		case 0xFFFFFF01:
			{
				incoming.Move(p)
			}
		case 0xffffff02:
			{
				incoming.Behaviour(p)
			}
		}
	}
}
