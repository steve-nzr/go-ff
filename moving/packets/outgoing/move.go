package outgoing

import (
	"flyff/common/service/cache"
	"flyff/common/service/external"
)

// DestPos packet
func DestPos(p *cache.Player) *external.Packet {
	packet := external.StartMergePacket(p.EntityID, uint16(0x00C1), 0x0000FF00).
		WriteFloat32(float32(p.Moving.Vec.X)).
		WriteFloat32(float32(p.Moving.Vec.Y)).
		WriteFloat32(float32(p.Moving.Vec.Z)).
		WriteUInt8(1)

	return packet.Finalize()
}
