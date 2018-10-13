package in

import (
	"flyff/core/net"
	"flyff/world/entities"
	"flyff/world/packets/out"
	"flyff/world/service/gamemap"
)

func DestPos(pe *entities.PlayerEntity, p *net.Packet) {
	pe.Moving.X = float64(p.ReadFloat32())
	pe.Moving.Y = float64(p.ReadFloat32())
	pe.Moving.Z = float64(p.ReadFloat32())

	destPosPacket := out.MakeDestPos(pe)
	gamemap.Manager.SendFrom(pe, &destPosPacket)
}
