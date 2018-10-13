package out

import (
	"flyff/core/net"
	"flyff/world/entities"
)

func MakeDestPos(pe *entities.PlayerEntity) net.Packet {
	p := net.StartMergePacket(uint32(pe.ID), uint16(0x00C1), 0x0000FF00).
		WriteFloat32(float32(pe.Moving.X)).
		WriteFloat32(float32(pe.Moving.Y)).
		WriteFloat32(float32(pe.Moving.Z)).
		WriteUInt8(1)

	return p
}
