package incoming

import (
	"go-ff/chat/packets/outgoing"
	"go-ff/common/service/cache"
	"go-ff/common/service/external"
	"go-ff/common/service/messaging"
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
