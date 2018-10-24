package incoming

import (
	"flyff/chat/packets/outgoing"
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"fmt"
)

// Chat packet (Basic chat & commands)
func Chat(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	fmt.Println("Processing")
	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Chat(player, p.Packet.ReadString()),
		To:     cache.FindIDAround(player),
	})
}
