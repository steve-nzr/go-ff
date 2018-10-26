package incoming

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/common/service/messaging"
	"flyff/entity/def/packets"
	"flyff/moving/feature/move"
	"flyff/moving/packets/outgoing"
)

// Behaviour packet handler
func Behaviour(p *external.PacketHandler) {
	player := cache.FindByNetID(p.ClientID)
	if player == nil {
		return
	}

	behaviourPacket := new(packets.Behaviour)
	behaviourPacket.V = p.Packet.Read3DVector()
	behaviourPacket.Vd = p.Packet.Read3DVector()
	behaviourPacket.F = p.Packet.ReadFloat32()
	behaviourPacket.State = p.Packet.ReadUInt32()
	behaviourPacket.StateFlag = p.Packet.ReadUInt32()
	behaviourPacket.Motion = p.Packet.ReadUInt32()
	behaviourPacket.MotionEx = p.Packet.ReadInt32()
	behaviourPacket.Loop = p.Packet.ReadInt32()
	behaviourPacket.MotionOptions = p.Packet.ReadUInt32()
	behaviourPacket.TickCount = p.Packet.ReadInt64()

	clientDestPos := behaviourPacket.V.Add(*behaviourPacket.Vd)
	serverDestPos := player.Position.Vec.Add(*behaviourPacket.Vd)
	if clientDestPos.Distance(serverDestPos) >= 3.0 {
		go move.ProcessDestPosMove(player.NetClientID, clientDestPos)
	}

	messaging.Publish(messaging.ConnectionTopic, &external.PacketEmitter{
		Packet: outgoing.Behaviour(player, behaviourPacket).Finalize(),
		To:     cache.FindIDAroundOnly(player),
	})
}
