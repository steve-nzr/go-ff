package in

import (
	"flyff/core/net"
	"flyff/world/game/structure"
	"flyff/world/packets/out"
	"flyff/world/service/gamemap"
)

func DestPos(wc *structure.WorldClient, p *net.Packet) {
	wc.PlayerEntity.Moving.Destination.X = float64(p.ReadFloat32())
	wc.PlayerEntity.Moving.Destination.Y = float64(p.ReadFloat32())
	wc.PlayerEntity.Moving.Destination.Z = float64(p.ReadFloat32())

	destPosPacket := out.MakeDestPos(wc.PlayerEntity)
	gamemap.Manager.SendFrom(wc, &destPosPacket)
}
