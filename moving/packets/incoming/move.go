package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/moving/def/packets"
	"flyff/moving/feature/move"
	"flyff/moving/packets/outgoing"

	"github.com/golang/geo/r3"
)

// DestPos packet from player
func DestPos(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	player.Moving.Destination = *p.Packet.Read3DVector()
	player.Moving.Motion = 0
	move.SaveMovingComponent(player)

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.DestPos(player).Finalize(),
		To:     cache.FindIDAround(player),
	})

	go move.ProcessDestPosMove(player.NetClientID, player.Moving.Destination)
}

// Move by key handler
func Move(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	behaviourPacket := new(packets.Behaviour)
	behaviourPacket.Construct(p.Packet)

	player.Moving.Destination = r3.Vector{}
	player.Moving.Motion = behaviourPacket.Motion
	player.Moving.Angle = float64(behaviourPacket.Angle)
	move.SaveMovingComponent(player)

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Move(player, behaviourPacket).Finalize(),
		To:     cache.FindIDAroundOnly(player),
	})

	if behaviourPacket.Motion == 5 {
		go move.ProcessMove(p.ClientID, player.Moving.Angle)
	}
}
