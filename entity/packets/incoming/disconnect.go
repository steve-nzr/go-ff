package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/entity/packets/outgoing"
)

// Disconnect from World
func Disconnect(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	cache.Connection.Where("net_client_id = ?", p.ClientID).Delete(player)

	// Make Visible ----
	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.DeleteObj(player),
		To:     cache.FindIDAround(player),
	})
}
