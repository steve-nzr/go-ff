package outgoing

import (
	"flyff/common/def/packet/snapshottype"
	"flyff/common/service/cache"
	"flyff/common/service/external"
	"flyff/moving/def/packets"
)

// Behaviour packet emitter
func Behaviour(p *cache.Player, b *packets.Behaviour) *external.Packet {
	return external.StartMergePacket(p.EntityID, snapshottype.Moverbehavior, 0x0000FF00).
		Write3DVector(b.V).
		Write3DVector(b.Vd).
		WriteFloat32(b.Angle).
		WriteUInt32(b.State).
		WriteUInt32(b.StateFlag).
		WriteUInt32(b.Motion).
		WriteInt32(b.MotionEx).
		WriteInt32(b.Loop).
		WriteUInt32(b.MotionOptions).
		WriteInt64(b.TickCount)
}
