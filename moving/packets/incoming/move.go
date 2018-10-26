package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/moving/def/packets"
	"flyff/moving/feature/move"
	"flyff/moving/packets/outgoing"
	"fmt"

	"github.com/golang/geo/r3"
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

// Move by key handler
func Move(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	behaviourPacket := new(packets.Behaviour)
	behaviourPacket.Construct(p.Packet)

	player.Moving.Vec = r3.Vector{}
	move.SaveMovingPosition(player)

	fmt.Println("Motion :", behaviourPacket.Motion)
}
