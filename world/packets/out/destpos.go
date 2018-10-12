package out

import (
	"flyff/core/net"
	"flyff/world/game/structure"
)

func MakeDestPos(pe *structure.PlayerEntity) net.Packet {
	p := net.StartMergePacket(uint32(pe.ID), uint16(0x00C1), 0x0000FF00).
		WriteFloat32(float32(pe.Moving.Destination.X)).
		WriteFloat32(float32(pe.Moving.Destination.Y)).
		WriteFloat32(float32(pe.Moving.Destination.Z)).
		WriteUInt8(1)

	return p
}
