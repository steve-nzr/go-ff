package packets

import (
	"github.com/golang/geo/r3"
)

// Behaviour packet struct
type Behaviour struct {
	V             *r3.Vector
	Vd            *r3.Vector
	F             float32
	State         uint32
	StateFlag     uint32
	Motion        uint32
	MotionEx      int32
	Loop          int32
	MotionOptions uint32
	TickCount     int64
}
