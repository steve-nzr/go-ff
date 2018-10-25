package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/moving/packets/outgoing"
)

// DestPos packet from player
func DestPos(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	player.Moving.Vec.X = float64(p.Packet.ReadFloat32())
	player.Moving.Vec.Y = float64(p.Packet.ReadFloat32())
	player.Moving.Vec.Z = float64(p.Packet.ReadFloat32())

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.DestPos(player).Finalize(),
		To:     cache.FindIDAround(player),
	})
}
