package incoming

import (
	"go-ff/pkg/service/cache"
	"go-ff/pkg/service/external"
	"go-ff/pkg/service/messaging"
	"go-ff/cmd/moving/def/packets"
	"go-ff/cmd/moving/feature/move"
	"go-ff/cmd/moving/packets/outgoing"
)

// Behaviour packet handler
func Behaviour(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	behaviourPacket := new(packets.Behaviour)
	behaviourPacket.Construct(p.Packet)

	clientDestPos := behaviourPacket.V.Add(*behaviourPacket.Vd)
	serverDestPos := player.Position.Vec.Add(*behaviourPacket.Vd)
	distance := clientDestPos.Distance(serverDestPos)
	if distance >= 3.0 && distance <= 15.0 {
		go move.ProcessDestPosMove(player.NetClientID, clientDestPos)
	}

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Behaviour(player, behaviourPacket).Finalize(),
		To:     cache.FindIDAroundOnly(player),
	})
}
