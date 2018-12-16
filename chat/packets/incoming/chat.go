package incoming

import (
	"flyff/chat/packets/outgoing"
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
)

// Chat packet (Basic chat & commands)
func Chat(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Chat(player, p.Packet.ReadString()),
		To:     cache.FindIDAround(player),
	})
}
