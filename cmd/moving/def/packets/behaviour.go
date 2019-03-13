package packets

import (
	"go-ff/pkg/service/external"

	"github.com/golang/geo/r3"
)

// Behaviour packet struct
type Behaviour struct {
	V             *r3.Vector
	Vd            *r3.Vector
	Angle         float32
	State         uint32
	StateFlag     uint32
	Motion        uint32
	MotionEx      int32
	Loop          int32
	MotionOptions uint32
	TickCount     int64
}

// Construct from Packet
func (b *Behaviour) Construct(p *external.Packet) {
	b.V = p.Read3DVector()
	b.Vd = p.Read3DVector()
	b.Angle = p.ReadFloat32()
	b.State = p.ReadUInt32()
	b.StateFlag = p.ReadUInt32()
	b.Motion = p.ReadUInt32()
	b.MotionEx = p.ReadInt32()
	b.Loop = p.ReadInt32()
	b.MotionOptions = p.ReadUInt32()
	b.TickCount = p.ReadInt64()
}
