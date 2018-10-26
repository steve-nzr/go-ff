package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/moving/feature/move"
	"flyff/moving/packets/outgoing"
)

// DestPos packet from player
func DestPos(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	player.Moving.Vec = *p.Packet.Read3DVector()
	move.SaveMovingPosition(player)

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.DestPos(player).Finalize(),
		To:     cache.FindIDAround(player),
	})

	go move.ProcessDestPosMove(player.NetClientID, player.Moving.Vec)
}
