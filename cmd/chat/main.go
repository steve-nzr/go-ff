package main

import (
	"encoding/json"
	"github.com/Steve-Nzr/go-ff/cmd/chat/packets/incoming"
	"github.com/Steve-Nzr/go-ff/pkg/def/packet/packettype"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/dotenv"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
	"github.com/Steve-Nzr/go-ff/pkg/service/messaging"
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
			log.Print(err)
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
