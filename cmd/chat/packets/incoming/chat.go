package incoming

import (
	"github.com/Steve-Nzr/go-ff/cmd/chat/packets/outgoing"
	"github.com/Steve-Nzr/go-ff/pkg/service/cache"
	"github.com/Steve-Nzr/go-ff/pkg/service/external"
	"github.com/Steve-Nzr/go-ff/pkg/service/messaging"
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
